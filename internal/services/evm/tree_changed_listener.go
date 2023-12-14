package evm

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/evm-worldcoin-saver-svc/pkg/worldid"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rarimo/evm-worldcoin-saver-svc/internal/config"
	oracletypes "github.com/rarimo/rarimo-core/x/oraclemanager/types"
	"github.com/rarimo/saver-grpc-lib/broadcaster"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

type listener struct {
	log         *logan.Entry
	filter      filterer
	broadcaster broadcaster.Broadcaster
	block       blockGetter
	listenerCfg
}

type listenerCfg struct {
	txCreatorAddr string
	homeChain     string
	fromBlock     uint64
	blockWindow   uint64
	maxBlocks     uint64
}

type filterer interface {
	FilterTreeChanged(*bind.FilterOpts, []*big.Int, []uint8, []*big.Int) (*worldid.WorldIdTreeChangedIterator, error)
}

type blockGetter interface {
	BlockNumber(context.Context) (uint64, error)
	HeaderByNumber(context.Context, *big.Int) (*ethtypes.Header, error)
}

func RunWorldCoinListener(ctx context.Context, cfg config.Config) {
	const runnerName = "worldcoin_listener"
	log := cfg.Log().WithField("who", runnerName)

	handler, err := worldid.NewWorldIdFilterer(cfg.Ethereum().ContractAddr, cfg.Ethereum().RPCClient)
	if err != nil {
		panic(errors.Wrap(err, "failed to init WorldCoin tree change filterer"))
	}

	l := listener{
		log:         log,
		broadcaster: cfg.Broadcaster(),
		filter:      handler,
		block:       cfg.Ethereum().RPCClient,
		listenerCfg: listenerCfg{
			txCreatorAddr: cfg.Broadcaster().Sender(),
			homeChain:     cfg.Ethereum().NetworkName,
			fromBlock:     cfg.Ethereum().StartFromBlock,
			blockWindow:   cfg.Ethereum().BlockWindow,
			maxBlocks:     cfg.Ethereum().MaxBlocksPerRequest,
		},
	}

	running.WithBackOff(ctx, log, runnerName,
		l.subscription,
		30*time.Second, 5*time.Second, 30*time.Second)
}

func (l *listener) subscription(ctx context.Context) error {
	lastBlock, err := l.defineLastBlock(ctx)
	if err != nil {
		return fmt.Errorf("failed to define last block: %w", err)
	}

	l.log.Infof("Starting subscription from %d to %d", l.fromBlock, lastBlock)
	defer l.log.Info("Subscription finished")

	iter, err := l.filter.FilterTreeChanged(&bind.FilterOpts{
		Start:   l.fromBlock,
		End:     &lastBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to filter WorldID TreeChanged events: %w", err)
	}

	l.processEvents(ctx, iter)
	// https://ethereum.stackexchange.com/questions/8199/are-both-the-eth-newfilter-from-to-fields-inclusive
	// End in FilterLogs is inclusive
	l.fromBlock = lastBlock + 1
	return nil
}

func (l *listener) defineLastBlock(ctx context.Context) (uint64, error) {
	lastBlock, err := l.block.BlockNumber(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block number: %w", err)
	}

	lastBlock -= l.blockWindow
	if lastBlock < l.fromBlock {
		l.log.Infof("Skipping window: start %d > finish %d", l.fromBlock, lastBlock)
		return 0, nil
	}

	if l.fromBlock+l.maxBlocks < lastBlock {
		l.log.Debugf("maxBlockPerRequest limit would be exceeded: setting last block to %d instead of %d", l.fromBlock+l.maxBlocks, lastBlock)
		lastBlock = l.fromBlock + l.maxBlocks
	}

	return lastBlock, nil
}

func (l *listener) processEvents(ctx context.Context, iter *worldid.WorldIdTreeChangedIterator) {
	for iter.Next() {
		evt := iter.Event
		if evt == nil {
			l.log.Error("got nil event")
			continue
		}

		l.log.WithFields(logan.F{
			"tx_hash":   evt.Raw.TxHash,
			"tx_index":  evt.Raw.TxIndex,
			"log_index": evt.Raw.Index,
		}).Debugf("Got event: preRoot=%s kind=%d postRoot=%s", evt.PreRoot, evt.Kind, evt.PostRoot)

		msg, err := l.msgFromEvent(ctx, evt)
		if err != nil {
			l.log.WithError(err).WithField("tx_hash", evt.Raw.TxHash.String()).Error("failed to convert event to WorldCoin msg")
			continue
		}

		if err = l.broadcaster.BroadcastTx(ctx, msg); err != nil {
			l.log.WithError(err).WithField("tx_hash", evt.Raw.TxHash.String()).Error(err, "failed to broadcast WorldCoin identity transfer msg")
			continue
		}
	}
}

func (l *listener) msgFromEvent(ctx context.Context, evt *worldid.WorldIdTreeChanged) (*oracletypes.MsgCreateWorldCoinIdentityTransferOp, error) {
	header, err := l.block.HeaderByNumber(ctx, new(big.Int).SetUint64(evt.Raw.BlockNumber))
	if err != nil {
		return nil, fmt.Errorf("failed to get block header by number from event: %w", err)
	}

	return &oracletypes.MsgCreateWorldCoinIdentityTransferOp{
		Creator:     l.txCreatorAddr,
		Contract:    evt.Raw.Address.Hex(),
		Chain:       l.homeChain,
		PrevState:   hexutil.Encode(evt.PreRoot.Bytes()),
		State:       hexutil.Encode(evt.PostRoot.Bytes()),
		Timestamp:   strconv.FormatUint(header.Time, 10),
		BlockNumber: evt.Raw.BlockNumber,
	}, nil
}

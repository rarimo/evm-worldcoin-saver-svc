package voting

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/evm-worldcoin-saver-svc/internal/config"
	"github.com/rarimo/evm-worldcoin-saver-svc/pkg/worldid"
	rarimocore "github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/rarimo/saver-grpc-lib/voter"
	"github.com/rarimo/saver-grpc-lib/voter/verifiers"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type verifier struct {
	log    *logan.Entry
	filter filterer
	header headerGetter
	chain  string
}

type filterer interface {
	FilterTreeChanged(*bind.FilterOpts, []*big.Int, []uint8, []*big.Int) (*worldid.WorldIdTreeChangedIterator, error)
}

type headerGetter interface {
	HeaderByNumber(context.Context, *big.Int) (*ethtypes.Header, error)
}

var _ voter.Verifier = &verifier{}

func NewWorldCoinVerifier(cfg config.Config) voter.Verifier {
	filter, err := worldid.NewWorldIdFilterer(cfg.Ethereum().ContractAddr, cfg.Ethereum().RPCClient)
	if err != nil {
		panic(fmt.Errorf("failed to init TreeChanged event filterer for verifier: %w", err))
	}

	return &verifier{
		log:    cfg.Log(),
		filter: filter,
		header: cfg.Ethereum().RPCClient,
		chain:  cfg.Ethereum().NetworkName,
	}
}

func (v *verifier) Verify(ctx context.Context, operation rarimocore.Operation) (rarimocore.VoteType, error) {
	if operation.OperationType != rarimocore.OpType_WORLDCOIN_IDENTITY_TRANSFER {
		v.log.Debugf("Voted NO: invalid operation type")
		return rarimocore.VoteType_NO, verifiers.ErrInvalidOperationType
	}

	var op rarimocore.WorldCoinIdentityTransfer
	if err := proto.Unmarshal(operation.Details.Value, &op); err != nil {
		v.log.Debugf("Voted NO: failed to unmarshal")
		return rarimocore.VoteType_NO, err
	}

	err := v.verify(ctx, op)
	if err == nil {
		return rarimocore.VoteType_YES, nil
	}

	v.log.WithError(err).Debugf("Voted NO: received an error from verifier")
	switch errors.Cause(err) {
	case verifiers.ErrUnsupportedNetwork:
		return rarimocore.VoteType_NO, verifiers.ErrUnsupportedNetwork
	case verifiers.ErrWrongOperationContent:
		return rarimocore.VoteType_NO, nil
	default:
		return rarimocore.VoteType_NO, err
	}
}

func (v *verifier) verify(ctx context.Context, op rarimocore.WorldCoinIdentityTransfer) error {
	evt, err := v.getEvent(ctx, op)
	if err != nil {
		return fmt.Errorf("failed to get event: %w", err)
	}

	num := new(big.Int).SetUint64(evt.Raw.BlockNumber)
	header, err := v.header.HeaderByNumber(ctx, num)
	if err != nil {
		return fmt.Errorf("failed to get block header by number from event: %w", err)
	}

	if !proto.Equal(&op, opFromExternalData(evt, v.chain, header.Time)) {
		return verifiers.ErrWrongOperationContent
	}

	return nil
}

func (v *verifier) getEvent(ctx context.Context, op rarimocore.WorldCoinIdentityTransfer) (*worldid.WorldIdTreeChanged, error) {
	iter, err := v.filter.FilterTreeChanged(&bind.FilterOpts{
		Context: ctx,
		Start:   op.BlockNumber,
		End:     &op.BlockNumber, // end block is inclusive
	}, toBigIntFilter(op.PrevState), nil, toBigIntFilter(op.State))

	if err != nil {
		return nil, fmt.Errorf("failed to filter WorldID changed events: %w", err)
	}
	if !iter.Next() || iter.Event == nil {
		return nil, errors.New("event not found")
	}

	res := *iter.Event
	if iter.Next() {
		return nil, errors.New("multiple events found for given block, preRoot and postRoot")
	}

	return &res, nil
}

func opFromExternalData(evt *worldid.WorldIdTreeChanged, chain string, timestamp uint64) *rarimocore.WorldCoinIdentityTransfer {
	return &rarimocore.WorldCoinIdentityTransfer{
		Contract:    evt.Raw.Address.Hex(),
		Chain:       chain,
		PrevState:   hexutil.EncodeBig(evt.PreRoot),
		State:       hexutil.EncodeBig(evt.PostRoot),
		Timestamp:   strconv.FormatUint(timestamp, 10),
		BlockNumber: evt.Raw.BlockNumber,
	}
}

func toBigIntFilter(s string) []*big.Int {
	b, ok := new(big.Int).SetString(s, 16)
	if !ok {
		return nil
	}
	return []*big.Int{b}
}

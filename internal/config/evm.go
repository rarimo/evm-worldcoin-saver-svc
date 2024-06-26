package config

import (
	"context"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Ethereum struct {
	ContractAddr        common.Address    `fig:"contract_addr,required"`
	RPCClient           *ethclient.Client `fig:"rpc,required"`
	StartFromBlock      uint64            `fig:"start_from_block"`
	NetworkName         string            `fig:"network_name,required"`
	BlockWindow         uint64            `fig:"block_window,required"`
	MaxBlocksPerRequest uint64            `fig:"max_blocks_per_request"`
}

func (c *config) Ethereum() *Ethereum {
	return c.ethereum.Do(func() interface{} {
		cfg := Ethereum{}

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, evmHooks).
			From(kv.MustGetStringMap(c.getter, "evm")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out evm config"))
		}

		if cfg.StartFromBlock == 0 {
			block, err := cfg.RPCClient.BlockNumber(context.TODO())
			if err != nil {
				panic(errors.Wrap(err, "failed to fetch last block"))
			}

			cfg.StartFromBlock = block
		}

		return &cfg
	}).(*Ethereum)
}

var evmHooks = figure.Hooks{
	"common.Address": func(raw interface{}) (reflect.Value, error) {
		v, err := cast.ToStringE(raw)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "expected string")
		}

		return reflect.ValueOf(common.HexToAddress(v)), nil
	},
	"*ethclient.Client": func(raw interface{}) (reflect.Value, error) {
		v, err := cast.ToStringE(raw)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "expected string")
		}

		client, err := ethclient.Dial(v)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to dial eth rpc")
		}

		return reflect.ValueOf(client), nil
	},
}

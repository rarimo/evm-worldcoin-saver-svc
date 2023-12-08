package config

import (
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/tendermint/tendermint/rpc/client/http"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func (c *config) Cosmos() *grpc.ClientConn {
	return c.cosmos.Do(func() interface{} {
		var cfg struct {
			Addr string `fig:"addr"`
		}

		if err := figure.Out(&cfg).From(kv.MustGetStringMap(c.getter, "cosmos")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out cosmos config"))
		}

		con, err := grpc.Dial(cfg.Addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second, // wait time before ping if no activity
			Timeout: 20 * time.Second, // ping timeout
		}))
		if err != nil {
			panic(errors.Wrap(err, "failed to dial cosmos grpc"))
		}

		return con
	}).(*grpc.ClientConn)
}

func (c *config) Tendermint() *http.HTTP {
	return c.tendermint.Do(func() interface{} {
		var cfg struct {
			Addr string `fig:"addr"`
		}

		if err := figure.Out(&cfg).From(kv.MustGetStringMap(c.getter, "core")).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out core config"))
		}

		client, err := http.New(cfg.Addr, "/websocket")
		if err != nil {
			panic(errors.Wrap(err, "failed to create tendermint client"))
		}

		if err := client.Start(); err != nil {
			panic(errors.Wrap(err, "failed to start tendermint client"))
		}

		return client
	}).(*http.HTTP)
}

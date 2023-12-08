package voting

import (
	"context"
	"sync"
	"time"

	"github.com/rarimo/evm-worldcoin-saver-svc/internal/config"
	rarimocore "github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/rarimo/saver-grpc-lib/voter"
	"gitlab.com/distributed_lab/running"
)

const opQueryWorldCoin = "tm.event='Tx' AND new_operation.operation_type='WORLDCOIN_IDENTITY_TRANSFER'"

func RunVoter(ctx context.Context, cfg config.Config) {
	v := voter.NewVoter(cfg.Ethereum().NetworkName, cfg.Log(), cfg.Broadcaster(), map[rarimocore.OpType]voter.Verifier{
		rarimocore.OpType_WORLDCOIN_IDENTITY_TRANSFER: NewWorldCoinVerifier(cfg),
	})

	// catchup tends to panic on startup and doesn't handle it by itself, so we wrap it into retry loop
	running.UntilSuccess(ctx, cfg.Log(), "voter-catchup", func(ctx context.Context) (bool, error) {
		voter.
			NewCatchupper(cfg.Cosmos(), v, cfg.Log()).
			Run(ctx)
		return true, nil
	}, 1*time.Second, 5*time.Second)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		voter.
			NewSubscriber(v, cfg.Tendermint(), cfg.Cosmos(), opQueryWorldCoin, cfg.Log(), cfg.Subscriber()).
			Run(ctx)
	}()

	wg.Wait()
}

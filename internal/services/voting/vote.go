package voting

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/rarimo/evm-worldcoin-saver-svc/internal/config"
	oracletypes "github.com/rarimo/rarimo-core/x/oraclemanager/types"
	rarimocore "github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/rarimo/saver-grpc-lib/voter"
	"github.com/rarimo/saver-grpc-lib/voter/verifiers"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type WorldCoinVerifier struct {
	log               *logan.Entry
	oracleQueryClient oracletypes.QueryClient
}

var _ voter.Verifier = &WorldCoinVerifier{}

func NewWorldCoinVerifier(cfg config.Config) *WorldCoinVerifier {
	return &WorldCoinVerifier{
		log:               cfg.Log(),
		oracleQueryClient: oracletypes.NewQueryClient(cfg.Cosmos()),
	}
}

func (s *WorldCoinVerifier) Verify(ctx context.Context, operation rarimocore.Operation) (rarimocore.VoteType, error) {
	if operation.OperationType != rarimocore.OpType_WORLDCOIN_IDENTITY_TRANSFER {
		s.log.Debugf("Voted NO: invalid operation type")
		return rarimocore.VoteType_NO, verifiers.ErrInvalidOperationType
	}

	var op rarimocore.WorldCoinIdentityTransfer
	if err := proto.Unmarshal(operation.Details.Value, &op); err != nil {
		s.log.Debugf("Voted NO: failed to unmarshal")
		return rarimocore.VoteType_NO, err
	}

	err := s.verify(ctx, op)
	if err == nil {
		return rarimocore.VoteType_YES, nil
	}

	s.log.WithError(err).Debugf("Voted NO: received an error from verifier")
	switch errors.Cause(err) {
	case verifiers.ErrUnsupportedNetwork:
		return rarimocore.VoteType_NO, verifiers.ErrUnsupportedNetwork
	case verifiers.ErrWrongOperationContent:
		return rarimocore.VoteType_NO, nil
	default:
		return rarimocore.VoteType_NO, err
	}
}

func (s *WorldCoinVerifier) verify(ctx context.Context, op rarimocore.WorldCoinIdentityTransfer) error {
	// TODO. Get data from Ethereum, then compare with Rarimo op (current logic is wrong)
	msg := oracletypes.MsgCreateWorldCoinIdentityTransferOp{
		Contract:  op.Contract,
		Chain:     op.Chain,
		PrevState: op.PrevState,
		State:     op.State,
		Timestamp: op.Timestamp,
	}
	// we can avoid the transfer and just use the op
	resp, err := s.oracleQueryClient.WorldCoinIdentityTransfer(ctx, &oracletypes.QueryGetWorldCoinIdentityTransferRequest{Msg: msg})
	if err != nil {
		return errors.Wrap(err, "failed to fetch operation details from core")
	}

	if !proto.Equal(&resp.Transfer, &op) {
		return verifiers.ErrWrongOperationContent
	}

	return nil
}

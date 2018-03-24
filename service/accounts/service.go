package accounts

import (
	"context"

	accountProto "github.com/homespendapi/service/accounts/proto"
	utilsProto "github.com/homespendapi/service/utils/proto"

	"google.golang.org/grpc"
)

type Service struct{}

func NewService(s *grpc.Server) {
	accountProto.RegisterAccountsServer(s, &Service{})
}

func (s *Service) CreateAccount(ctx context.Context, in *accountProto.CreateAccountRequest) (*accountProto.CreateAccountResponse, error) {
	return &accountProto.CreateAccountResponse{
		RespMsg: &utilsProto.RespMsg{
			Msg:  "Success",
			Stat: utilsProto.MsgStat_success,
		},
	}, nil
}

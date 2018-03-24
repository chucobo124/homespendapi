package budgets

import (
	"context"

	budgetProto "github.com/homespendapi/service/budgets/proto"
	utilsProto "github.com/homespendapi/service/utils/proto"

	"google.golang.org/grpc"
)

type Service struct{}

func NewService(s *grpc.Server) {
	budgetProto.RegisterBudgetsServer(s, &Service{})
}

func (s *Service) CreateBudget(ctx context.Context, in *budgetProto.CreateBudgetRequest) (*budgetProto.CreateBudgetResponse, error) {
	return &budgetProto.CreateBudgetResponse{
		RespMsg: &utilsProto.RespMsg{
			Msg:  "Success",
			Stat: utilsProto.MsgStat_success,
		},
	}, nil
}

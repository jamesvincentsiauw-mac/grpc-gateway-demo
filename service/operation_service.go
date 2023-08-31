package service

import (
	"context"
	pb "github.com/jamesvincentsiauw-mac/grpc-gateway-demo/proto/gen/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OperationService struct {
}

func count(first uint32, second uint32, operator pb.Operator) float32 {
	switch operator {
	case pb.Operator_ADDITION:
		return float32(first + second)
	case pb.Operator_SUBTRACTION:
		return float32(first - second)
	case pb.Operator_DIVISION:
		return float32(first / second)
	case pb.Operator_MULTIPLY:
		return float32(first * second)
	default:
		return 0
	}
}

func (o *OperationService) Operate(_ context.Context, request *pb.OperationRequest) (*pb.OperationResponse, error) {
	_, ok := pb.Operator_value[request.GetOperator().String()]
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "operator not allowed")
	}
	result := count(request.GetFirst(), request.GetSecond(), request.GetOperator())

	return &pb.OperationResponse{Result: result}, nil
}

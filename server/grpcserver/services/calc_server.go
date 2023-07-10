package services

import context "context"

type calcServer struct{}

var count int64

func NewCalcServer() CalcServer {
	return calcServer{}
}

func (calcServer) mustEmbedUnimplementedCalcServer() {}

func (calcServer) Plus(ctx context.Context, req *PlusRequest) (*PlusResponse, error) {

	count += int64(req.N)

	res := PlusResponse{
		Result: count,
	}

	return &res, nil
}

func (calcServer) Minus(ctx context.Context, req *MinusRequest) (*MinusResponse, error) {

	count -= int64(req.N)

	res := MinusResponse{
		Result: count,
	}

	return &res, nil
}

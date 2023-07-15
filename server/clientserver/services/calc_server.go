package services

import (
	"context"
)

type CalcService interface {
	Plus(n uint64, out *int) error
	Minus(n int64, out *int) error
	GetSummation(out *int) error
}

type calcService struct {
	calcClient CalcClient
}

func NewCalcService(calcClient CalcClient) CalcService {
	return calcService{calcClient}
}

func (serv calcService) Plus(n uint64, out *int) error {
	req := PlusRequest{
		N: n,
	}

	res, err := serv.calcClient.Plus(context.Background(), &req)
	if err != nil {
		return err
	}

	*out = int(res.Result)

	return nil
}

func (serv calcService) Minus(n int64, out *int) error {
	req := MinusRequest{
		N: n,
	}

	res, err := serv.calcClient.Minus(context.Background(), &req)
	if err != nil {
		return err
	}

	*out = int(res.Result)

	return nil
}

func (serv calcService) GetSummation(out *int) error {
	req := SumRequest{}

	res, err := serv.calcClient.GetSummation(context.Background(), &req)
	if err != nil {
		return err
	}

	*out = int(res.Result)

	return nil
}

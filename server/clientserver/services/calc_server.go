package services

import (
	"context"
	"fmt"
)

type CalcService interface {
	Plus(n uint32) error
	Minus(n int32) error
}

type calcService struct {
	calcClient CalcClient
}

func NewCalcService(calcClient CalcClient) CalcService {
	return calcService{calcClient}
}

func (serv calcService) Plus(n uint32) error {
	req := PlusRequest{
		N: n,
	}

	res, err := serv.calcClient.Plus(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func (serv calcService) Minus(n int32) error {
	req := MinusRequest{
		N: n,
	}

	res, err := serv.calcClient.Minus(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

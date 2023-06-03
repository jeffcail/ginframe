package _rpc

import (
	"errors"
)

type Arith struct{}

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

package server

import (
	calculatorPB "CalculatorApi/proto/calculator"
	"context"
	"fmt"
)

type Server struct{}

func (*Server) Add(ctx context.Context, req *calculatorPB.CalculatorReq) (*calculatorPB.CalculatorResp, error) {
	fmt.Printf("Add function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()
	z := x + y

	res := &calculatorPB.CalculatorResp{
		Expression: fmt.Sprintf("Add(%d, %d) = %d", x, y, z),
		Result:     z,
	}

	return res, nil
}
func (*Server) Sub(ctx context.Context, req *calculatorPB.CalculatorReq) (*calculatorPB.CalculatorResp, error) {
	fmt.Printf("Sub function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()
	z := x - y

	res := &calculatorPB.CalculatorResp{
		Expression: fmt.Sprintf("Sub(%d, %d) = %d", x, y, z),
		Result:     z,
	}

	return res, nil
}
func (*Server) Mult(ctx context.Context, req *calculatorPB.CalculatorReq) (*calculatorPB.CalculatorResp, error) {
	fmt.Printf("Mult function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()
	z := x * y

	res := &calculatorPB.CalculatorResp{
		Expression: fmt.Sprintf("Mult(%d, %d) = %d", x, y, z),
		Result:     z,
	}

	return res, nil
}
func (*Server) Fibonacci(
	req *calculatorPB.CalculatorReq,
	stream calculatorPB.CalculatorApi_FibonacciServer) error {
	fmt.Printf("Mult function is invoked with %v \n", req)

	var x = req.GetX()

	var chche = make([]int64, x+1)
	chche[0] = 0
	chche[1] = 1

	for i := 0; i <= int(x); i++ {
		var result = getFibonacci(i, chche)
		var res = &calculatorPB.CalculatorResp{
			Expression: fmt.Sprintf("Fibonacci(%d) = %d", i, result),
			Result:     result,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

func getFibonacci(index int, cache []int64) int64 {
	if index == 0 || index == 1 || cache[index] != 0 {
		return cache[index]
	}

	cache[index] = getFibonacci(index-1, cache) + getFibonacci(index-2, cache)
	return cache[index]
}

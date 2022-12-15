package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calculatorPB "CalculatorApi/proto/calculator"

	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) Add(ctx context.Context, req *calculatorPB.CalculatorReq) (*calculatorPB.CalculatorResp, error) {
	fmt.Printf("Add function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()
	z := x + y

	res := &calculatorPB.CalculatorResp{
		Expression: fmt.Sprintf("%d + %d = %d", x, y, z),
		Result:     z,
	}
	fmt.Printf("%v", res)

	return res, nil
}
func (*Server) Sub(ctx context.Context, req *calculatorPB.CalculatorReq) (*calculatorPB.CalculatorResp, error) {
	fmt.Printf("Sub function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()
	z := x - y

	res := &calculatorPB.CalculatorResp{
		Expression: fmt.Sprintf("%d - %d = %d", x, y, z),
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
		Expression: fmt.Sprintf("%d * %d = %d", x, y, z),
		Result:     z,
	}

	return res, nil
}

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	calculatorPB.RegisterCalculatorApiServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

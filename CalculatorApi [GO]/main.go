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
	fmt.Printf("Sum function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()

	res := &calculatorPB.CalculatorResp{
		Z: x + y,
	}

	return res, nil
}
func (*Server) Sub(ctx context.Context, req *calculatorPB.CalculatorReq) (*calculatorPB.CalculatorResp, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()

	res := &calculatorPB.CalculatorResp{
		Z: x - y,
	}

	return res, nil
}
func (*Server) Mult(ctx context.Context, req *calculatorPB.CalculatorReq) (*calculatorPB.CalculatorResp, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	x := req.GetX()
	y := req.GetY()

	res := &calculatorPB.CalculatorResp{
		Z: x * y,
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

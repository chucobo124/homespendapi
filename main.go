//go:generate protoc -I ${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/homespendapi/service/budgets/proto/budget.proto
//go:generate protoc -I ${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/homespendapi/service/utils/proto/utils.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/homespendapi/service/budgets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9200"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Setup Middleware
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(LogIntercepter),
		),
	)

	// Regest Service
	budgets.NewService(s)
	// End Regest Service

	generateServiceInfo(s)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func generateServiceInfo(s *grpc.Server) {
	for service, info := range s.GetServiceInfo() {
		fmt.Printf("GRPC Port %s \n", port)
		fmt.Println("---------------------------------------------------------------------------")
		fmt.Print("Service: ")
		fmt.Print(service + "\n")
		fmt.Println("---------------------------------------------------------------------------")
		for _, method := range info.Methods {
			fmt.Println("Grpc Methods:")
			fmt.Println("---------------------------------------------------------------------------")
			fmt.Printf("%s | IsClientStream: %t | IsServerStream: %t \n", method.Name, method.IsClientStream, method.IsServerStream)
		}
		fmt.Println("---------------------------------------------------------------------------")
	}
}

func LogIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	inTime, _ := ctx.Deadline()
	resp, err = handler(ctx, req)
	duration := time.Now().Nanosecond() - inTime.Nanosecond()
	log.Printf("GRPC ---- %s | tooks %d ns", info.FullMethod, duration)
	return
}

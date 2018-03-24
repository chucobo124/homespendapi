package main

import (
	"encoding/json"
	"log"
	"time"

	budgetsProto "github.com/homespendapi/service/budgets/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost"
	prot    = ":9200"
)

func main() {
	createBudgetReq := &budgetsProto.CreateBudgetRequest{}
	createBudget(address, prot, createBudgetReq)

}

func createBudget(address string, prot string, req *budgetsProto.CreateBudgetRequest) {
	uri := address + prot
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := budgetsProto.NewBudgetsClient(conn)
	r, err := c.CreateBudget(ctx, req)
	if err != nil {
		log.Fatalf("could not CreateBudget: %v", err)
	}
	msg, _ := json.Marshal(r.RespMsg)
	log.Printf(string(msg))

}

package main

import (
	"encoding/json"
	"log"
	"time"

	accountsProto "github.com/homespendapi/service/accounts/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost"
	prot    = ":9200"
)

func main() {
	creatAccountReq := &accountsProto.CreateAccountRequest{}
	creatAccount(address, prot, creatAccountReq)

}

func creatAccount(address string, prot string, req *accountsProto.CreateAccountRequest) {
	uri := address + prot
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := accountsProto.NewAccountsClient(conn)
	r, err := c.CreateAccount(ctx, req)
	if err != nil {
		log.Fatalf("could not CreatAccount: %v", err)
	}
	msg, _ := json.Marshal(r.RespMsg)
	log.Printf(string(msg))

}

func fetchRpc(uri string, runClient func(*grpc.ClientConn)) {

}

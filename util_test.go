package main

import (
	"context"
	"github.com/cita-cloud/executor_noop_go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

func TestReadToml(t *testing.T) {
	config, err := readToml("example/test-chain-0/config.toml", "executor_evm")
	if err != nil {
		t.Fatal(err)
	}
	if config.ExecutorPort != 50002 {
		t.Errorf("config.ExecutorPort is %d, want 5002", config.ExecutorPort)
	}
}

func TestGrpcServer(t *testing.T) {
	serverAddr := "localhost:50002"
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewExecutorServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	header := new(pb.BlockHeader)
	header.Height = 100
	req := new(pb.Block)
	req.Header = header
	req.Header.Height = 100

	res, err := client.Exec(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if res.Status.Code != 0 {
		t.Errorf("res.Status.Code is %d, want 200", res.Status.Code)
	}
}

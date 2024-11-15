package main

import (
	"context"
	"github.com/cita-cloud/executor_noop_go/pb"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

var (
	runCommand = &cli.Command{
		Name:   "run",
		Usage:  "run executor server",
		Action: runServer,
		Flags: []cli.Flag{
			configFileFlag,
			privateFileFlag,
		},
	}
	configFileFlag = &cli.PathFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "executor server config file",
		DefaultText: "config.toml",
		Value:       "config.toml",
		TakesFile:   true,
	}
	privateFileFlag = &cli.PathFlag{
		Name:        "private_key_path",
		Aliases:     []string{"p"},
		Usage:       "executor server private key path",
		DefaultText: "private_key",
		Value:       "private_key",
		TakesFile:   true,
	}
)

type ExecutorServiceServerImpl struct {
	pb.UnimplementedExecutorServiceServer
}

func (*ExecutorServiceServerImpl) Exec(ctx context.Context, block *pb.Block) (*pb.HashResponse, error) {
	log.Printf("get block height: %v", block.Header.Height)
	rawBytes, err := proto.Marshal(block)
	if err != nil {
		return &pb.HashResponse{
			Status: &pb.StatusCode{
				Code: 122,
			},
			Hash: &pb.Hash{
				Hash: make([]byte, 32),
			},
		}, nil
	}
	hash := sm3.Sm3Sum(rawBytes)
	return &pb.HashResponse{
		Status: &pb.StatusCode{
			Code: 0,
		},
		Hash: &pb.Hash{
			Hash: hash,
		},
	}, nil
}

func (*ExecutorServiceServerImpl) Call(ctx context.Context, request *pb.CallRequest) (*pb.CallResponse, error) {
	return &pb.CallResponse{
		Value: nil,
	}, nil
}

type RPCServiceServerImpl struct {
	pb.UnimplementedRPCServiceServer
}

func (*RPCServiceServerImpl) GetTransactionReceipt(ctx context.Context, hash *pb.Hash) (*pb.Receipt, error) {
	//TODO implement me
	panic("implement me")
}

func (*RPCServiceServerImpl) GetCode(ctx context.Context, request *pb.GetCodeRequest) (*pb.ByteCode, error) {
	//TODO implement me
	panic("implement me")
}

func (*RPCServiceServerImpl) GetBalance(ctx context.Context, request *pb.GetBalanceRequest) (*pb.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (*RPCServiceServerImpl) GetTransactionCount(ctx context.Context, request *pb.GetTransactionCountRequest) (*pb.Nonce, error) {
	//TODO implement me
	panic("implement me")
}

func (*RPCServiceServerImpl) GetAbi(ctx context.Context, request *pb.GetAbiRequest) (*pb.ByteAbi, error) {
	//TODO implement me
	panic("implement me")
}

func (*RPCServiceServerImpl) EstimateQuota(ctx context.Context, request *pb.CallRequest) (*pb.ByteQuota, error) {
	//TODO implement me
	panic("implement me")
}

func (*RPCServiceServerImpl) GetReceiptProof(ctx context.Context, hash *pb.Hash) (*pb.ReceiptProof, error) {
	//TODO implement me
	panic("implement me")
}

func (*RPCServiceServerImpl) GetRootsInfo(ctx context.Context, number *pb.BlockNumber) (*pb.RootsInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (R *RPCServiceServerImpl) GetStorageAt(ctx context.Context, request *pb.GetStorageAtRequest) (*pb.Hash, error) {
	//TODO implement me
	panic("implement me")
}

type HealthServerImpl struct {
	pb.UnimplementedHealthServer
}

func (*HealthServerImpl) Check(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("healthcheck entry!")
	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_SERVING}, nil
}

func runServer(ctx *cli.Context) error {
	path := ctx.Path(configFileFlag.Name)
	config, err := readToml(path, "executor_evm")
	if err != nil {
		return err
	}
	logPath := config.LogConfig.RollingFilePath + "/" + config.LogConfig.ServiceName + ".log"
	dir := filepath.Dir(logPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create directory: %v\n", err)
	}
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	lis, err := net.Listen("tcp", ":"+strconv.FormatUint(uint64(config.ExecutorPort), 10))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterExecutorServiceServer(grpcServer, &ExecutorServiceServerImpl{})
	pb.RegisterRPCServiceServer(grpcServer, &RPCServiceServerImpl{})
	pb.RegisterHealthServer(grpcServer, &HealthServerImpl{})

	log.Printf("Starting executor server on port %v...\n", config.ExecutorPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
	return nil
}

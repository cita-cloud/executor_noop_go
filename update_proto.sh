git submodule deinit --all -f
git submodule update --force --init --remote --recursive
rm -rf ./pb/*
protoc --proto_path=cloud-proto/protos --go_out=. --go-grpc_out=. --go_opt=Mblockchain.proto=./pb --go_opt=Mcommon.proto=./pb --go_opt=Mexecutor.proto=./pb --go_opt=Mvm/evm.proto=./pb --go-grpc_opt=Mblockchain.proto=./pb --go-grpc_opt=Mcommon.proto=./pb --go-grpc_opt=Mexecutor.proto=./pb --go-grpc_opt=Mvm/evm.proto=./pb cloud-proto/protos/blockchain.proto cloud-proto/protos/common.proto cloud-proto/protos/executor.proto cloud-proto/protos/vm/evm.proto
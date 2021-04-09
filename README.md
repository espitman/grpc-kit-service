# gRPC Kit Service

1. remove go.mod & go.sum
2. go mod init ```service name```
3. replace all ```grpc-kit-service``` with ```service name```
4. replace ```github.com/espitman/protos-kit/kit``` with new proto repo address
5. rename ```kit.RegisterKitServiceServer``` in ```cmd/start.go```

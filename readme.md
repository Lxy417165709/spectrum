
## protoc To grpc go
```shell
export PATH=$PATH:/Users/lixueyue/go/bin 
protoc  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative common/pb/*.proto


protoc --go_out=plugins=grpc:. common/pb/*.proto
```

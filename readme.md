
## protoc To grpc go
```shell
export PATH=$PATH:/Users/lixueyue/go/bin 
protoc  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative common/pb/*.proto


protoc --go_out=plugins=grpc:. common/pb/*.proto
```


## 注意点
2021-03-19 23:08:09
good表与建表sql语句定义不同！！！记得修改
还有 desk 表

order 表也会变
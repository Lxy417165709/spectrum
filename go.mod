module test

go 1.14

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.32.0

require (
	github.com/astaxie/beego v1.12.2
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb // indirect
	golang.org/x/sys v0.0.0-20201013132646-2da7054afaeb // indirect
	google.golang.org/genproto v0.0.0-20201013134114-7f9ee70cb474 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
)

module github.com/dreamstax/kai-piper

require (
	cloud.google.com/go/longrunning v0.5.1
	github.com/dreamstax/go-genproto v0.0.0-20231102092056-7aab8b064f6d
	github.com/golang/protobuf v1.5.3
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)

// replace github.com/dreamstax/go-genproto => ../go-genproto

go 1.20

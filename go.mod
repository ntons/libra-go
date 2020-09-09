module github.com/ntons/libra-go

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/ntons/log-go v0.0.0-20200814034523-0ad12f491b2b
	github.com/ntons/tongo v0.0.0-20200811072755-8cc8b86c2a60
	go.uber.org/zap v1.15.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200813231717-0a73ddcff9b8 // indirect
	google.golang.org/genproto v0.0.0-20200808173500-a06252235341
	google.golang.org/grpc v1.31.0
	google.golang.org/grpc/examples v0.0.0-20200810225334-2983360ff4e7
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
)

replace github.com/ntons/tongo => /root/ntons/tongo

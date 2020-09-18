module github.com/ntons/libra-go

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/ntons/log-go v0.0.0-20200814034523-0ad12f491b2b
	github.com/ntons/tongo/template v0.0.0-20200915100842-af1058d95ac9
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20200904194848-62affa334b73
	golang.org/x/sys v0.0.0-20200915084602-288bc346aa39 // indirect
	golang.org/x/tools v0.0.0-20200813231717-0a73ddcff9b8 // indirect
	google.golang.org/genproto v0.0.0-20200915202801-9f80d0600517
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

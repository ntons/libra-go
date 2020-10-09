module github.com/ntons/libra-go/protos

go 1.15

// go.mod here is used to exclude protos from the top-level Go module.
// We need pin grpc@v1.29.1, but google.golang.org/api imported by googleapis
// require grpc@v1.31.x

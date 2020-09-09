package sdk

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type healthServer struct {
	grpc_health_v1.UnimplementedHealthServer
	ctx context.Context
}

func newHealth(ctx context.Context) *healthServer {
	return &healthServer{ctx: ctx}
}

func (h *healthServer) Check(
	ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (
	resp *grpc_health_v1.HealthCheckResponse, err error) {
	//log.Debugf("health check: %v", req)
	var status = grpc_health_v1.HealthCheckResponse_UNKNOWN
	select {
	case <-h.ctx.Done():
		status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	default:
		status = grpc_health_v1.HealthCheckResponse_SERVING
	}
	resp = &grpc_health_v1.HealthCheckResponse{Status: status}
	return
}

func (h *healthServer) Watch(
	req *grpc_health_v1.HealthCheckRequest,
	stream grpc_health_v1.Health_WatchServer) (err error) {
	//log.Debugf("health watch: %v", req)
	var status = grpc_health_v1.HealthCheckResponse_UNKNOWN
	select {
	case <-h.ctx.Done():
		status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
		stream.Send(&grpc_health_v1.HealthCheckResponse{Status: status})
	default:
		status = grpc_health_v1.HealthCheckResponse_SERVING
		stream.Send(&grpc_health_v1.HealthCheckResponse{Status: status})
	}
	select {
	case <-h.ctx.Done():
		status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
		stream.Send(&grpc_health_v1.HealthCheckResponse{Status: status})
	case <-stream.Context().Done():
	}
	return
}

package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validatable interface {
	Validate() error
}

//Validate request validation middleware
func Validate(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if v, ok := req.(validatable); ok {
		err = v.Validate()
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "request validation failed: %s", err)
		}
	}
	resp, err = handler(ctx, req)
	return
}

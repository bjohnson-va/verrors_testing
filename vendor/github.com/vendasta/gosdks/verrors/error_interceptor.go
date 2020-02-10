package verrors

import (
	"context"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

var (
	// DefaultErrorMask is a configuration mask for the ErrorConverterServerInterceptor
	// Masks all errors except:
	// NotFound
	// InvalidArgument
	// FailedPrecondition
	//
	// You can delete or edit masks to build your own configuration:
	//
	// delete(verrors.DefaultErrorMask, verrors.AlreadyExists)
	// verrors.DefaultErrorMask[util.Unimplemented] = "🚧 Under construction 🚧"
	DefaultErrorMask = map[ErrorType]string{
		Unknown:           "Something went wrong",
		Internal:          "Internal Server Error",
		PermissionDenied:  "Permission Denied",
		ResourceExhausted: "Resource Exhausted, please try again later",
		DeadlineExceeded:  "Deadline Exceeded",
		Unimplemented:     "Unimplemented",
		Unauthenticated:   "Unauthenticated",
		AlreadyExists:     "Entity already exists",
		Aborted:           "Request Aborted",
		Unavailable:       "Service Unavailable",
		Canceled:          "Request was canceled",
	}
)

// ErrorConverterServerInterceptor converts service errors and simple error types into grpc errors
// This removes the need to convert all your errors into grpc errors in the application layer.
// This should be placed BEFORE your logging interceptor
//
// maskErrorsMap is a map of ErrorType's to Error messages, any overrides present in the map will be used when converting
// errors of that type into GRPC errors.
// See DefaultErrorMask for an example configuration
// Usage:
// grpcServer := serverconfig.CreateGrpcServer(verrors.ErrorConverterServerInterceptor(DefaultErrorMask))
func ErrorConverterServerInterceptor(maskErrorsMap map[ErrorType]string,
) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			_, ok := status.FromError(err)
			if ok {
				return nil, err
			}
			serviceError := FromErrorWithContext(ctx, err)
			messageOverride, ok := maskErrorsMap[serviceError.errType]
			if !ok {
				// if we don't have an override just convert the error and move on:
				return nil, serviceError.GRPCError()
			}
			return nil, New(serviceError.errType, messageOverride).GRPCError()
		}
		return resp, nil
	}
}

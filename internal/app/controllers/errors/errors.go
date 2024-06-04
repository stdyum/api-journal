package errors

import (
	grpcErr "github.com/stdyum/api-common/grpc"
	httpErr "github.com/stdyum/api-common/http"
)

var (
	HttpErrorsMap = map[error]any{}

	GRpcErrorsMap = map[error]any{}
)

func Register() {
	httpErr.RegisterErrors(HttpErrorsMap)
	grpcErr.RegisterErrors(GRpcErrorsMap)
}

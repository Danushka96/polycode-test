package greeting_service

import (
	"fmt"
	"github.com/cloudimpl/polycode-runtime/go/sdk"
	"pos-app/models"
)

func Hello(ctx sdk.ServiceContext, req models.HelloRequest) (models.HelloResponse, error) {
	return models.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}

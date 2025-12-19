package greeting_service

import (
	"fmt"
	"pos-app/models"

	"github.com/cloudimpl/polycode-runtime/go/sdk"
)

func Hello(ctx sdk.ServiceContext, req models.HelloRequest) (models.HelloResponse, error) {
	ctx.Logger().Info().Msg("Hello request test log message")
	return models.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}

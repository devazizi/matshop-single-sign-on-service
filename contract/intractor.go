package contract

import (
	"context"
	"sso/dto"
)

type Intractor interface {
	RegisterClient(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
}

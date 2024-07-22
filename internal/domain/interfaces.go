package domain

import (
	"context"
	"example/template/internal/domain/dto"
	"example/template/internal/domain/entity"
)

type (
	UserRepo interface {
		Create(ctx context.Context, d dto.CreateUser) (*entity.User, error)
		Find(ctx context.Context, filter dto.UserFilter) (*entity.User, error)
	}
)

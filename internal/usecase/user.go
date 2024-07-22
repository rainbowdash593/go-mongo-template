package usecase

import (
	"context"
	"example/template/internal/domain"
	"example/template/internal/domain/dto"
	"example/template/internal/domain/entity"
)

type UserUseCase struct {
	repo domain.UserRepo
}

func NewUserService(r domain.UserRepo) UserUseCase {
	return UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, d dto.CreateUser) (*entity.User, error) {
	u, err := uc.repo.Create(ctx, d)
	return u, err
}

func (uc *UserUseCase) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	u, err := uc.repo.Find(ctx, dto.UserFilter{ID: id})
	return u, err
}

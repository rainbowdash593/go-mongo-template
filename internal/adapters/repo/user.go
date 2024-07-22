package repo

import (
	"context"
	"errors"
	"example/template/internal/adapters/repo/models"
	"example/template/internal/domain/dto"
	"example/template/internal/domain/entity"
	"example/template/internal/domain/exceptions"
	"example/template/pkg/mongodb"
	"example/template/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db *mongodb.Database
	c  *mongo.Collection
}

func NewUserRepo(db *mongodb.Database) *UserRepo {
	return &UserRepo{db: db, c: db.DB.Collection("users")}
}

func (r UserRepo) Find(ctx context.Context, dto dto.UserFilter) (*entity.User, error) {
	var (
		user models.User
		err  error
	)
	filter := bson.M{}

	objID, _ := primitive.ObjectIDFromHex(dto.ID)
	if !objID.IsZero() {
		filter["_id"] = objID
	}

	if dto.Email != "" {
		filter["email"] = dto.Email
	}

	if dto.Name != "" {
		filter["name"] = dto.Name
	}
	err = r.c.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = exceptions.ErrUserNotFound
		} else {
			err = utils.WrapErrors(exceptions.ErrUnhandled, err)
		}
		return nil, err
	}

	return &entity.User{
		Id:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r UserRepo) Create(ctx context.Context, d dto.CreateUser) (*entity.User, error) {
	existingUser, err := r.Find(ctx, dto.UserFilter{Email: d.Email})

	if err != nil && !errors.Is(err, exceptions.ErrUserNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return nil, exceptions.ErrUserAlreadyExists
	}

	result, err := r.c.InsertOne(ctx, d)
	insertedID := result.InsertedID.(primitive.ObjectID)

	if err != nil {
		return nil, utils.WrapErrors(exceptions.ErrUnhandled, err)
	}

	return r.Find(ctx, dto.UserFilter{ID: insertedID.Hex()})
}

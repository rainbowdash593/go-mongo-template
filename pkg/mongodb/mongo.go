package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	_connectTimeout = 10 * time.Second
)

type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
	ctx    context.Context
}

func New(dsn string, database string) (*Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), _connectTimeout)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Database{Client: client, DB: client.Database(database), ctx: ctx}, nil
}

func (d Database) Close() error {
	err := d.Client.Disconnect(d.ctx)
	return err
}

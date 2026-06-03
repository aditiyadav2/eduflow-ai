package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("users"),
	}
}

func (r *Repository) CreateUser(ctx context.Context, user User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

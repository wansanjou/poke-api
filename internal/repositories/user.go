package repositories

import (
	"context"
	"time"

	"github.com/wansanjou/poke-api/internal/core/domains"
	"github.com/wansanjou/poke-api/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	mc  *mongo.Client
	db  string
	col string
}

func NewUserRepository(mc *mongo.Client, db string) ports.UserRepository {
	col := "users"
	_, err := mc.Database(db).Collection(col).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
	return &userRepository{mc, db, col}
}

func (u *userRepository) Create(ctx context.Context, data domains.User) (*domains.User, error) {
	return u.insertOne(ctx, data)
}

func (u *userRepository) FindByUsername(ctx context.Context, username string) (*domains.User, error) {
	filter := bson.D{{Key: "username", Value: username}}
	return u.findOne(ctx, filter)
}

func (u *userRepository) findOne(ctx context.Context, filter bson.D) (*domains.User, error) {
	out := domains.User{}
	col := u.mc.Database(u.db).Collection(u.col)
	if err := col.FindOne(ctx, filter).Decode(&out); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &out, nil
}

func (u *userRepository) insertOne(ctx context.Context, in domains.User) (*domains.User, error) {
	in.CreatedAt = time.Now().UTC()
	col := u.mc.Database(u.db).Collection(u.col)
	result, err := col.InsertOne(ctx, in)
	if err != nil {
		return nil, err
	}
	oid, _ := result.InsertedID.(primitive.ObjectID)
	in.ID = oid
	return &in, nil
}

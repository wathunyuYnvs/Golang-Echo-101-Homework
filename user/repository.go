package user

import (
	"context"
	"myecho/db"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	resource   *db.Resource
	collection *mongo.Collection
}

type Repository interface {
	GetAll() (Users, error)
	CreateOne(User) (User, error)
	GetByID(id string) (*User, error)
	FindLogin(username string) (*User, error)
	UpdateOne(string, User) (User, error)
}

func NewUserRepository(resource *db.Resource) Repository {
	collection := resource.DB.Collection("user")
	repository := &UserRepository{resource: resource, collection: collection}
	return repository
}

// GetAll to get all users
func (ur *UserRepository) GetAll() (Users, error) {
	users := Users{}
	ctx, cancel := initContext()
	defer cancel()

	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return Users{}, err
	}

	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			logrus.Print(err)
		}
		users = append(users, user)
	}
	return users, nil
}

// GetByID to get user by id
func (ur *UserRepository) GetByID(id string) (*User, error) {
	var user User

	ctx, cancel := initContext()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	err := ur.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateOne to create new user
func (ur *UserRepository) CreateOne(u User) (User, error) {
	ctx, cancel := initContext()
	defer cancel()
	_, err := ur.collection.InsertOne(ctx, u)

	if err != nil {
		return User{}, err
	}

	return u, nil
}

// FindLogin to get user by id
func (ur *UserRepository) FindLogin(username string) (*User, error) {
	var user User

	ctx, cancel := initContext()
	defer cancel()

	err := ur.collection.FindOne(ctx, bson.M{"email": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateOne to create new user
func (ur *UserRepository) UpdateOne(id string, u User) (User, error) {
	ctx, cancel := initContext()
	defer cancel()
	_id, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return User{}, e
	}
	update := bson.M{
		"$set": bson.M{
			"name":     u.Name,
			"age":      u.Age,
			"password": u.Password,
		},
	}
	filter := bson.D{primitive.E{Key: "_id", Value: _id}}
	if _, err := ur.collection.UpdateOne(ctx, filter, update); err != nil {
		return User{}, err
	}

	return u, nil
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

package db

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s"
)

type Resource struct {
	DB *mongo.Database
}

func (r *Resource) Close() {
	logrus.Warning("Closing all db connections")
}

func Init() (*Resource, error) {
	// 	username := os.Getenv("MONGODB_USERNAME")
	// 	password := os.Getenv("MONGODB_PASSWORD")
	// 	dbName := os.Getenv("MONGODB_DB_NAME")
	clusterEndpoint := "localhost" //os.Getenv("MONGODB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, clusterEndpoint)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		logrus.Errorf("Failed to create client: %v", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Errorf("Failed to connect to server: %v", err)
		return nil, err
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Errorf("Failed to ping cluster: %v", err)
		return nil, err
	}

	return &Resource{DB: client.Database("myecho")}, nil
}

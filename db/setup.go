package db

import (
	"context"
	"digicert/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DBClient *mongo.Client
)

func Connect(ctx context.Context, config config.Config) (*mongo.Database, error) {
	//var uri = fmt.Sprintf("%s://%s:%s@%s/%s", config.DBDriver, config.DBUsername, config.DBPassword, config.DBHost, config.DBName)
	var uri = fmt.Sprintf("%s://%s:%s", config.DBDriver, config.DBHost, config.DBPort)

	var clientOptions = options.Client().ApplyURI(uri)

	DBClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	var ping = DBClient.Ping(context.Background(), readpref.Primary())

	if ping != nil {
		log.Fatal(ping.Error())
	} else {
		log.Info(" MongoDb [Connected].")
	}

	return DBClient.Database(config.DBName), nil
}

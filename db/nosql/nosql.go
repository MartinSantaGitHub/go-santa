package nosql

import (
	"context"
	"log"
	"os"

	"helpers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbNoSql struct {
	Connection *mongo.Client
}

// region "Connection"

/* Connect connects to the database */
func (db *DbNoSql) Connect() error {
	connTimeout := os.Getenv("CTX_TIMEOUT")
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONN"))
	ctx, cancel := helpers.GetTimeoutCtx(connTimeout)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return err
	}

	db.Connection = client

	return nil
}

/* IsConnection makes a ping to the Database */
func (db *DbNoSql) IsConnection() bool {
	err := db.Connection.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err.Error())

		return false
	}

	return true
}

// endregion

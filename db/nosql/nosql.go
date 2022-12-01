package nosql

import (
	"context"
	"errors"
	"log"
	"os"

	"helpers"
	m "models/nosql"

	"go.mongodb.org/mongo-driver/bson"
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

// region "Greetings"

/* SaveName saves a name to the file */
func (db *DbNoSql) SaveName(name string) (bool, error) {
	var user m.User

	col := getCollection(db, "greetings", "users")
	user, isFound, err := getUser(col, name)

	if !isFound {
		ctxInsert, cancelInsert := helpers.GetTimeoutCtx(os.Getenv("CTX_TIMEOUT"))

		defer cancelInsert()

		user = m.User{
			Name:   name,
			Active: true,
		}

		_, err = col.InsertOne(ctxInsert, user)

		if err != nil {
			return false, err
		}
	}

	if !user.Active {
		updateString := bson.M{
			"$set": bson.M{"active": true},
		}

		ctxUpdate, cancelUpdate := helpers.GetTimeoutCtx(os.Getenv("CTX_TIMEOUT"))

		defer cancelUpdate()

		_, err = col.UpdateByID(ctxUpdate, user.Id, updateString)
	}

	return isFound && user.Active, err
}

/* GetNames gets the names saved in the file */
func (db *DbNoSql) GetNames() ([]string, error) {
	col := getCollection(db, "greetings", "users")

	ctxCount, cancelCount := helpers.GetTimeoutCtx(os.Getenv("CTX_TIMEOUT"))

	defer cancelCount()

	total, err := col.CountDocuments(ctxCount, bson.D{})

	if err != nil || total == 0 {
		return nil, err
	}

	ctxFind, cancelFind := helpers.GetTimeoutCtx(os.Getenv("CTX_TIMEOUT"))

	defer cancelFind()

	cur, err := col.Find(ctxFind, bson.D{})

	if err != nil {
		return nil, err
	}

	user := m.User{}
	names := make([]string, total)
	ctxCursor := context.TODO()
	idx := 0

	defer cur.Close(ctxCursor)

	for cur.Next(ctxCursor) {
		err := cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		names[idx] = user.Name

		idx++
	}

	return names, nil
}

// endregion

// region "Helpers"

func getCollection(db *DbNoSql, dbName string, colName string) *mongo.Collection {
	database := db.Connection.Database(dbName)
	collection := database.Collection(colName)

	return collection
}

func getUser(col *mongo.Collection, name string) (m.User, bool, error) {
	var user m.User

	filter := bson.M{
		"name": name,
	}

	ctxFind, cancelFind := helpers.GetTimeoutCtx(os.Getenv("CTX_TIMEOUT"))

	defer cancelFind()

	err := col.FindOne(ctxFind, filter).Decode(&user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return user, false, nil
	}

	if err != nil {
		return user, false, err
	}

	return user, true, nil
}

// endregion

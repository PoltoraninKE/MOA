package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bankAccount struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string
	Categories []interface{}
	Cash       int
}

const mongoDbConntectionString = "mongodb://localhost:27017/"

func setupRouter() *gin.Engine {
	router := gin.Default()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbConntectionString))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("MOA_staging").Collection("bankAccount")

	router.GET("", func(cont *gin.Context) {
		cont.String(http.StatusOK, "pong")
	})

	router.GET("/bankAccount/:name", func(cont *gin.Context) {
		userName := cont.Params.ByName("name")
		var result bankAccount

		err := coll.FindOne(context.TODO(), bson.D{{Key: "Name", Value: userName}}).Decode(&result)

		if err == mongo.ErrNoDocuments {
			fmt.Printf("User with name %s\n not found", userName)
			return
		}
		if err != nil {
			panic(err)
		}

		fmt.Print(result)
	})

	return router
}

func main() {
	router := setupRouter()

	router.Run(":80")
}

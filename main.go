package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vyash/controllers"
	"github.com/vyash/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO() //we are not dealing with any timeouts , TODO will create a simple context object with no cancellation thing inside
	mongoconn := options.Client().ApplyURI("mongodb+srv://<>:<>@cluster0.tis34.mongodb.net/thepolyglotdeveloper?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error")
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection establish")

	usercollection = (*mongo.Collection)(mongoclient.Database("thepolyglotdeveloper").Collection("microserviceusers"))
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}
func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/api/v1")
	usercontroller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9000"))
}

//controller -> interact with service -> service interact with database
//decode =unmarshal= json to struct return byte slice
// we use ctx because can can the process if it takes ,ore thsan 10 sec or 15 sec whatever

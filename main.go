package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BogdanJas/Hospital/controlers"
	"github.com/BogdanJas/Hospital/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
var (
	server 			*gin.Engine
	patientservice 	 services.PatientService
	patientcontrolers 	 controlers.PatientController
	ctx 			 context.Context
	patientcollection 	 *mongo.Collection
	mongoclient 	 *mongo.Client
	err 			 error
)

func init(){
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx,mongoconn)
	if err!= nil{
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx,readpref.Primary())
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")
	patientcollection = mongoclient.Database("patientdb").Collection("patients")
	patientservice = services.NewPatientService(patientcollection,ctx)
	patientcontrolers = controlers.New(patientservice)
	server = gin.Default()

}

func main(){
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/hospital")
	patientcontrolers.RegisterpatientRoutes(basepath)

	log.Fatal(server.Run(":8080"))
}
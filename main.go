package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spur-dev/api/controllers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dgraph-io/badger"
	"github.com/julienschmidt/httprouter"
)

var SC controllers.SessionController
var RegionName = goDotEnvVariable("REGION")
var RawBucketName = goDotEnvVariable("RAW_VIDEOS_BUCKET")
var FinalBucketName = goDotEnvVariable("FINAL_VIDEOS_BUCKET")
var PORT = goDotEnvVariable("PORT")

func connectDynamoDB() *dynamodb.DynamoDB {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: &RegionName,
	})))
}

func openBadgerCache() *badger.DB {
	opts := badger.DefaultOptions(goDotEnvVariable("CACHE_LOCATION"))
	opts.Logger = nil // Removing badger startup and shutdown logs
	cache, err := badger.Open(opts)

	if err != nil {
		log.Fatalln(err)
		panic("Could not start cache")
	}

	return cache
}

func connectS3() *s3.S3 {
	return s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(RegionName),
	})))
}

func init() {
	SC = *controllers.NewSessionController(connectDynamoDB(), openBadgerCache(), connectS3())

	// Creating dynamo db table
	err := SC.CreateVideosTable()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}

	// TODO: Create videos bucket if doesnot exist
}

func main() {
	router := httprouter.New()

	router.GET("/new-video/:uid", SC.NewVideoHandler)
	router.GET("/video/:vid", SC.GetVideoMetadaHandler)
	router.DELETE("/video/:vid", SC.CancelRecordingHandler)
	router.POST("/update-video-state/:vid", SC.UpdateVideoStateHandler)

	fmt.Printf("Starting http server on %s \n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}

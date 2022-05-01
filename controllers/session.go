package controllers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dgraph-io/badger"

	"github.com/google/uuid"
)

var TableName = "Videos"

func GenerateUniqueVideoId(uid string) string {
	id := uuid.NewString()
	return id
}

type SessionController struct {
	db    *dynamodb.DynamoDB
	cache *badger.DB
	s3    *s3.S3
}

func NewSessionController(db *dynamodb.DynamoDB, cache *badger.DB, s3 *s3.S3) *SessionController {
	return &SessionController{db, cache, s3}
}

func (sc SessionController) CreateVideosTable() error {
	_, err := sc.db.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("VID"),
				AttributeType: aws.String("S"),
			},
			// {
			// 	AttributeName: aws.String("Timestamp"),
			// 	AttributeType: aws.String("N"),
			// },
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("VID"),
				KeyType:       aws.String("HASH"),
			},
			// {
			// 	AttributeName: aws.String("Timestamp"),
			// 	KeyType:       aws.String("RANGE"),
			// },
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   &TableName,
	})

	return err
}

// func (sc SessionController) CreateRawVideosBuckets() error {
// func (sc SessionController) CreateFinalVideosBuckets() error {

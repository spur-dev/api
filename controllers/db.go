package controllers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/spur-dev/api/models"
)

func (sc SessionController) GetVideoMetadata(vid string) (models.MetaData, error) {
	result, err := sc.db.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"VID": {
				S: aws.String(vid),
			},
		},

		TableName: aws.String(TableName),
	})

	if err != nil {
		// log.Fatalln(err)
		fmt.Println("Expecting error here")
		fmt.Println(err)
		return models.MetaData{}, err
	}

	v := models.MetaData{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &v)

	return v, err
}

func (sc SessionController) CreateVideoMetadata(v models.MetaData) error {
	// TODO: check if entry exists
	fmt.Println("Putting video metadata")
	fmt.Println(v)
	_, err := sc.db.PutItem(&dynamodb.PutItemInput{ // This will override existing entry
		Item: map[string]*dynamodb.AttributeValue{
			"VID": {
				S: aws.String(v.VID),
			},
			"Timestamp": {
				N: aws.String(strconv.Itoa(v.Timestamp)),
			},
			"UID": {
				S: aws.String(v.UID),
			},
			"State": {
				S: aws.String(v.State),
			},
		},

		TableName: &TableName,
	})

	return err
}

func (sc SessionController) UpdateVideoMetadataStatus(vid string, s string) error {
	// TODO: check if entry exists
	_, err := sc.db.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#S": aws.String("Status"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Status": {
				S: aws.String(s),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"VID": {
				S: aws.String(vid),
			},
		},
		TableName: &TableName,

		UpdateExpression: aws.String("SET #S = :Status"),
	})

	return err
}

func (sc SessionController) DeleteVideoMetadata(id string) error {
	// TODO: This needs to do further processing to also delete from raw bucket and final bucket

	_, err := sc.db.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"VID": {
				S: aws.String(id),
			},
		},

		TableName: &TableName,
	})

	return err

}

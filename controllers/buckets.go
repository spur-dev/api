package controllers

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// TODO : Error handling
func (sc SessionController) DeleteVideoFromBucket(videoName string, bucketName string) {

	// func deleteObject(filename string) (resp *s3.DeleteObjectOutput) {
	fmt.Println("Deleting: ", videoName)
	_, err := sc.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(videoName),
	})

	if err != nil {
		fmt.Printf("Error when deleting %s from bucket %s", videoName, bucketName)
		fmt.Println("I really don't know anything about this error")
		log.Fatal(err)
	}

	// return resp
	//   }
	// return nil
}

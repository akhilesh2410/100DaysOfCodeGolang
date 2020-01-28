package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToS3(filename string) error {
	//Create env credential here which will fetch access and secret key from environment variable
	//make sure we have set environment variable with accesskey and secretkey
	//set AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY for access key
	//set AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY for secret key
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsEast2RegionID),
		Credentials: credentials.NewEnvCredentials(),
	}))
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	//test with 64 MB chunk size
	uploader.PartSize = 1024 * 1024 * 64
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	keyName := filepath.Base(filename)
	fmt.Println("uploading file to s3 : ", keyName)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("akhileshs3"),
		Key:    aws.String(keyName),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil
}

func main() {
	err := UploadFileToS3("/Users/akhileshverma/test.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR in uploading image")
		return
	}
	fmt.Println("Successfully uploaded")

}

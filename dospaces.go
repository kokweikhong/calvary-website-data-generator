package main

import (

	// Additional imports needed for examples below

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListObjects(directory string) ([]string, error) {
	key := "DO00YV7AJWYPTFYTJM32"
	secret := "+8ncKmZNIrP0jRbUwjl5PGzxejERasB3CLIhQJ2WrWI"

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String("https://sgp1.digitaloceanspaces.com"),
		Region:           aws.String("sgp1"),
		S3ForcePathStyle: aws.Bool(false), // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(newSession)

	// list all objects in a bucket
	bucket := "calvarycarpentry-cloud-storage"
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(1000),
		Prefix:  aws.String(directory),

		// Prefix: aws.String(""),
		// Delimiter: aws.String(""),
		// ContinuationToken: aws.String(""),
		// FetchOwner: aws.Bool(true),
		// StartAfter: aws.String(""),
	}

	result, err := s3Client.ListObjectsV2(input)
	if err != nil {
		return nil, err
	}

	var keys []string

	for _, object := range result.Contents {
		// check if object is a directory
		if *object.Size == 0 {
			continue
		}
		keys = append(keys, *object.Key)
	}

	return keys, nil
}

package cfs3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	accountID    = ""
	accessKeyID  = ""
	accessSecret = ""
	bucketName   = ""
)

type R2Client struct {
	client *s3.Client
}

func NewTest() *R2Client {
	return NewR2Client(bucketName, accountID, accessKeyID, accessSecret)
}

func NewR2Client(bucketName, accountId, accessKeyId, accessKeySecret string) *R2Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
	})
	return &R2Client{
		client: client,
	}
}

// ListObjectsV2
func (c *R2Client) ListObjectsV2(ctx context.Context, bucketName string, prefix string) (*s3.ListObjectsV2Output, error) {
	listObjectsOutput, err := c.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range listObjectsOutput.Contents {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}
	listBucketsOutput, err := c.client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range listBucketsOutput.Buckets {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}

	return listObjectsOutput, nil
}

// putObject
func (c *R2Client) PutObject(ctx context.Context, bucketName string, key string, body io.Reader) (*s3.PutObjectOutput, error) {
	putObjectOutput, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		log.Fatal(err)
	}
	return putObjectOutput, nil
}

// getObject
func (c *R2Client) GetObject(ctx context.Context, bucketName string, key string) (*s3.GetObjectOutput, error) {
	getObjectOutput, err := c.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal(err)
	}
	return getObjectOutput, nil
}

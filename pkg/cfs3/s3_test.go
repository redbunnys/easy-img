package cfs3_test

import (
	"bytes"
	"context"
	s3 "navpage/pkg/cfs3"
	"strings"
	"testing"
)

func TestListObjectsV2(t *testing.T) {
	client := s3.NewTest()
	client.ListObjectsV2(context.Background(), "easy-img", "")
}

// putObject
func TestPutObject(t *testing.T) {
	client := s3.NewTest()
	client.PutObject(context.Background(), "easy-img", "test/test.txt", strings.NewReader("Hello R2 PutObject!"))
}

// getObject
func TestGetObject(t *testing.T) {
	client := s3.NewTest()
	data, _ := client.GetObject(context.Background(), "easy-img", "test/test.txt")
	//读取
	buf := new(bytes.Buffer)
	buf.ReadFrom(data.Body)
	body := buf.String()
	t.Log(body)
}

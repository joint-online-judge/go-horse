package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

var Minio *minio.Client

const (
	ProblemConfigBucketName = "joj2-problem-config"
	SubmissionBucketName    = "joj2-submission"
)

func ConnectMinio() {
	conf := config.Conf
	endpoint := fmt.Sprintf("%s:%d", conf.S3Host, conf.S3Port)
	accessKeyId := conf.S3Username
	secretAccessKey := conf.S3Password
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Debugf("minio client: %+v", minioClient)
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Debugf("minio buckets: %+v", buckets)
	EnsureBucket(minioClient, ProblemConfigBucketName)
	EnsureBucket(minioClient, SubmissionBucketName)
	Minio = minioClient
}

func EnsureBucket(minioClient *minio.Client, bucketName string) {
	if found, err := minioClient.BucketExists(
		context.Background(), bucketName,
	); err != nil {
		logrus.Fatal(err)
	} else if found {
		return
	}
	if err := minioClient.MakeBucket(
		context.Background(), bucketName,
		minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true},
	); err != nil {
		logrus.Fatal(err)
	}
}

func PutSubmission(prefix string, files []*multipart.FileHeader) error {
	return putObjects(SubmissionBucketName, prefix, files)
}

func PutConfig(prefix string, files []*multipart.FileHeader) error {
	return putObjects(ProblemConfigBucketName, prefix, files)
}

func putObjects(bucketName, prefix string, files []*multipart.FileHeader) error {
	for _, file := range files {
		reader, err := file.Open()
		if err != nil {
			return err
		}
		uploadInfo, err := Minio.PutObject(
			context.Background(),
			bucketName,
			file.Filename,
			reader,
			file.Size,
			minio.PutObjectOptions{ContentType: "application/octet-stream"},
		)
		if err != nil {
			return err
		}
		logrus.Debugf("Successfully uploaded bytes: %+v", uploadInfo)
	}
	return nil
}

package storage

import (
	"archive/zip"
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
	var err error
	conf := config.Conf
	endpoint := fmt.Sprintf("%s:%d", conf.S3Host, conf.S3Port)
	accessKeyId := conf.S3Username
	secretAccessKey := conf.S3Password
	Minio, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Debugf("minio client: %+v", Minio)
	EnsureBucket(ProblemConfigBucketName)
	EnsureBucket(SubmissionBucketName)
}

func EnsureBucket(bucketName string) {
	if found, err := Minio.BucketExists(
		context.Background(), bucketName,
	); err != nil {
		logrus.Fatal(err)
	} else if found {
		return
	}
	if err := Minio.MakeBucket(
		context.Background(), bucketName,
		minio.MakeBucketOptions{Region: "", ObjectLocking: false},
	); err != nil {
		logrus.Fatal(err)
	}
}

func PutSubmission(prefix string, files []*multipart.FileHeader) error {
	for _, file := range files {
		reader, err := file.Open()
		if err != nil {
			return err
		}
		defer reader.Close()
		uploadInfo, err := Minio.PutObject(
			context.Background(),
			SubmissionBucketName,
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

func PutProblemConfig(prefix string, file *multipart.FileHeader) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()
	zipFile, err := zip.NewReader(fileReader, file.Size)
	if err != nil {
		return err
	}
	for _, file := range zipFile.File {
		reader, err := file.Open()
		if err != nil {
			return err
		}
		defer reader.Close()
		if file.FileInfo().IsDir() {
			continue
		}
		uploadInfo, err := Minio.PutObject(
			context.Background(),
			SubmissionBucketName,
			file.FileInfo().Name(),
			reader,
			file.FileInfo().Size(),
			minio.PutObjectOptions{ContentType: "application/octet-stream"},
		)
		if err != nil {
			return err
		}
		logrus.Debugf("Successfully uploaded bytes: %+v", uploadInfo)
	}
	return nil
}

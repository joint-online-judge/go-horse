package storage

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func ConnectMinio() {
	conf := config.Conf
	endpoint := fmt.Sprintf("%s:%d", conf.S3Host, conf.S3Port)
	accessKeyID := conf.S3Username
	secretAccessKey := conf.S3Password
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("minioClient: %+v", minioClient)
}

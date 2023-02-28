package storage

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

var Minio *minio.Client

func ConnectMinio() {
	conf := config.Conf
	endpoint := fmt.Sprintf("%s:%d", conf.S3Host, conf.S3Port)
	accessKeyId := conf.S3Username
	secretAccessKey := conf.S3Password
	Minio, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Debugf("minio client: %+v", Minio)
}

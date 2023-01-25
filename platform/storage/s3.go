package storage

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/sirupsen/logrus"
)

func ConnectS3() {
	conf := configs.Conf
	cfg := &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			conf.S3Username,
			conf.S3Password,
			"",
		),
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		logrus.Fatalf("failed to create s3 session: %+v", err)
	}
	client := s3.New(sess)
	client.Handlers.Build.PushBack(func(r *request.Request) {
		if r.HTTPRequest.Method != http.MethodGet {
			return
		}
		r.HTTPRequest.URL.Host = fmt.Sprintf(
			"http://%s:%d",
			conf.S3Host,
			conf.S3Port,
		)
	})
}

package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type conf struct {
	// server config
	Debug             bool     `env:"DEBUG"               envDefault:"false"`
	Https             bool     `env:"HTTPS"               envDefault:"false"`
	Host              string   `env:"HOST"                envDefault:"127.0.0.1"`
	Port              int      `env:"PORT"                envDefault:"34765"`
	Domain            string   `env:"DOMAIN"`
	RootPath          string   `env:"ROOT_PATH"`
	ForwardedAllowIPs []string `env:"FORWARDED_ALLOW_IPS" envDefault:"127.0.0.1" envSeparator:","`

	// postgresql config
	DBHost     string `env:"DB_HOST"     envDefault:"localhost"`
	DBPort     int    `env:"DB_PORT"     envDefault:"5432"`
	DBUsername string `env:"DB_USERNAME" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"pass"`
	DBName     string `env:"DB_NAME"     envDefault:"horse_production"`
	DBEcho     bool   `env:"DB_ECHO"     envDefault:"true"`

	// redis config
	RedisHost     string `env:"REDIS_HOST"     envDefault:"localhost"`
	RedisPort     int    `env:"REDIS_PORT"     envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDbIndex  int    `env:"REDIS_DB_INDEX" envDefault:"0"`

	// rabbitmq config
	RabbitmqHost     string `env:"RABBITMQ_HOST"     envDefault:"localhost"`
	RabbitmqPort     int    `env:"RABBITMQ_PORT"     envDefault:"5672"`
	RabbitmqUsername string `env:"RABBITMQ_USERNAME" envDefault:"guest"`
	RabbitmqPassword string `env:"RABBITMQ_PASSWORD" envDefault:"guest"`
	RabbitmqVhost    string `env:"RABBITMQ_VHOST"    envDefault:""`

	// s3 config (any s3 compatible service)
	S3Host     string `env:"S3_HOST"     envDefault:""`
	S3Port     int    `env:"S3_PORT"     envDefault:"80"`
	S3Username string `env:"S3_USERNAME" envDefault:""`
	S3Password string `env:"S3_PASSWORD" envDefault:""`

	// lakefs config
	LakefsS3Domain string `env:"LAKEFS_S3_DOMAIN" envDefault:"s3.lakefs.example.com"`
	LakefsHost     string `env:"LAKEFS_HOST"      envDefault:""`
	LakefsPort     int    `env:"LAKEFS_PORT"      envDefault:"34766"`
	LakefsUsername string `env:"LAKEFS_USERNAME"  envDefault:"lakefs"`
	LakefsPassword string `env:"LAKEFS_PASSWORD"  envDefault:"lakefs"`

	// buckets
	BucketConfig     string `env:"BUCKET_CONFIG"     envDefault:"s3://joj-config"`
	BucketSubmission string `env:"BUCKET_SUBMISSION" envDefault:"s3://joj-submission"`

	// jwt config
	JwtSecret string `env:"JWT_SECRET"         envDefault:"secret"`
	// JwtAlgorithm     string `env:"JWT_ALGORITHM"      envDefault:"HS256"`
	JwtExpireSeconds int64 `env:"JWT_EXPIRE_SECONDS" envDefault:"1209600"` // 14 days, in seconds

	// oauth config
	OauthJaccount       bool   `env:"OAUTH_JACCOUNT"        envDefault:"False"`
	OauthJaccountId     string `env:"OAUTH_JACCOUNT_ID"     envDefault:""`
	OauthJaccountSecret string `env:"OAUTH_JACCOUNT_SECRET" envDefault:""`

	OauthGithub       bool   `env:"OAUTH_GITHUB"        envDefault:"False"`
	OauthGithubId     string `env:"OAUTH_GITHUB_ID"     envDefault:""`
	OauthGithubSecret string `env:"OAUTH_GITHUB_SECRET" envDefault:""`

	// rollbar `envDefault:config`
	RollbarAccessToken string `env:"ROLLBAR_ACCESS_TOKEN" envDefault:""`
}

var Conf *conf

func Initalize() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	cfg := conf{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse config: %+v", err)
	}
	log.Infof("config object: %+v", cfg)
	Conf = &cfg
}

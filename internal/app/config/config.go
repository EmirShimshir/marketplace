package config

import (
	"github.com/EmirShimshir/marketplace/internal/adapter/authProvider/jwt"
	v1 "github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1"
	"github.com/EmirShimshir/marketplace/pkg/database/postgres"
	"github.com/EmirShimshir/marketplace/pkg/database/redis"
	"github.com/EmirShimshir/marketplace/pkg/gomail"
	"github.com/EmirShimshir/marketplace/pkg/logrus"
	"github.com/EmirShimshir/marketplace/pkg/minio"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Logger   logrus.Config
	Postgres postgres.Config
	JWT      jwt.Config
	Redis    redis.Config
	Minio    minio.Config
	Web      v1.Config
	Gomail   gomail.Config
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := bindEnvConfig(); err != nil {
		panic(errors.Wrap(err, "error reading env"))
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "error reading config file"))
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(errors.Wrap(err, "error unmarshaling config file"))
	}
	return cfg
}

func bindEnvConfig() error {
	bindings := make(map[string]string)
	bindings["web.host"] = "HOST"
	bindings["web.port"] = "PORT"

	bindings["jwt.secret"] = "JWT_SECRET"

	bindings["postgres.database"] = "POSTGRES_DATABASE"
	bindings["postgres.host"] = "POSTGRES_HOST"
	bindings["postgres.port"] = "POSTGRES_PORT"
	bindings["postgres.user"] = "POSTGRES_USER"
	bindings["postgres.password"] = "POSTGRES_PASSWORD"

	bindings["redis.uri"] = "REDIS_URI"

	bindings["minio.endpoint"] = "MINIO_ENDPOINT"
	bindings["minio.user"] = "MINIO_ROOT_USER"
	bindings["minio.password"] = "MINIO_ROOT_PASSWORD"
	bindings["minio.bucketName"] = "MINIO_BUCKET_NAME"
	bindings["minio.bucketName"] = "MINIO_BUCKET_NAME"
	bindings["minio.host"] = "MINIO_HOST"

	bindings["gomail.host"] = "GOMAIL_HOST"
	bindings["gomail.port"] = "GOMAIL_PORT"
	bindings["gomail.senderEmail"] = "GOMAIL_SENDEREMAIL"
	bindings["gomail.password"] = "GOMAIL_PASSWORD"

	for name, binding := range bindings {
		if err := viper.BindEnv(name, binding); err != nil {
			return err
		}
	}

	return nil
}

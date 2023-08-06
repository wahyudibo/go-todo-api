package config

// Config stores configuration object and parse value from environment variables
type Config struct {
	AppPort    int    `envconfig:"APP_PORT" default:"8080"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     int    `envconfig:"DB_PORT" default:"3306"`
	DBUser     string `envconfig:"DB_USER" default:"root"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`

	StorageDriver                 string `envconfig:"STORAGE_DRIVER" default:"local"`
	StoragePath                   string `envconfig:"STORAGE_PATH"`
	LocalStorageDownloadPrefixUrl string `envconfig:"LOCAL_STORAGE_DOWNLOAD_PREFIX_URL"`

	AWSProfile   string `envconfig:"AWS_PROFILE"`
	S3BucketName string `envconfig:"S3_BUCKET_NAME"`
}

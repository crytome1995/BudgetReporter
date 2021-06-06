package config

// Configuration exported
type Configuration struct {
	Aws      AwsConfiguration
	Database DatabaseConfiguration
}

// AwsConfiguration exported
type AwsConfiguration struct {
	DynamoRegion string
	SnsRegion    string
	SnsArn       string
}

// DatabaseConfiguration exported
type DatabaseConfiguration struct {
	TableName string
}

package config

type Config struct {
	AppName string `env:"APP_NAME"`

	PublicServer  HTTP
	DebugServer   HTTP
	KafkaConusmer Kafka
	CH            Clickhouse
}

func New() {}

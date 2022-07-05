package store

type Config struct {
	DataBaseURL string `json:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}

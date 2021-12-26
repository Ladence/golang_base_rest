package storage

type Config struct {
	Dialect string `json:"dialect" yaml:"dialect"`
	Name string `json:"database_name" yaml:"database_name"`
	Host string `json:"database_host" yaml:"database_host"`
	Username string `json:"database_username" yaml:"database_username"`
	Password string `json:"database_password" yaml:"database_password"`
	Port string `json:"port" yaml:"port"`
	MaxConnectAttempts int `json:"max_connect_attempts" yaml:"max_connect_attempts"`
}

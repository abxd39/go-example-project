package domain

type Config struct {
	Server
	Database  Database `json:"database"`
	Logger    Logger   `json:"logger"`
	RateLimit int      `json:"rateLimit"`
}

type Server struct {
	Mode    string `json:"mode"`
	Port    int    `json:"port"`
	Address string `json:"address"`
}

type Database struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
	Host         string `json:"host"`
}

type Logger struct {
	Level string `json:"level"`
}

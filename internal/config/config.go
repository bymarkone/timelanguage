package config

type Cloud struct {
	Provider string
	Address string
	Username string
	Token string
}

type Config struct {
	Cloud Cloud ``
}
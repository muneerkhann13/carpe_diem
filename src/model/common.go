package model

type Config struct {
	MysqlDatabase []struct {
		Host     string `json:"Host"`
		Name     string `json:"Name"`
		Username string `json:"Username"`
		Password string `json:"Password"`
	} `json:"MysqlDatabase"`
}

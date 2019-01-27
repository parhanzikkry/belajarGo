package database

import "belajarGo/config"

var cfg config.Config

func SetConnection() (string, error) {
	cfg = config.GetConfig()
	return `
		dbname=` + cfg.DatabaseConfig.DBName + `
		host=` + cfg.DatabaseConfig.Host + `
		user=` + cfg.DatabaseConfig.User + `
		password=` + cfg.DatabaseConfig.Password + `
		sslmode=` + cfg.DatabaseConfig.SslMode + `
	`, nil
}

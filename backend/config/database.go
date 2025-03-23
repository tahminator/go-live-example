package config

import "fmt"

type Database struct {
	Host     string
	Port     string
	Password string
	User     string
	Name     string
}

// URL generator.
func (db *Database) Url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}

package db

import (
	"github.com/dimiro1/health"
	"gorm.io/gorm"
)

//Checker is a checker that check a given db
type Checker struct {
	DB *gorm.DB
}

// NewChecker returns a new url.Checker with the given URL
func NewPostgresChecker(db *gorm.DB) Checker {
	return Checker{
		DB: db,
	}
}

// Check execute queries in the database
// The first is a simple one used to verify if the database is up
// If is Up then another query is executed, querying for the database version
func (c Checker) Check() health.Health {
	var (
		version string
	)

	h := health.NewHealth()

	if c.DB == nil {
		h.Down().AddInfo("error", "Empty resource")
		return h
	}

	tx := c.DB.Raw(`Select version()`).Scan(&version)
	if tx.Error != nil || version == "" {
		h.Down().AddInfo("error", tx.Error.Error())
		return h
	}

	h.Up().AddInfo("version", version)

	return h
}

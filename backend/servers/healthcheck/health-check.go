package healthcheck

import (
	"alfred/config"
	postgreshealth "alfred/servers/healthcheck/db"
	"github.com/dimiro1/health"
	"gorm.io/gorm"
	"net/http"
)

func RunHealthCheckServer(config *config.Config, db *gorm.DB) error {

	//postgres health check
	postgresql := postgreshealth.NewPostgresChecker(db)
	handler := health.NewHandler()
	postgresql.Check()
	handler.AddChecker("Postgres", postgresql)
	http.Handle("/health/", handler)
	return http.ListenAndServe(config.HealthCheck.Host+":"+config.HealthCheck.Port, nil)
}

func HealthCheckRunner(config *config.Config, db *gorm.DB) error {
	go func() {
		_ = RunHealthCheckServer(config, db)
	}()
	return nil
}

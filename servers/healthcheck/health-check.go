package healthcheck

import (
	"github.com/dimiro1/health"
	"goarcc/config"
	postgreshealth "goarcc/servers/healthcheck/db"
	"gorm.io/gorm"
	"net/http"
)

// RunHealthCheckServer : will start the health check server
func RunHealthCheckServer(config *config.Config, db *gorm.DB) error {

	//postgres health check
	postgresql := postgreshealth.NewPostgresChecker(db)
	handler := health.NewHandler()
	postgresql.Check()
	handler.AddChecker("Postgres", postgresql)
	http.Handle("/health/", handler)
	return http.ListenAndServe(config.HealthCheck.Host+":"+config.HealthCheck.Port, nil)
}

// HealthCheckRunner : will responsible to run health check server in non blocking way.
func HealthCheckRunner(config *config.Config, db *gorm.DB) error {
	go func() {
		_ = RunHealthCheckServer(config, db)
	}()
	return nil
}

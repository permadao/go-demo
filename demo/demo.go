package demo

import (
	"net/http"

	"github.com/go-co-op/gocron"
	"github.com/permadao/go-demo/logger"
	"github.com/permadao/go-demo/storage/redis"
	"github.com/permadao/go-demo/storage/sql"
)

var log = logger.New("demo")

type Demo struct {
	apiServer *http.Server
	scheduler *gocron.Scheduler

	sql   *sql.SQL
	redis *redis.Redis
}

func New(dsn, redisURL string) *Demo {
	return &Demo{
		// sql:   sql.New(dsn),
		// redis: redis.New(redisURL),
	}
}

func (d *Demo) Run(port string) {
	// d.sql.Migrate()
	go d.runAPI(port)
	go d.runJobs()
}

func (d *Demo) Close() {
	d.closeAPI()
}

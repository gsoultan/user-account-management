package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gsoultan/user-account-management/app/account"
	"github.com/gsoultan/user-account-management/database"
	"github.com/gsoultan/user-account-management/helpers/cache"
	"github.com/gsoultan/user-account-management/migration"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultPort            = "8080"
	defaultRoutingProtocol = "http://localhost:7878"
)

func main() {
	var (
		addr = defaultPort
		//rsurl = defaultRoutingProtocol

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		//routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")

		//ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	connectionString := "postgresql://gormroach@localhost:26257/uam?sslmode=disable"
	db, err := database.NewCockroachDB(connectionString)
	if err != nil {
		os.Exit(1)
	}
	defer db.GetConnection().(*gorm.DB).Close()

	//Execute Migration
	migration.Migrate(db)

	cache := cache.NewRedisInstance("tcp", "localhost", 32774)

	accountRepository := account.NewCockroachDBRepository(db.GetConnection())
	accountCommandService, transactionService := account.NewCommandService(accountRepository, cache)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/v1/accounts/", account.MakeHttpHandler(accountCommandService, transactionService, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

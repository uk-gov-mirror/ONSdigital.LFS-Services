package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"services/api"
	"services/api/ws"
	"services/config"
	"time"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if config.Config.LogFormat == "Terminal" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if config.Config.LogLevel == "Debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Command line flag overrides the configuration file
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Info().
		Str("startTime", time.Now().String()).
		Msg("LFS Services: Starting up")

	router := mux.NewRouter()

	importsHandler := api.NewImportsHandler()
	auditHandler := api.NewAuditHandler()
	batchHandler := api.NewBatchHandler()
	loginHandler := api.NewLoginHandler()

	router.HandleFunc("/batches/monthly/{month}/{year}", batchHandler.CreateMonthlyBatchHandler).Methods(http.MethodPost)
	router.HandleFunc("/batches/quarterly/{quarter}/{year}", batchHandler.CreateQuarterlyBatchHandler).Methods(http.MethodPost)

	router.HandleFunc("/imports/survey/gb/{week}/{year}", importsHandler.SurveyUploadGBHandler).Methods(http.MethodPost)
	router.HandleFunc("/imports/survey/ni/{month}/{year}", importsHandler.SurveyUploadNIHandler).Methods(http.MethodPost)
	router.HandleFunc("/imports/address", importsHandler.AddressUploadHandler).Methods(http.MethodPost)

	router.HandleFunc("/audits", auditHandler.HandleAllAuditRequest).Methods(http.MethodGet)
	router.HandleFunc("/audits/year/{year}", auditHandler.HandleYearAuditRequest).Methods(http.MethodGet)
	router.HandleFunc("/audits/month/{year}/{month}", auditHandler.HandleMonthAuditRequest).Methods(http.MethodGet)
	router.HandleFunc("/audits/week/{year}/{week}", auditHandler.HandleWeekAuditRequest).Methods(http.MethodGet)

	router.HandleFunc("/login/{user}", loginHandler.LoginHandler).Methods(http.MethodGet)

	router.HandleFunc("/ws", ws.WebSocketHandler{}.ServeWs).Methods(http.MethodGet)

	listenAddress := config.Config.Service.ListenAddress

	writeTimeout, err := time.ParseDuration(config.Config.Service.WriteTimeout)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LFS").
			Msgf("writeTimeout configuration error")
	}

	readTimeout, err := time.ParseDuration(config.Config.Service.ReadTimeout)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LFS").
			Msgf("readTimeout configuration error")
	}

	// we'll allow anything for now. May need or want to restrict this to just the UI when we know its endpoint
	origins := []string{"*"}
	var c = handlers.AllowedOrigins(origins)

	handlers.CORS(c)(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         listenAddress,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	log.Info().
		Str("listenAddress", listenAddress).
		Str("writeTimeout", writeTimeout.String()).
		Str("readTimeout", readTimeout.String()).
		Msg("LFS Services: Waiting for requests")

	err = srv.ListenAndServe()
	log.Fatal().
		Err(err).
		Str("service", "LFS").
		Msgf("ListenAndServe failed")
}

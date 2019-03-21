package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/honestbee/Zen/config"
	"github.com/honestbee/Zen/examiner"
	"github.com/honestbee/Zen/grpc"
	"github.com/honestbee/Zen/models"
	"github.com/honestbee/Zen/resolvers"
	"github.com/honestbee/Zen/router"
	"github.com/honestbee/Zen/zendesk"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.InfoLevel)

	conf, err := config.New()
	if err != nil {
		logger.Fatal().Err(err).Msgf("new config file failed")
	}

	service, err := models.New(conf)
	if err != nil {
		logger.Fatal().Err(err).Msgf("new model service failed")
	}

	zend, err := zendesk.NewZenDesk(conf)
	if err != nil {
		logger.Fatal().Err(err).Msgf("new zendesk failed")
	}

	exam, err := examiner.NewExaminer(conf, &logger, service, zend)
	if err != nil {
		logger.Fatal().Err(err).Msgf("new examiner failed")
	}

	grpcSvr, err := grpc.New(conf, &logger, service, exam, zend)
	if err != nil {
		logger.Fatal().Err(err).Msgf("new grpc failed")
	}

	graphql, err := resolvers.New(conf, &logger, service, exam, zend)
	if err != nil {
		logger.Fatal().Err(err).Msgf("new graphql failed")
	}

	hmux, err := router.New(conf, &logger, service, exam, zend, graphql)
	if err != nil {
		logger.Fatal().Err(err).Msgf("new router failed")
	}

	srv := &http.Server{
		Addr:         conf.HTTP.ListenAddr,
		Handler:      hmux,
		ReadTimeout:  time.Duration(conf.HTTP.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(conf.HTTP.WriteTimeoutSec) * time.Second,
		IdleTimeout:  time.Duration(conf.HTTP.IdleTimeoutSec) * time.Second,
	}

	if conf.Datadog.Enable {
		// Start the datadog tracer with options.
		var ddHost, ddPort string
		if ddHost = os.Getenv("DATADOG_HOST"); ddHost == "" {
			ddHost = conf.Datadog.Host
		}
		if ddPort = os.Getenv("DATADOG_APM_PORT"); ddPort == "" {
			ddPort = conf.Datadog.Port
		}
		tracer.Start(
			tracer.WithServiceName("helpcenter-zendesk"),
			tracer.WithGlobalTag("env", conf.Datadog.Env),
			tracer.WithAgentAddr(ddHost+":"+ddPort),
			tracer.WithDebugMode(conf.Datadog.Debug),
		)
		defer tracer.Stop()
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	notify := make(chan struct{})
	go func() {
		grpcLis, err := net.Listen("tcp", conf.GRPC.ListenAddr)
		if err != nil {
			logger.Error().Err(err).Msgf("grpc server listen error")
		}

		close(notify)

		if err := grpcSvr.Serve(grpcLis); err != nil {
			logger.Error().Err(err).Msgf("grpc server serve error")
		}
	}()

	go func() {
		<-notify
		if err := srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Info().Msgf("http server close")
			} else {
				logger.Error().Err(err).Msgf("http server error")
			}
		}
	}()

	logger.Info().Msgf("server started")

	<-done
	logger.Info().Msgf("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msgf("https server shutdown failed")
	}

	grpcSvr.GracefulStop()

	if err = service.Close(); err != nil {
		logger.Error().Err(err).Msgf("service close failed")
	}

	if err = exam.Close(); err != nil {
		logger.Error().Err(err).Msgf("examiner close failed")
	}

	logger.Info().Msgf("server shutdown")
}

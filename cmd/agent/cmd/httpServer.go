// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/spidernet-io/rocktemplate/api/v1/server"
	"github.com/spidernet-io/rocktemplate/api/v1/server/restapi"
	"github.com/spidernet-io/rocktemplate/api/v1/server/restapi/healthy"
	"go.uber.org/zap"
)

// ---------- readiness Healthy Hander
type readinessHealthyHander struct {
	logger *zap.Logger
}

func (s *readinessHealthyHander) Handle(r healthy.GetHealthyReadinessParams) middleware.Responder {
	// return healthy.NewGetHealthyReadinessInternalServerError()
	return healthy.NewGetHealthyReadinessOK()
}

// ---------- liveness Healthy Hander
type livenessHealthyHander struct {
	logger *zap.Logger
}

func (s *livenessHealthyHander) Handle(r healthy.GetHealthyLivenessParams) middleware.Responder {
	return healthy.NewGetHealthyLivenessOK()
}

// ---------- startup Healthy Hander
type startupHealthyHander struct {
	logger *zap.Logger
}

func (s *startupHealthyHander) Handle(r healthy.GetHealthyStartupParams) middleware.Responder {

	return healthy.NewGetHealthyStartupOK()
}

// ====================

func SetupHttpServer() {
	logger := rootLogger.Named("http")

	if globalConfig.HttpPort == 0 {
		logger.Sugar().Warn("http server is disabled")
		return
	}
	logger.Sugar().Infof("setup http server at port %v", globalConfig.HttpPort)

	spec, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	if err != nil {
		logger.Sugar().Fatalf("failed to load Swagger spec, reason=%v ", err)
	}

	api := restapi.NewHTTPServerAPIAPI(spec)
	api.Logger = func(s string, i ...interface{}) {
		logger.Sugar().Infof(s, i)
	}

	// setup route
	api.HealthyGetHealthyReadinessHandler = &readinessHealthyHander{logger: logger.Named("route: readiness health")}
	api.HealthyGetHealthyLivenessHandler = &livenessHealthyHander{logger: logger.Named("route: liveness health")}
	api.HealthyGetHealthyStartupHandler = &startupHealthyHander{logger: logger.Named("route: startup health")}

	//
	srv := server.NewServer(api)
	srv.EnabledListeners = []string{"http"}
	// srv.EnabledListeners = []string{"unix"}
	// srv.SocketPath = "/var/run/http-server-api.sock"

	// dfault to listen on "0.0.0.0" and "::1"
	// srv.Host = "0.0.0.0"
	srv.Port = int(globalConfig.HttpPort)
	srv.ConfigureAPI()

	go func() {
		e := srv.Serve()
		s := "http server break"
		if e != nil {
			s += fmt.Sprintf(" reason=%v", e)
		}
		logger.Fatal(s)
	}()
	return
}

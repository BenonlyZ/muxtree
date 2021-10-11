package main

import (
	"muxtree/internal/service"
	web "muxtree/pkg/v1"
	"net/http"
)

func main() {
	svc := web.NewSdkHttpServer("test-service", web.MetricFilterBuilder)
	svc.Route(http.MethodGet, "/", service.Home)
	svc.Route(http.MethodGet, "/user", service.User)
	svc.Route(http.MethodGet, "/user/create", service.CreateUser)
	svc.Route(http.MethodGet, "/order", service.Order)
	svc.Start(":8087")
}

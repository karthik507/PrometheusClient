package main

import (
"fmt"
"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequests = prometheus.NewConterVec(
		prometheus.CounterOpts{
			Name : "http_request_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method","path","status"},
	)
	
	latencyRequests := prometheus.NewHistorgramVec(
		prometheus.HistorgramOpts{
			Name: "http_requests_duration_seconds",
			Help: "Http request latency",
			Buckets: prometheus.DefBuckets,
			},
			[]string{"method" , "path"},
	)
	
	inFlightRequests := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_inflight_requests",
			Help: "Current number of active requests",
		}
	)
	
	errorRequests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_error_requests",
			Help: "Total Http Errors",
		},
		[]string{"method", "path", "status"},
	)
	
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		start := time.Now()
		c.Next()
		status:= c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()
		httpRequest.WithLabelValues(method, path, status).Inc()
	}
}

func main (){
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(latencyRequests)
	prometheus.MustRegister(inFlightRequests)
	prometheus.MustRegister(errorRequests)
	r := gin.Default()
	r.Use(PrometheusMiddleware())
	r.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "healthy",
		})
	}
	r.GET("/metrics" , gin.WrapH(promhttp.Handler()))
	r.Run(":8080")
}







package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func CustomLogger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	// After the request has been processed
	duration := time.Since(start)
	path := c.Path()
	status := c.Response().StatusCode()

	// Log the request and response
	log.Printf("Method: %s, Path: %s, Status: %d, Duration: %v", c.Method(), path, status, duration)

	// Record metrics
	httpRequestsTotal.WithLabelValues(path).Inc()
	httpRequestDuration.WithLabelValues(path).Observe(duration.Seconds())

	return err
}

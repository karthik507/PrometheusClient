# Prometheus Metrics with Gin (Go)

This project demonstrates how to integrate **Prometheus monitoring** into a Go HTTP service built using the Gin framework.
It exposes application metrics such as request count and request latency, which can be scraped by Prometheus and visualized in Grafana.

---

## Overview

This example implements:

* Prometheus metrics using the Go client
* Gin middleware for automatic request instrumentation
* `/metrics` endpoint for Prometheus scraping
* Request counters and latency histograms

Metrics collected:

* `http_requests_total` → Total API requests
* `http_request_duration_seconds` → Request latency
* Labels → `method`, `path`, `status`

---

## Architecture

```
Go API (Gin)
     │
     ▼
/metrics endpoint
     │
     ▼
Prometheus Scrapes Metrics
     │
     ▼
Grafana Dashboards
```

---

## Project Structure

```
.
├── main.go
├── middleware.go
└── README.md
```

Example structure for production:

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── middleware/
│       └── prometheus.go
```

---

## Installation

Clone the repository

```
git clone [https://github.com/yourusername/prometheus-gin-example.git](https://github.com/karthik507/PrometheusClient.git)
cd prometheus-gin-example
```

Install dependencies

```
go mod tidy
```

Dependencies used:

* Gin
* Prometheus Go Client

---

## Running the Application

Start the Go server

```
go run main.go
```

Server will start at

```
http://localhost:8080
```

---

## Metrics Endpoint

Prometheus metrics are exposed at:

```
http://localhost:8080/metrics
```

Example output:

```
http_requests_total{method="GET",path="/hello",status="200"} 5

http_request_duration_seconds_bucket{method="GET",path="/hello",le="0.5"} 5
```

---

## Example API

Test the API:

```
curl http://localhost:8080/hello
```

Response:

```
{"message":"hello world"}
```

Each request updates the Prometheus metrics.

---

## Middleware Implementation

The middleware records request metrics automatically.

```
start := time.Now()

c.Next()

duration := time.Since(start).Seconds()

httpRequests.WithLabelValues(method, path, status).Inc()
httpDuration.WithLabelValues(method, path).Observe(duration)
```

Metrics captured:

* Request count
* Request latency
* HTTP method
* API route
* Status code

---

## Prometheus Configuration

Create `prometheus.yml`:

```
global:
  scrape_interval: 5s

scrape_configs:
  - job_name: "go-service"
    static_configs:
      - targets: ["localhost:8080"]
```

Run Prometheus:

```
docker run -p 9090:9090 \
-v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml \
prom/prometheus
```

Open Prometheus UI:

```
http://localhost:9090
```

Test query:

```
http_requests_total
```

---

## Grafana Dashboard

Run Grafana:

```
docker run -d -p 3000:3000 grafana/grafana
```

Login:

```
http://localhost:3000
username: admin
password: admin
```

Add **Prometheus Data Source**

```
URL: http://localhost:9090
```

Example queries:

Request Rate

```
rate(http_requests_total[1m])
```

Request Latency (P95)

```
histogram_quantile(0.95,
rate(http_request_duration_seconds_bucket[5m]))
```

Requests by Endpoint

```
sum(rate(http_requests_total[5m])) by (path)
```

---

## Example Metrics

```
http_requests_total{method="GET",path="/hello",status="200"} 10

http_request_duration_seconds_bucket{method="GET",path="/hello",le="0.1"} 6
```

---

## Concepts Covered

* Prometheus metrics
* Counter metrics
* Histogram metrics
* Middleware instrumentation
* Metrics labels
* Observability fundamentals

---

## Golden Signals

This example demonstrates two of the **SRE Golden Signals**:

| Signal  | Metric                          |
| ------- | ------------------------------- |
| Traffic | `http_requests_total`           |
| Latency | `http_request_duration_seconds` |

Other signals include:

* Errors
* Saturation

---

## References

Prometheus Documentation
https://prometheus.io/docs/

Prometheus Go Client
https://github.com/prometheus/client_golang

Gin Web Framework
https://github.com/gin-gonic/gin

---

## License

MIT License

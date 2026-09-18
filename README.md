# Logger Go (`logger-go`)

[![Go Reference](https://pkg.go.dev/badge/github.com/pinedadaniel/logger-go.svg)](https://pkg.go.dev/github.com/pinedadaniel/logger-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`logger-go` is a lightweight, high-performance, strongly typed structured logging library for Go. Designed according to the principles of Clean Architecture, it allows you to decouple the implementation details of the logging engine (Logrus) from the logic of your application or microservices.

---

## 🚀 Features

* **Zero Coupling:** Abstracts third-party libraries; your services consume only the `pkg/log` package.

* **Native Structured JSON:** Formatted by default with ISO-8601 (`RFC3339Nano`), ready for ingestion into Datadog, ELK, New Relic, or CloudWatch.

* **Strongly Typed Fields:** Prevents runtime typing errors when constructing log fields (`log.String`, `log.Int`, `log.Err`).

* **Context Support (`context.Context`):** Ready for distributed traceability propagation (`trace_id`, `span_id`).

* **Clean Package-Level Syntax:** Invokes logging methods without propagating or injecting logger instances at each layer.

---

## 📦 setting-up

```bash
go get github.com/tu-usuario/logger-go@v1.0.0

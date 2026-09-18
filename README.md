# Logger Go (`logger-go`)

[![Go Reference](https://pkg.go.dev/badge/github.com/tu-usuario/logger-go.svg)](https://pkg.go.dev/github.com/tu-usuario/logger-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/tu-usuario/logger-go)](https://goreportcard.com/report/github.com/tu-usuario/logger-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`logger-go` es una librería de logging estructurado ligera, de alto rendimiento y fuertemente tipada para Go. Diseñada bajo los principios de *Clean Architecture*, permite desacoplar los detalles de implementación del motor de logs (Logrus) de la lógica de tu aplicación o microservicios.

---

## 🚀 Características

* **Cero Acoplamiento:** Abstrae librerías de terceros; tus servicios consumen únicamente el paquete `pkg/log`.
* **JSON Estructurado Nativo:** Formateado por defecto con ISO-8601 (`RFC3339Nano`) listo para ingesta en Datadog, ELK, New Relic o CloudWatch.
* **Campos Fuertemente Tipados:** Evita errores de tipado en tiempo de ejecución al construir campos de log (`log.String`, `log.Int`, `log.Err`).
* **Soporte de Contexto (`context.Context`):** Preparado para propagación de trazabilidad distribuida (`trace_id`, `span_id`).
* **Sintaxis Límpida a nivel de Paquete:** Invoca métodos de logging sin propagar o inyectar instancias del logger en cada capa.

---

## 📦 Instalación

```bash
go get [github.com/tu-usuario/logger-go@v1.0.0](https://github.com/tu-usuario/logger-go@v1.0.0)

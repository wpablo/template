# rates

Descripción: Rates for Fixed Term Deposits
CMDB: FTD
ID de programa: 5510
Head y Sponsor: Mara Solange Chaizaz y Patricia Lugano

## Starter golang API application

This is the default golang API templated generated by generator-golangsantander

### Features included:
 * End to end endpoint implementation (router -> controller -> service -> apicall mock)
 * Go module definitions
 * Service test using golden files
 * Configuration examples including local values using the `santander-go-kit` configuration provider
 * `santander-go-kit` usage, including app initialization and REST client
 * Default router initialization with resilience features support (panic recovery, circuit breaker)
 * `X-Debug` and `X-Debug-Verbose` header support for internal apicall diagnostics via logs
 * Prometheus metrics server on port `8081`

## Running the project

**Requirements:** golang 1.14.x

To setup your local development environment see the following guide: https://confluence.ar.bsch/display/ASA/Entorno+de+desarrollo+local

**Build:**

```shell
    go build main/main.go
```

**Execute (build included):**

```shell
    go run main/main.go
```

**Run tests:**

Using the native go command:

```shell
    go test ./...
```

Alternatively, using the included formatter and test runner (recommended):

```shell
    ./test.sh
```

To install the dependencies see https://confluence.ar.bsch/display/ASA/Entorno+de+desarrollo+local#Entornodedesarrollolocal-Instalaci%C3%B3nlocaldeformateadoresylinter

**Test coverage:**

The test runner includes the package code coverage. To get a detailed code report, run the following runner to generate an html view of it 

```shell
    ./test_coverage.sh
```

## Debugging

Delve (the golang debugger) is require to perform debug sessions, regardless of the IDE: https://github.com/go-delve/delve/tree/master/Documentation/installation
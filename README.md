[![GitHub Release](https://img.shields.io/github/v/release/jhdrn/go-recoil)](https://github.com/jhdrn/go-recoil/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/jhdrn/go-recoil.svg)](https://pkg.go.dev/github.com/jhdrn/go-recoil)
[![go.mod](https://img.shields.io/github/go-mod/go-version/jhdrn/go-recoil)](go.mod)
[![LICENSE](https://img.shields.io/github/license/jhdrn/go-recoil)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/jhdrn/go-recoil/build.yml?branch=main)](https://github.com/jhdrn/go-recoil/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/jhdrn/go-recoil)](https://goreportcard.com/report/github.com/jhdrn/go-recoil)
[![Codecov](https://codecov.io/gh/jhdrn/go-recoil/branch/main/graph/badge.svg)](https://codecov.io/gh/jhdrn/go-recoil)

## Description
`go-recoil` is a small HTTP handler abstraction library. It enables using a `func(r *http.Request) Response` signature for HTTP handlers and comes with helper functions for fluid response building. It is standard library compatible and works very well with for example the [chi router](https://github.com/go-chi/chi).



## Usage


``` go

import (
    "net/http"

    "github.com/jhdrn/go-recoil"
    "github.com/jhdrn/go-recoil/response"
)

http.HandleFunc("/", recoil.HandlerFunc(func(r *http.Request) recoil.Response {
    return response.OK()
}))

http.Handle("/", recoil.Handler(func(r *http.Request) recoil.Response {
    return response.OK()
}))


```



## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
//
// Copyright (c) 2018
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package mainflux

import (
	"encoding/json"
	"net/http"
)

const version string = "0.5.0"

type response struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

// Version exposes an HTTP handler for retrieving service version.
func Version(service string) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		res := response{service, version}

		data, _ := json.Marshal(res)

		rw.Write(data)
	})
}

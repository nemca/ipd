/*
Copyright © 2022 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handlers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/nemca/ipd/internal/config"
	"github.com/nemca/ipd/models"
)

// RootHandler main http handler
type RootHandler struct {
	config *config.Config
}

func NewRootHandler(config *config.Config) *RootHandler {
	return &RootHandler{config: config}
}

func (h *RootHandler) GetIP(w http.ResponseWriter, r *http.Request) {
	resp := new(models.Response)
	toJSON := false

	output := r.URL.Query().Get("output")
	if output == "json" {
		toJSON = true
	}

	forwarderFor := r.Header.Get(h.config.HTTP.ForwardedHeader)
	if forwarderFor != "" {
		resp.IP = forwarderFor
		fmt.Fprintf(w, resp.Make(toJSON))
		return
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.IP = host
	fmt.Fprintf(w, resp.Make(toJSON))
}

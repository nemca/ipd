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
)

// RootHandler main http handler
type RootHandler struct {
	config *config.Config
}

func NewRootHandler(config *config.Config) *RootHandler {
	return &RootHandler{config: config}
}

func (h *RootHandler) GetIP(w http.ResponseWriter, r *http.Request) {
	forwarderFor := r.Header.Get(h.config.HTTP.ForwardedHeader)
	if forwarderFor != "" {
		fmt.Fprintf(w, forwarderFor)
		return
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, host)
}

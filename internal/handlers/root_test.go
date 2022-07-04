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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nemca/ipd/internal/config"
	"gotest.tools/assert"
)

var (
	ip string = "192.0.2.1"
)

var (
	version string = "test"
	build   string = "test"
)

func TestGetIP(t *testing.T) {
	cfg, err := config.Init(version, build)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	cfg.HTTP.ForwardedHeader = "X-Forwarded-For"
	req.Header.Add(cfg.HTTP.ForwardedHeader, ip)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router(cfg).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ip, string(body))
}

func router(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", NewRootHandler(cfg).GetIP).Methods(http.MethodGet)
	return r
}

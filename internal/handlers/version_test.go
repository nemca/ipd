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

	"github.com/nemca/ipd/internal/config"
	"gotest.tools/assert"
)

const (
	wantVersion string = `{"version":"test","commint":"test"}`
)

func TestGetVersion(t *testing.T) {
	cfg, err := config.Init(version, build)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router(cfg).ServeHTTP(rr, req)
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, wantVersion, string(body))
}

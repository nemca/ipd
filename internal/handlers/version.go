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
	"encoding/json"
	"net/http"

	"github.com/nemca/ipd/internal/config"
	"github.com/nemca/ipd/models"
)

// VersionHandler handler for version and build info
type VersionHandler struct {
	config *config.Config
}

// NewVersionHandler returns new VersionHandler
func NewVersionHandler(config *config.Config) *VersionHandler {
	return &VersionHandler{config: config}
}

// GetVersion returns version and build info in JSON format
func (h *VersionHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	resp := models.VersionResponse{
		Version: h.config.Version,
		Commit:  h.config.Build,
	}

	data, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

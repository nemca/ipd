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

package models

import "encoding/json"

// Response helds IP address for response
type Response struct {
	IP string `json:"ip"`
}

// String implements fmt.Stringer interface
func (r *Response) String() string {
	return r.IP
}

// Make builds response string in text or JSON formats
func (r *Response) Make(j bool) string {
	if j {
		data, _ := json.Marshal(r)
		return string(data)
	}
	return r.IP
}

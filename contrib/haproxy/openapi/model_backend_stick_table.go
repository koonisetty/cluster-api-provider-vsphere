/*
Copyright 2020 The Kubernetes Authors.

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

/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BackendStickTable struct for BackendStickTable
type BackendStickTable struct {
	Expire  *int32 `json:"expire,omitempty"`
	Keylen  *int32 `json:"keylen,omitempty"`
	Nopurge bool   `json:"nopurge,omitempty"`
	Peers   string `json:"peers,omitempty"`
	Size    *int32 `json:"size,omitempty"`
	Store   string `json:"store,omitempty"`
	Type    string `json:"type,omitempty"`
}

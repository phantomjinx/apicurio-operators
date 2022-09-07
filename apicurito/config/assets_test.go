/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NoErrorAndNotEmptyBytes(t *testing.T, path string, callable func(path string) ([]byte, error)) {
	t.Helper()

	object, err := callable(path)

	assert.Nil(t, err)
	assert.NotEmpty(t, object)
}
func NoErrorAndNotEmptyString(t *testing.T, path string, callable func(path string) (string, error)) {
	t.Helper()

	object, err := callable(path)

	assert.Nil(t, err)
	assert.NotEmpty(t, object)
}

func NoErrorAndContains(t *testing.T, path string, contains string, callable func(path string) ([]string, error)) {
	t.Helper()

	elements, err := callable(path)

	assert.Nil(t, err)
	assert.Contains(t, elements, contains)
}
func NoErrorAndNotContains(t *testing.T, path string, contains string, callable func(path string) ([]string, error)) {
	t.Helper()

	elements, err := callable(path)

	assert.Nil(t, err)
	assert.NotContains(t, elements, contains)
}
func NoErrorAndEmpty(t *testing.T, path string, callable func(path string) ([]string, error)) {
	t.Helper()

	elements, err := callable(path)

	assert.Nil(t, err)
	assert.Empty(t, elements)
}

func ErrorBytes(t *testing.T, path string, callable func(path string) ([]byte, error)) {
	t.Helper()

	_, err := callable(path)
	assert.NotNil(t, err)
}
func ErrorString(t *testing.T, path string, callable func(path string) (string, error)) {

	t.Helper()

	_, err := callable(path)
	assert.NotNil(t, err)
}

func TestGetAsset(t *testing.T) {
	NoErrorAndNotEmptyBytes(t, "samples/apicur_v1_apicurito_cr.yaml", Asset)
	NoErrorAndNotEmptyBytes(t, "/samples/apicur_v1_apicurito_cr.yaml", Asset)
	NoErrorAndNotEmptyString(t, "samples/apicur_v1_apicurito_cr.yaml", AssetAsString)
}

func TestAssets(t *testing.T) {
	NoErrorAndContains(t, "samples", "samples/apicur_v1_apicurito_cr.yaml", Assets)
	NoErrorAndNotContains(t, "/samples/", "kustomize.yaml", Assets)
	NoErrorAndEmpty(t, "/dirnotexist", Assets)

	items, err := Assets("manifests")
	assert.Nil(t, err)

	for _, res := range items {
		if strings.Contains(res, "apicurito.clusterserviceversion.yaml") {
			assert.Fail(t, "Assets should not return nested files")
		}
		if strings.Contains(res, "bases") {
			assert.Fail(t, "Assets should not return nested dirs")
		}
	}

	NoErrorAndContains(t, "/manifests/bases", "manifests/bases/apicurito.clusterserviceversion.yaml", Assets)
}

func TestCRDAssets(t *testing.T) {
	NoErrorAndNotEmptyBytes(t, "/crd/bases/apicur.io_apicuritoes.yaml", Asset)
}

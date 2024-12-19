// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/mod/module"
)

const (
	defaultGoProxyURL string = "https://proxy.golang.org"
)

var (
	ErrInvalidEscapedModulePath = errors.New("invalid escaped module path")
	ErrModuleNotFound           = errors.New("module not found in proxy")
)

// downloadModule downloads the content a given module path/version.
func downloadModule(module *module.Version) ([]byte, error) {
	cleanPath := strings.ToLower(module.Path)
	urlPath := fmt.Sprintf("%s/%s/@v/%s.zip", goProxyURLOpt, cleanPath, module.Version)
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		return body, nil
	}
	if resp.StatusCode == 410 && strings.Contains(string(body), "invalid escaped module path") {
		return nil, ErrInvalidEscapedModulePath
	}
	if resp.StatusCode == 404 {
		return nil, ErrModuleNotFound
	}

	return nil, fmt.Errorf("unhandled response status code: %v", resp.StatusCode)
}

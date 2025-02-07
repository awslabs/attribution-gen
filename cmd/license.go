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

	licenseclassifier "github.com/google/licenseclassifier/v2"
	"github.com/google/licenseclassifier/v2/assets"
)

var (
	ErrorLicenseNotFound = errors.New("license not found")
	ErrorLicenseUnknown  = errors.New("license type unknown")
)

type LicenseType int

const (
	defaultConfidenceTreshHold = 0.9
)

// License represents a sotfware distribution and usage license
type License struct {
	// The name of the license
	Name string
	// If the name of the license if unkown, this field will be
	// filled
	Data []byte
}

// newLicenseClassifier instantiate a new classifer and returns
// a licenseClassifierWrapper
func newLicenseClassifier(
	confidenceThreshold float64,
) (*licenseClassifierWrapper, error) {
	classifier, err := assets.DefaultClassifier()
	if err != nil {
		return nil, err
	}

	// TODO: it looks like the default clf doesn't support changing confidence threshold,
	// we can use the confidenceThreshold level to filter out results less than it
	return &licenseClassifierWrapper{
		confidenceThreshold: confidenceThreshold,
		classifier:          classifier,
	}, nil
}

// licenseClassifierWrapper is a wrapper around licenseclassifier.License
type licenseClassifierWrapper struct {
	confidenceThreshold float64
	classifier          *licenseclassifier.Classifier
}

// detectLicense takes the byte of a given license and returns its name
func (lcw *licenseClassifierWrapper) detectLicense(data []byte) (string, error) {
	matches := lcw.classifier.Match(data).Matches
	if matches.Len() == 0 {
		return "", ErrorLicenseUnknown
	}

	best := 0
	for i := range matches {
		if matches.Less(best, i) {
			best = i
		}
	}

	licenseName := matches[best].Name
	// license detection library can return "Copyright" which is not a valid license name/type
	if licenseName == "Copyright" {
		return "", ErrorLicenseUnknown
	}
	return licenseName, nil
}

package main

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
	LicenseUnknown LicenseType = iota
	LicenseApache20
	LicenseMIT
	LicenseBSD3Clause
	LicenseISC
	LicenseOther

	defaultConfidenceTreshHold = 0.9
)

// getLicenseType takes a license name and returns the equivalent
// LicenseType constant
func getLicenseType(name string) LicenseType {
	switch name {
	case "Apache-2.0":
		return LicenseApache20
	case "MIT":
		return LicenseMIT
	case "BSD-3-Clause":
		return LicenseMIT
	case "ISC":
		return LicenseISC
	default:
		return LicenseOther
	}
}

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
	return licenseName, nil
}

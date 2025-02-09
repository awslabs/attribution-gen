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

const (
	defaultAttributionFileHeaderTemplate = `# Open Source Software Attribution

[//]: # (File generated by attribution-gen. DO NOT EDIT.)

The {{ .ModuleName }} source code is licensed under the
{{ .ModuleLicense }} license. A copy of this license is available in the
[LICENSE](LICENSE) file in the root source code directory and is included,
along with this document, in any images containing {{ .ModuleName }}
binaries.

## Package dependencies

The module {{ .ModuleName }} depends on a number of Open Source Go packages. Direct
dependencies are listed in the ./go.mod file.
Those direct package dependencies have some dependencies of their own (known as
"transitive dependencies")

In this part of the Attribution document, we list our dependent packages and
include an indication of the Open Source License under which that package is
distributed. For any package *NOT* distributed under the terms of the Apache
License version 2.0, we include the full text of the package's License below.

{{ if .Tree.Root.Dependencies -}}
{{ range $dependency := .Tree.Root.Dependencies -}}
* ` + "`" + "{{ .Version.Path }}" + "`" + `
{{ end -}}
{{ end -}}
`

	defaultAttributionModuleBlockTemplate = `
{{ .TitlePrefix }} {{ .Name }}

{{ if .LicenseIdentifier -}}
License Identifier: {{ .LicenseIdentifier }}

{{ end -}}
{{ if .License -}}
{{ .License }}

{{ end -}}
{{ if .Dependencies -}}
Subdependencies:
{{ range $dependency := .Dependencies -}}
* ` + "`" + "{{ .Version.Path }}" + "`" + `
{{ end }}
{{- end }}
`
)

type attributionModuleVars struct {
	TitlePrefix       string
	Name              string
	LicenseIdentifier string
	License           string
	Dependencies      []*Module
}

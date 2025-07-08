module github.com/geomyidia/util

go 1.23.0

toolchain go1.24.4

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

// Overrides
// https://github.com/geomyidia/util/security/dependabot/2:
require gopkg.in/yaml.v3 v3.0.0-20220521103104-8f96da9f5d5e // indirect

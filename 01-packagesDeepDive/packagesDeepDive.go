package _0201_packagesDeepDive_

import (
	nestedPackages "go-advanced/01-packagesDeepDive/01-nestedPackages"
	configuringPackages "go-advanced/01-packagesDeepDive/02-configuringPackages"
)

var Demos = struct {
	NestedPackages      func()
	ConfiguringPackages func()
}{
	NestedPackages:      nestedPackages.Demo,
	ConfiguringPackages: configuringPackages.Demo,
}

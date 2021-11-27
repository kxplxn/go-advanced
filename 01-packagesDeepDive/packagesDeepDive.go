package _0201_packagesDeepDive_

import (
	nestedPackages "go-advanced/01-packagesDeepDive/01-nestedPackages"
	configuringPackages "go-advanced/01-packagesDeepDive/02-configuringPackages"
	importingPackages "go-advanced/01-packagesDeepDive/03-importingPackages"
	alternateImportMethods "go-advanced/01-packagesDeepDive/04-alternateImportMethods"
	documentingPackages "go-advanced/01-packagesDeepDive/05-documentingPackages"
	theInitFunction "go-advanced/01-packagesDeepDive/06-theInitFunction"
)

var Demos = struct {
	NestedPackages         func()
	ConfiguringPackages    func()
	ImportingPackages      func()
	AlternateImportMethods func()
	DocumentingPackages    func()
	TheInitFunction        func()
}{
	NestedPackages:         nestedPackages.Demo,
	ConfiguringPackages:    configuringPackages.Demo,
	ImportingPackages:      importingPackages.Demo,
	AlternateImportMethods: alternateImportMethods.Demo,
	DocumentingPackages:    documentingPackages.Demo,
	TheInitFunction:        theInitFunction.Demo,
}

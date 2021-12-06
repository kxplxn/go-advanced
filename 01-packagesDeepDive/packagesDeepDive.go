package _0201_packagesDeepDive_

import (
	nestedPackages "go-advanced/01-packagesDeepDive/01-nestedPackages"
	configuringPackages "go-advanced/01-packagesDeepDive/02-configuringPackages"
	importingPackages "go-advanced/01-packagesDeepDive/03-importingPackages"
	alternateImportMethods "go-advanced/01-packagesDeepDive/04-alternateImportMethods"
	documentingPackages "go-advanced/01-packagesDeepDive/05-documentingPackages"
	theInitFunction "go-advanced/01-packagesDeepDive/06-theInitFunction"
	modules "go-advanced/01-packagesDeepDive/07-modules"
)

func Demos() {
	nestedPackages.Demo()
	configuringPackages.Demo()
	importingPackages.Demo()
	alternateImportMethods.Demo()
	documentingPackages.Demo()
	theInitFunction.Demo()
	modules.Demo()
}

package main

import (
	packagesDeepDive "go-advanced/01-packagesDeepDive"
	concurrency "go-advanced/02-concurrency"
)

func main() {
	packagesDeepDive.Demos.NestedPackages()
	packagesDeepDive.Demos.ConfiguringPackages()
	packagesDeepDive.Demos.ImportingPackages()
	packagesDeepDive.Demos.AlternateImportMethods()
	packagesDeepDive.Demos.DocumentingPackages()
	packagesDeepDive.Demos.TheInitFunction()
	packagesDeepDive.Demos.Modules()

	concurrency.Demos.ImprovingPerformance()
}

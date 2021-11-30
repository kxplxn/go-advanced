package main

import (
	packagesDeepDive "go-advanced/01-packagesDeepDive"
	//concurrency "go-advanced/02-concurrency"
	goroutines "go-advanced/03-goroutines"
)

func main() {
	packagesDeepDive.Demos.NestedPackages()
	packagesDeepDive.Demos.ConfiguringPackages()
	packagesDeepDive.Demos.ImportingPackages()
	packagesDeepDive.Demos.AlternateImportMethods()
	packagesDeepDive.Demos.DocumentingPackages()
	packagesDeepDive.Demos.TheInitFunction()
	packagesDeepDive.Demos.Modules()

	//concurrency.Demos.ImprovingPerformance()

	goroutines.Demos.Goroutines()
	goroutines.Demos.AnonymousFunctions()
	goroutines.Demos.AssigningFunctions()
	goroutines.Demos.ReturningFunctions()
	goroutines.Demos.Closures()
	goroutines.Demos.AnonymousGoroutines()
	goroutines.Demos.WaitGroups()
	goroutines.Demos.AtomicOperations()
}

package main

import (
	packagesDeepDive "go-advanced/01-packagesDeepDive"
	//concurrency "go-advanced/02-concurrency"
	goroutines "go-advanced/03-goroutines"
	channels "go-advanced/04-channels"
	databases "go-advanced/05-databases"
)

func main() {
	packagesDeepDive.Demos.NestedPackages()
	packagesDeepDive.Demos.ConfiguringPackages()
	packagesDeepDive.Demos.ImportingPackages()
	packagesDeepDive.Demos.AlternateImportMethods()
	packagesDeepDive.Demos.DocumentingPackages()
	packagesDeepDive.Demos.TheInitFunction()
	packagesDeepDive.Demos.Modules()

	// concurrency.Demos.ImprovingPerformance() // takes a long time

	goroutines.Demos.Goroutines()
	goroutines.Demos.AnonymousFunctions()
	goroutines.Demos.AssigningFunctions()
	goroutines.Demos.ReturningFunctions()
	goroutines.Demos.Closures()
	goroutines.Demos.AnonymousGoroutines()
	goroutines.Demos.WaitGroups()
	goroutines.Demos.AtomicOperations()
	goroutines.Demos.Mutexes()
	goroutines.Demos.RaceConditions()
	goroutines.Demos.PreventingDataRaces()
	goroutines.Demos.UsingRaceDetector()

	channels.Demos.GoroutinesAndChannels()
	channels.Demos.UnbufferedChannels()
	channels.Demos.BufferedChannels()
	channels.Demos.RangeAndClose()
	channels.Demos.SelectStatement()
	channels.Demos.NonblockingSelect()
	channels.Demos.ChannelsAndWaitGroups()
	channels.Demos.Pipelines()

	databases.Demos.AccessingDatabases()
	databases.Demos.SingleRowDataRetrieval()
	databases.Demos.MultipleRowsDataRetrieval()
	databases.Demos.PreparedStatements()
	databases.Demos.DataManipulation()
	databases.Demos.HandlingErrors()
}

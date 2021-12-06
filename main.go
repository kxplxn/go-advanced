package main

import (
	packagesDeepDive "go-advanced/01-packagesDeepDive"
	//concurrency "go-advanced/02-concurrency"
	goroutines "go-advanced/03-goroutines"
	channels "go-advanced/04-channels"
	databases "go-advanced/05-databases"
	testingAndDeploying "go-advanced/06-testingAndDeploying"
)

func main() {
	packagesDeepDive.Demos()
	//concurrency.Demos() // takes a looooong time
	goroutines.Demos()
	channels.Demos()
	databases.Demos()
	testingAndDeploying.Demos()
}

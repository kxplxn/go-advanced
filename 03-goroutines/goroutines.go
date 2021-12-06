package _0203_goroutines_

import (
	goroutines "go-advanced/03-goroutines/01-goroutines"
	anonymousFunctions "go-advanced/03-goroutines/02-closures/01-anonymousFunctions"
	assigningFunctions "go-advanced/03-goroutines/02-closures/02-assigningFunctions"
	returningFunctions "go-advanced/03-goroutines/02-closures/03-returningFunctions"
	closures "go-advanced/03-goroutines/02-closures/04-closures"
	goroutinesAsClosures "go-advanced/03-goroutines/03-goroutinesAsClosures"
	waitGroups "go-advanced/03-goroutines/04-waitGroups"
	atomicOperations "go-advanced/03-goroutines/05-atomicOperations"
	mutexes "go-advanced/03-goroutines/06-mutexes"
	raceConditions "go-advanced/03-goroutines/07-raceConditions"
	preventingDataRaces "go-advanced/03-goroutines/08-preventingDataRaces"
	usingRaceDetector "go-advanced/03-goroutines/09-usingRaceDetector"
)

func Demos() {
	goroutines.Demo()
	anonymousFunctions.Demo()
	assigningFunctions.Demo()
	returningFunctions.Demo()
	closures.Demo()
	goroutinesAsClosures.Demo()
	waitGroups.Demo()
	atomicOperations.Demo()
	mutexes.Demo()
	raceConditions.Demo()
	preventingDataRaces.Demo()
	usingRaceDetector.Demo()
}

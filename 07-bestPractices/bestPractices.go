package _0207_bestPractices_

import (
	documenting "go-advanced/07-bestPractices/01-documenting"
	structuring "go-advanced/07-bestPractices/02-structuring"
	errorHandling "go-advanced/07-bestPractices/03-errorHandling"
	dataHandling "go-advanced/07-bestPractices/04-dataHandling"
	concurrency "go-advanced/07-bestPractices/05-concurrency"
	testing "go-advanced/07-bestPractices/06-testing"
)

func Demos() {
	documenting.Demo()
	structuring.Demo()
	errorHandling.DemoPanic()
	errorHandling.DemoTypeSwitch()
	dataHandling.Demo()
	concurrency.Demo()
	testing.Demo()
}

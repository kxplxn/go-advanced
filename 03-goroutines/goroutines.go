package _0203_goroutines_

import (
	goroutines "go-advanced/03-goroutines/01-goroutines"
	anonymousFunctions "go-advanced/03-goroutines/02-closures/01-anonymousFunctions"
	assigningFunctions "go-advanced/03-goroutines/02-closures/02-assigningFunctions"
	returningFunctions "go-advanced/03-goroutines/02-closures/03-returningFunctions"
)

var Demos = struct {
	Goroutines         func()
	AnonymousFunctions func()
	AssigningFunctions func()
	ReturningFunctions func()
}{
	Goroutines:         goroutines.Demo,
	AnonymousFunctions: anonymousFunctions.Demo,
	AssigningFunctions: assigningFunctions.Demo,
	ReturningFunctions: returningFunctions.Demo,
}

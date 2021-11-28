package _0203_goroutines_

import (
	goroutines "go-advanced/03-goroutines/01-goroutines"
	anonymousFunctions "go-advanced/03-goroutines/02-closures/01-anonymousFunctions"
)

var Demos = struct {
	Goroutines         func()
	AnonymousFunctions func()
}{
	Goroutines:         goroutines.Demo,
	AnonymousFunctions: anonymousFunctions.Demo,
}

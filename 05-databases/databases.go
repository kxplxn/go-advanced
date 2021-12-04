package _0205_databases_

import (
	accessingDatabases "go-advanced/05-databases/01-accessingDatabases"
)

var Demos = struct {
	AccessingDatabases func()
}{
	AccessingDatabases: accessingDatabases.Demo,
}

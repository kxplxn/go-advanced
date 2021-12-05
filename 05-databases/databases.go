package _0205_databases_

import (
	accessingDatabases "go-advanced/05-databases/01-accessingDatabases"
	dataRetrieval "go-advanced/05-databases/02-dataRetrieval"
)

var Demos = struct {
	AccessingDatabases        func()
	SingleRowDataRetrieval    func()
	MultipleRowsDataRetrieval func()
}{
	AccessingDatabases:        accessingDatabases.Demo,
	SingleRowDataRetrieval:    dataRetrieval.SingleRowDemo,
	MultipleRowsDataRetrieval: dataRetrieval.MultipleRowsDemo,
}

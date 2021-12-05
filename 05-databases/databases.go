package _0205_databases_

import (
	accessingDatabases "go-advanced/05-databases/01-accessingDatabases"
	dataRetrieval "go-advanced/05-databases/02-dataRetrieval"
	preparedStatements "go-advanced/05-databases/03-preparedStatements"
)

var Demos = struct {
	AccessingDatabases        func()
	SingleRowDataRetrieval    func()
	MultipleRowsDataRetrieval func()
	PreparedStatements        func()
}{
	AccessingDatabases:        accessingDatabases.Demo,
	SingleRowDataRetrieval:    dataRetrieval.SingleRowDemo,
	MultipleRowsDataRetrieval: dataRetrieval.MultipleRowsDemo,
	PreparedStatements:        preparedStatements.Demo,
}

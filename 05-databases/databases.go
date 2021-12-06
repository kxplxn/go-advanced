package _0205_databases_

import (
	accessingDatabases "go-advanced/05-databases/01-accessingDatabases"
	dataRetrieval "go-advanced/05-databases/02-dataRetrieval"
	preparedStatements "go-advanced/05-databases/03-preparedStatements"
	dataManipulation "go-advanced/05-databases/04-dataManipulation"
	transactions "go-advanced/05-databases/05-transactions"
	errorsNullsAndUnknownColumns "go-advanced/05-databases/06-errorsNullsAndUnknownColumns"
)

func Demos() {
	accessingDatabases.Demo()
	dataRetrieval.SingleRowDemo()
	dataRetrieval.MultipleRowsDemo()
	preparedStatements.Demo()
	dataManipulation.Demo()
	transactions.Demo()
	errorsNullsAndUnknownColumns.DemoHandlingErrors()
	errorsNullsAndUnknownColumns.DemoHandlingNulls()
	errorsNullsAndUnknownColumns.DemoHandlingUnknownTypes()
}

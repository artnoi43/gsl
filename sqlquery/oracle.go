package sqlquery

import (
	"errors"
	"fmt"
)

func InsertAllOracle(items ...ModelCreate) (string, []interface{}, error) {
	if len(items) == 0 {
		return "", nil, errors.New("empty items slice")
	}

	tableName := items[0].TableName()
	cols := items[0].ColumnsCreate()
	lenCols := len(cols)

	var valuesAll []interface{}
	query := "insert all"
	clauseColumns := ClauseColumns(cols)

	bindPointer := 1
	for i := range items {
		query += fmt.Sprintf(
			" into %s %s values %s",
			tableName, clauseColumns, ClauseValues(Colon, uint(bindPointer), uint(lenCols)),
		)

		valuesAll = append(valuesAll, items[i].ValuesCreate()...)
		bindPointer += lenCols
	}

	// Dummy SELECT clause to avoid Oracle DB throwing error
	// https://stackoverflow.com/questions/73751/what-is-the-dual-table-in-oracle
	query += " select * from dual"

	return query, valuesAll, nil
}

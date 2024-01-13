package sqlquery

import (
	"errors"
	"fmt"
)

// InsertAll returns query and bind values for INSERT ALL into a table.
// All members of `items` must map to the same table.
func InsertAll(placeholder Placeholder, items ...ModelCreate) (string, []interface{}, error) {
	if len(items) == 0 {
		return "", nil, errors.New("empty items slice")
	}

	tableName := items[0].TableName()
	columns := items[0].ColumnsCreate()
	lenCols := len(columns)

	var valuesAll []interface{}

	query := fmt.Sprintf("insert all into %s ", tableName)
	query += ClauseColumns(columns)
	query += " values "

	bindPointer := 1
	lenItems := len(items)

	for i := range items {
		query += ClauseValues(placeholder, uint(bindPointer), uint(lenCols))

		if i != lenItems-1 {
			query += " "
		}

		valuesAll = append(valuesAll, items[i].ValuesCreate()...)

		bindPointer += lenCols
	}

	return query, valuesAll, nil
}

package sqlquery

import (
	"fmt"
)

type Placeholder uint8

const (
	QuestionMark Placeholder = iota + 1
	Dollar
	Colon
)

func ClauseColumns(columns []string) string {
	lenColumns := len(columns)
	if lenColumns == 0 {
		return ""
	}

	clauseColumns := "("
	for i := range columns {
		clauseColumns += columns[i]
		if i != lenColumns-1 {
			clauseColumns += ","
		}
	}

	clauseColumns += ")"

	return clauseColumns
}

func ClauseValuesQuestionMark(lenColumns uint) string {
	clause := "("
	for i := uint(0); i < lenColumns; i++ {
		clause += "?"

		if i != lenColumns-1 {
			clause += ","
		}
	}

	clause += ")"

	return clause
}

func ClauseValuesNumbered(start, lenColumns uint, placeholder rune) string {
	clause := "("
	for i := uint(0); i < lenColumns; i++ {
		clause += fmt.Sprintf("%c%d", placeholder, start+i)

		if i != lenColumns-1 {
			clause += ","
		}
	}

	clause += ")"

	return clause
}

func ClauseValues(placeholder Placeholder, start, lenColumns uint) string {
	if lenColumns == 0 {
		return ""
	}

	switch placeholder {
	case Colon:
		return ClauseValuesNumbered(start, lenColumns, ':')

	case Dollar:
		return ClauseValuesNumbered(start, lenColumns, '$')

	case QuestionMark:
		return ClauseValuesQuestionMark(lenColumns)
	}

	panic(fmt.Sprintf("invalid placeholder %d", placeholder))
}

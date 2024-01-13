package sqlquery

import (
	"testing"
)

type foo struct {
	id   uint64
	name string
	age  uint8
}

func (f *foo) TableName() string {
	return "FOO"
}

func (f *foo) ColumnsCreate() []string {
	return []string{"ID", "NAME", "AGE"}
}

func (f *foo) ValuesCreate() []interface{} {
	return []interface{}{f.id, f.name, f.age}
}

func (f *foo) MapColumnValuesCreate() map[string]interface{} {
	return map[string]interface{}{
		"ID":   f.id,
		"NAME": f.name,
		"AGE":  f.age,
	}
}

func (f *foo) ColumnsUpdate() []string {
	return []string{"NAME", "AGE"}
}

func (f *foo) ValuesUpdate() []interface{} {
	return []interface{}{f.name, f.age}
}

func (f *foo) MapColumnValuesUpdate() map[string]interface{} {
	return map[string]interface{}{
		"ID":   f.id,
		"NAME": f.name,
		"AGE":  f.age,
	}
}

func TestFmtClauseColumns(t *testing.T) {
	type test struct {
		cols     []string
		expected string
	}

	tests := []test{
		{
			cols:     []string{"a", "b", "c"},
			expected: "(a,b,c)",
		},
		{
			cols:     []string{"a"},
			expected: "(a)",
		},
	}

	for i := range tests {
		test := &tests[i]

		result := ClauseColumns(test.cols)
		if result != test.expected {
			t.Logf("Unexpected columns")
			t.Logf("Expecting: \"%s\"", test.expected)
			t.Logf("Got: \"%s\"", result)

			t.Fatalf("unexpected columns")
		}
	}
}

func TestFmtClauseValuesQMark(t *testing.T) {
	type test struct {
		cols     []string
		start    uint
		expected string
	}

	tests := []test{
		{cols: []string{}, start: 1, expected: ""},
		{cols: []string{"a", "b", "c"}, start: 1, expected: "(?,?,?)"},
		{cols: []string{"a", "b", "c"}, start: 2, expected: "(?,?,?)"},
		{cols: []string{"a", "b"}, start: 5, expected: "(?,?)"},
		{cols: []string{"a"}, start: 5, expected: "(?)"},
	}

	for i := range tests {
		test := &tests[i]
		result := ClauseValues(QuestionMark, test.start, uint(len(test.cols)))

		if result != test.expected {
			t.Logf("Unexpected values")
			t.Logf("Expecting: \"%s\"", test.expected)
			t.Logf("Got: \"%s\"", result)

			t.Fatalf("unexpected values")
		}
	}
}

func TestFmtClauseValuesColon(t *testing.T) {
	type test struct {
		cols     []string
		start    uint
		expected string
	}

	tests := []test{
		{cols: []string{}, start: 1, expected: ""},
		{cols: []string{"a", "b", "c"}, start: 1, expected: "(:1,:2,:3)"},
		{cols: []string{"a", "b", "c"}, start: 2, expected: "(:2,:3,:4)"},
		{cols: []string{"a", "b"}, start: 5, expected: "(:5,:6)"},
	}

	for i := range tests {
		test := &tests[i]
		result := ClauseValues(Colon, test.start, uint(len(test.cols)))

		if result != test.expected {
			t.Logf("Unexpected values")
			t.Logf("Expecting: \"%s\"", test.expected)
			t.Logf("Got: \"%s\"", result)

			t.Fatalf("unexpected values")
		}
	}
}

func TestFmtClauseValuesDollar(t *testing.T) {
	type test struct {
		cols     []string
		start    uint
		expected string
	}

	tests := []test{
		{cols: []string{}, start: 1, expected: ""},
		{cols: []string{"a", "b", "c"}, start: 1, expected: "($1,$2,$3)"},
		{cols: []string{"a", "b", "c"}, start: 2, expected: "($2,$3,$4)"},
		{cols: []string{"a", "b"}, start: 5, expected: "($5,$6)"},
	}

	for i := range tests {
		test := &tests[i]
		result := ClauseValues(Dollar, test.start, uint(len(test.cols)))

		if result != test.expected {
			t.Logf("Unexpected values")
			t.Logf("Expecting: \"%s\"", test.expected)
			t.Logf("Got: \"%s\"", result)

			t.Fatalf("unexpected values")
		}
	}
}

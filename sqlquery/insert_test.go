package sqlquery

import "testing"

func TestInsertAll(t *testing.T) {
	foos := []*foo{
		{id: 0, name: "a", age: 1},
		{id: 1, name: "b", age: 2},
		{id: 2, name: "c", age: 3},
	}

	items := make([]ModelCreate, len(foos))
	for i := range items {
		items[i] = foos[i]
	}

	query, values, err := InsertAll(Colon, items...)
	if err != nil {
		t.Errorf("mapCreates returns error: %s", err.Error())
	}

	expectedQuery := "insert all into FOO (ID,NAME,AGE) values"
	expectedQuery += " (:1,:2,:3)"
	expectedQuery += " (:4,:5,:6)"
	expectedQuery += " (:7,:8,:9)"

	if query != expectedQuery {
		t.Logf("Unexpected query")
		t.Logf("Expecting:\n\"%s\"", expectedQuery)
		t.Logf("Actual:\n\"%s\"", query)

		t.Fatalf("unexpected query")
	}

	expectedValues := []interface{}{0, "a", 1, 1, "b", 2, 2, "c", 3}

	if lenExpected, lenActual := len(expectedValues), len(values); lenExpected != lenActual {
		t.Logf("Unexpected len of bind vars")
		t.Logf("Expecting: %d", lenExpected)
		t.Logf("Actual: %d", lenActual)

		t.Fatalf("unexpected values")
	}
}

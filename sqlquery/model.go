package sqlquery

type ModelBase interface {
	TableName() string
}

type ModelWhere interface {
	ModelBase

	Where() map[string]interface{}
}

type ModelCreate interface {
	ModelBase

	ColumnsCreate() []string
	ValuesCreate() []interface{}
	MapColumnValuesCreate() map[string]interface{}
}

type ModelUpdate interface {
	ModelBase

	ColumnsUpdate() []string
	ValuesUpdate() []interface{}
	MapColumnValuesUpdate() map[string]interface{}
}

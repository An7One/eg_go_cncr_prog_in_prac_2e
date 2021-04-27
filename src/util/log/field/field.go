package field

// go:generate log_xfields_generator

// FieldType the type of field
type FieldType int

// constant fields of logs
const (
	UnkownType  FieldType = 0
	BoolType    FieldType = 1
	Int64Type   FieldType = 2
	Float64Type FieldType = 3
	StringType  FieldType = 4
	ObjectType  FieldType = 5
)

// Field interface of log field
type Field interface {
	Name() string
	Type() FieldType
	Value() interface{}
}

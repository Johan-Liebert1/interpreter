package callstack

type ActivationRecord struct {
	Name         string
	Type         string
	NestingLevel int
	Members      map[string]interface{}
}

func (ar *ActivationRecord) Init(name string, _type string, nestingLevel int) {
	ar.Name = name
	ar.Type = _type
	ar.NestingLevel = nestingLevel
	ar.Members = map[string]interface{}{}
}

func (ar *ActivationRecord) SetItem(key string, value interface{}) {
	ar.Members[key] = value
}

func (ar *ActivationRecord) GetItem(key string) (interface{}, bool) {
	value, exists := ar.Members[key]

	return value, exists
}

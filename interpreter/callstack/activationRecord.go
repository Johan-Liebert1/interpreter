package callstack

type ActivationRecord struct {
	Name         string
	Type         string
	NestingLevel int
	Members      map[string]interface{}
	AboveNode    *ActivationRecord
}

func (ar *ActivationRecord) Init() {
	ar.Members = map[string]interface{}{}
}

func (ar *ActivationRecord) SetItem(key string, value interface{}) {
	// helpers.ColorPrint(constants.LightMagenta, 1, 1, "setting key = ", key, " value = ", value)

	ar.Members[key] = value
}

func (ar *ActivationRecord) GetItem(key string) (interface{}, bool) {
	value, exists := ar.Members[key]

	if !exists && ar.AboveNode != nil {
		value, exists = ar.AboveNode.GetItem(key)
	}

	return value, exists
}

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

func (ar *ActivationRecord) GetActivaionRecordWithKey(key string) *ActivationRecord {
	result := ar

	if ar.AboveNode != nil {
		result = ar.AboveNode.GetActivaionRecordWithKey(key)
	}

	return result
}

func (ar *ActivationRecord) SetItem(key string, value interface{}) {
	// helpers.ColorPrint(constants.LightMagenta, 1, 0, "setting key = ", key, " value = ", value)

	arToSet := ar.GetActivaionRecordWithKey(key)

	arToSet.Members[key] = value
}

func (ar *ActivationRecord) GetItem(key string) (interface{}, bool) {
	value, exists := ar.Members[key]

	if !exists && ar.AboveNode != nil {
		value, exists = ar.AboveNode.GetItem(key)
	}

	// helpers.ColorPrint(constants.LightMagenta, 1, 0, "get key = ", key, " value = ", value)

	return value, exists
}

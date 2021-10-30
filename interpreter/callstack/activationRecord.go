package callstack

import (
	"programminglang/constants"
)

/*
Members = {
	varName: {
		varValue: interface{},
		varType: string,
	}
}
*/

type ActivationRecord struct {
	Name         string
	Type         string
	NestingLevel int
	Members      map[string]map[string]interface{}
	AboveNode    *ActivationRecord
}

func (ar *ActivationRecord) Init() {
	ar.Members = map[string]map[string]interface{}{}
}

func (ar *ActivationRecord) GetActivaionRecordWithKey(key string) *ActivationRecord {
	result := ar

	if ar.AboveNode != nil {
		result = ar.AboveNode.GetActivaionRecordWithKey(key)
	}

	return result
}

func (ar *ActivationRecord) SetItem(key string, value map[string]interface{}, setVarType bool) {

	arToSet := ar.GetActivaionRecordWithKey(key)

	var (
		typeToSet  interface{}
		valueToSet interface{} = value[constants.AR_KEY_VALUE]
	)

	// helpers.ColorPrint(constants.LightMagenta, 1, 1, constants.SpewPrinter.Sdump(key, value, setVarType))

	// set variable type at var decl, loop counter and function params
	if setVarType {
		typeToSet = value[constants.AR_KEY_TYPE]
	} else {
		typeToSet = arToSet.Members[key][constants.AR_KEY_TYPE]
	}

	finalValue := map[string]interface{}{
		constants.AR_KEY_TYPE:  typeToSet,
		constants.AR_KEY_VALUE: valueToSet,
	}

	// helpers.ColorPrint(constants.LightMagenta, 1, 0, "setting key = ", key, " value = ", finalValue)

	arToSet.Members[key] = finalValue
}

func (ar *ActivationRecord) GetItem(key string) (map[string]interface{}, bool) {
	value, exists := ar.Members[key]

	if !exists && ar.AboveNode != nil {
		value, exists = ar.AboveNode.GetItem(key)
	}

	// helpers.ColorPrint(constants.LightMagenta, 1, 0, "get key = ", key, " value = ", value)

	return value, exists
}

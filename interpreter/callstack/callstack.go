package callstack

type CallStack struct {
	Records []ActivationRecord
}

func (cs *CallStack) Push(item ActivationRecord) {
	cs.Records = append(cs.Records, item)

	// helpers.ColorPrint(constants.LightGreen, 1, 1, "Push = ", constants.SpewPrinter.Sdump(cs.Records))

}

func (cs *CallStack) Pop() ActivationRecord {
	// helpers.ColorPrint(constants.LightGreen, 1, constants.SpewPrinter.Sdump(cs))

	var poppedItem ActivationRecord

	if len(cs.Records) == 0 {
		return poppedItem
	}

	poppedItem, cs.Records = cs.Records[len(cs.Records)-1], cs.Records[:len(cs.Records)-1]

	// helpers.ColorPrint(constants.LightGreen, 1, 1, "Pop = ", constants.SpewPrinter.Sdump(cs.Records))

	return poppedItem
}

func (cs *CallStack) Peek() (ActivationRecord, bool) {
	if len(cs.Records) == 0 {
		return ActivationRecord{}, false
	}

	return cs.Records[len(cs.Records)-1], true
}

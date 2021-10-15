package callstack

type CallStack struct {
	Records []ActivationRecord
}

func (cs *CallStack) Push(item ActivationRecord) {
	cs.Records = append(cs.Records, item)
}

func (cs *CallStack) Pop() ActivationRecord {
	var poppedItem ActivationRecord

	poppedItem, cs.Records = cs.Records[len(cs.Records)-1], cs.Records[:len(cs.Records)-1]

	return poppedItem
}

func (cs *CallStack) Peek() ActivationRecord {
	return cs.Records[len(cs.Records)-1]
}

package item

// String is an item representation of Cobweb when Cobweb is broken by an unsuitable tool.
type String struct {
}

// EncodeItem...
func (String) EncodeItem() (id int32, meta int16) {
	return 287, 0
}

// TODO: In the future, when tripwire hook is added, string should be treated as a block and an item, and be in package block instead.
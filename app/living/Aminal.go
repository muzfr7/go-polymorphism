package living

// Animal ...
type Animal struct {
	Name string
	Age  uint8
}

// HasName ...
func (a Animal) HasName() bool {
	return true
}

// TellName ...
func (a Animal) TellName() string {
	return a.Name
}

// CanTalk ...
func (a Animal) CanTalk() bool {
	return true
}

// CanWalk ...
func (a Animal) CanWalk() bool {
	return true
}

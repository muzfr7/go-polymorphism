package living

// Person ...
type Person struct {
	Name string
	Age  uint8
}

// HasName ...
func (p Person) HasName() bool {
	return true
}

// TellName ...
func (p Person) TellName() string {
	return p.Name
}

// CanTalk ...
func (p Person) CanTalk() bool {
	return true
}

// CanWalk ...
func (p Person) CanWalk() bool {
	return true
}

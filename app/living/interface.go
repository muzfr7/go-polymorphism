package living

// Speakable ...
type Speakable interface {
	HasName() bool
	TellName() string
	CanTalk() bool
	CanWalk() bool
}

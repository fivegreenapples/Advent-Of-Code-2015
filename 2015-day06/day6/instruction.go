package day6

// Action describes how to change the status of a lightbulb
type Action int

const (
	// On turns a light on
	On Action = iota
	// Off turns a light off
	Off
	// Toggle changes a bulb from on to off or vice-versa
	Toggle
)

// Instruction describes which bulbs need to change and
// what action to take
type Instruction struct {
	Action      Action
	TopLeft     TwoD
	BottomRight TwoD
}

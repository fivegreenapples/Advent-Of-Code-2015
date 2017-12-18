package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type sourceType string
type gateType string

const (
	WIRE  sourceType = "wire"
	GATE             = "gate"
	VALUE            = "value"
)
const (
	AND    gateType = "and"
	OR              = "or"
	NOT             = "not"
	RSHIFT          = "rshift"
	LSHIFT          = "lshift"
)

type wireIdentifier string
type gateIdentifier int

type wire struct {
	source          sourceType
	sourceWire      wireIdentifier
	sourceGate      gateIdentifier
	sourceValue     uint16
	value           uint16
	valueDetermined bool
}

type gate struct {
	typ             gateType
	sourceA         wireIdentifier
	sourceB         wireIdentifier
	shiftAmount     uint
	value           uint16
	valueDetermined bool
}

type circuit struct {
	wires         map[wireIdentifier]*wire
	gates         map[gateIdentifier]*gate
	nextGateIndex gateIdentifier
}

func makeCircuit() circuit {
	return circuit{
		wires: map[wireIdentifier]*wire{},
		gates: map[gateIdentifier]*gate{},
	}
}

func (c *circuit) parseAllInstructions(in []string) {
	for _, i := range in {
		c.parseInstruction(i)
	}
}
func (c *circuit) parseInstruction(in string) {
	/*
		"bn RSHIFT 2 -> bo",
		"lf RSHIFT 1 -> ly",
		"fo RSHIFT 3 -> fq",
		"cj OR cp -> cq",
		"fo OR fz -> ga",
		"t OR s -> u",
		"lx -> a",
		"NOT ax -> ay",
		"he RSHIFT 2 -> hf",
	*/
	gateRegex := regexp.MustCompile(`^([a-z0-9]+) ([A-Z]+) ([a-z0-9]+) -> ([a-z]+)$`)
	notGateRegex := regexp.MustCompile(`^NOT ([a-z]+) -> ([a-z]+)$`)
	wireRegex := regexp.MustCompile(`^([a-z]+) -> ([a-z]+)$`)
	valueRegex := regexp.MustCompile(`^([0-9]+) -> ([a-z]+)$`)
	wireIsValueRegex := regexp.MustCompile(`^[0-9]+$`)

	if matches := gateRegex.FindStringSubmatch(in); matches != nil {
		sourceA := matches[1]
		sourceB := matches[3]
		gateTyp := matches[2]
		sink := matches[4]
		newGate := gate{}
		switch gateTyp {
		case "AND":
			newGate.sourceA = wireIdentifier(sourceA)
			newGate.sourceB = wireIdentifier(sourceB)
			newGate.typ = AND
		case "OR":
			newGate.sourceA = wireIdentifier(sourceA)
			newGate.sourceB = wireIdentifier(sourceB)
			newGate.typ = OR
		case "RSHIFT", "LSHIFT":
			shiftAmount, err := strconv.Atoi(sourceB)
			if err != nil {
				panic("shift amount failed parsing: " + err.Error())
			}
			newGate.sourceA = wireIdentifier(sourceA)
			newGate.shiftAmount = uint(shiftAmount)
			if gateTyp == "RSHIFT" {
				newGate.typ = RSHIFT
			} else {
				newGate.typ = LSHIFT
			}
		default:
			panic("gate not understood: " + gateTyp)
		}
		c.gates[c.nextGateIndex] = &newGate
		c.wires[wireIdentifier(sink)] = &wire{
			source:     GATE,
			sourceGate: c.nextGateIndex,
		}
		c.nextGateIndex++

		// we now check if sourceA is actually a value
		if wireIsValueRegex.MatchString(sourceA) {
			// if so we make a pretend wire using the value as its identifier
			wireVal, err := strconv.Atoi(sourceA)
			if err != nil {
				panic("failed parsing value of gate input: " + err.Error())
			}
			c.wires[wireIdentifier(sourceA)] = &wire{
				source:      VALUE,
				sourceValue: uint16(wireVal),
			}
		}

		return
	}
	if matches := notGateRegex.FindStringSubmatch(in); matches != nil {
		source := matches[1]
		sink := matches[2]
		newGate := gate{}
		newGate.typ = NOT
		newGate.sourceA = wireIdentifier(source)
		c.gates[c.nextGateIndex] = &newGate
		c.wires[wireIdentifier(sink)] = &wire{
			source:     GATE,
			sourceGate: c.nextGateIndex,
		}
		c.nextGateIndex++
		return
	}
	if matches := wireRegex.FindStringSubmatch(in); matches != nil {
		source := matches[1]
		sink := matches[2]
		c.wires[wireIdentifier(sink)] = &wire{
			source:     WIRE,
			sourceWire: wireIdentifier(source),
		}
		return
	}
	if matches := valueRegex.FindStringSubmatch(in); matches != nil {
		strVal := matches[1]
		sink := matches[2]
		val, err := strconv.Atoi(strVal)
		if err != nil {
			panic("value instruction failed parsing: " + err.Error())
		}
		c.wires[wireIdentifier(sink)] = &wire{
			source:      VALUE,
			sourceValue: uint16(val),
		}
		return
	}
	panic("instruction not understood: " + in)
}
func (c *circuit) valueForWire(w wireIdentifier) uint16 {
	wire, found := c.wires[w]
	if !found {
		panic("wire not found in circuit: " + w)
	}

	if wire.valueDetermined {
		return wire.value
	}

	switch wire.source {
	case VALUE:
		wire.value = uint16(wire.sourceValue)
	case WIRE:
		wire.value = c.valueForWire(wire.sourceWire)
	case GATE:
		wire.value = c.valueForGate(wire.sourceGate)
	default:
		panic("unhandled wire source type: " + wire.source)
	}
	wire.valueDetermined = true
	c.wires[w] = wire
	return wire.value
}
func (c *circuit) valueForGate(g gateIdentifier) uint16 {
	gate, found := c.gates[g]
	if !found {
		panic("gate not found in circuit: " + strconv.Itoa(int(g)))
	}

	if gate.valueDetermined {
		return gate.value
	}

	switch gate.typ {
	case AND:
		gate.value = c.valueForWire(gate.sourceA) & c.valueForWire(gate.sourceB)
	case OR:
		gate.value = c.valueForWire(gate.sourceA) | c.valueForWire(gate.sourceB)
	case RSHIFT:
		gate.value = c.valueForWire(gate.sourceA) >> gate.shiftAmount
	case LSHIFT:
		gate.value = c.valueForWire(gate.sourceA) << gate.shiftAmount
	case NOT:
		gate.value = ^c.valueForWire(gate.sourceA)
	default:
		panic("unhandled gate type: " + gate.typ)
	}
	gate.valueDetermined = true
	c.gates[g] = gate
	return gate.value
}

func (c *circuit) reset() {
	for g := range c.gates {
		c.gates[g].valueDetermined = false
	}
	for w := range c.wires {
		c.wires[w].valueDetermined = false
	}
}
func (c *circuit) setWireToValue(w wireIdentifier, value uint16) {
	c.wires[w] = &wire{
		source:      VALUE,
		sourceValue: value,
	}
}

func (c circuit) String() string {
	// First find all wires and sort alphabetically
	wires := []wireIdentifier{}
	for w := range c.wires {
		wires = append(wires, w)
	}
	sort.Slice(wires, func(i int, j int) bool {
		return wires[i] < wires[j]
	})

	// now loop over wires, outputting values
	out := ""
	for _, w := range wires {
		out += fmt.Sprintf("%s: %d\n", w, c.valueForWire(w))
	}
	return out
}

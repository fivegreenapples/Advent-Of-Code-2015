package day6

type grid struct {
	lights [][]bool
	numOn  int
}

// MakeGrid returns a lighting grid of the given width and height
// with all lights initially turned off
func MakeGrid(width, height int) grid {
	g := grid{
		lights: make([][]bool, width),
		numOn:  0,
	}
	for w := 0; w < width; w++ {
		g.lights[w] = make([]bool, height)
	}
	return g
}

func (g grid) NumberOfLightsTurnedOn() int {
	return g.numOn
}
func (g *grid) ApplyInstructionList(ii []Instruction) {
	for _, i := range ii {
		g.ApplyInstruction(i)
	}
}

func (g *grid) ApplyInstruction(i Instruction) {

	for x := i.TopLeft.X; x <= i.BottomRight.X; x++ {
		for y := i.TopLeft.Y; y <= i.BottomRight.Y; y++ {
			switch i.Action {
			case On:
				if !g.lights[x][y] {
					g.lights[x][y] = true
					g.numOn++
				}
			case Off:
				if g.lights[x][y] {
					g.lights[x][y] = false
					g.numOn--
				}
			case Toggle:
				g.lights[x][y] = !g.lights[x][y]
				if g.lights[x][y] {
					g.numOn++
				} else {
					g.numOn--
				}
			default:
				panic("unhandled action")
			}
		}
	}

}

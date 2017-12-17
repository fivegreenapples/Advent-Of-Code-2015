package day6

type grid2 struct {
	lights          [][]int
	totalBrightness int
}

// MakeGrid2 returns a lighting grid of the given width and height
// with all lights initially turned off
func MakeGrid2(width, height int) grid2 {
	g := grid2{
		lights:          make([][]int, width),
		totalBrightness: 0,
	}
	for w := 0; w < width; w++ {
		g.lights[w] = make([]int, height)
	}
	return g
}

func (g grid2) TotalBrightness() int {
	return g.totalBrightness
}
func (g *grid2) ApplyInstructionList(ii []Instruction) {
	for _, i := range ii {
		g.ApplyInstruction(i)
	}
}

func (g *grid2) ApplyInstruction(i Instruction) {

	for x := i.TopLeft.X; x <= i.BottomRight.X; x++ {
		for y := i.TopLeft.Y; y <= i.BottomRight.Y; y++ {
			switch i.Action {
			case On:
				g.lights[x][y]++
				g.totalBrightness++
			case Off:
				if g.lights[x][y] > 0 {
					g.lights[x][y]--
					g.totalBrightness--
				}
			case Toggle:
				g.lights[x][y] += 2
				g.totalBrightness += 2
			default:
				panic("unhandled action")
			}
		}
	}

}

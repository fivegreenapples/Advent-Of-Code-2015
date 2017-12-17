package main

type coord struct {
	x int
	y int
}

func (c *coord) applyDelta(d delta) {
	c.x += d.dX
	c.y += d.dY
}

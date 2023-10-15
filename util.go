package main

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var (
	black = Color{0, 0, 0, 255}
	white = Color{255, 255, 255, 255}
	gray  = Color{127, 127, 127, 255}
)

func setColor(pixels []uint8, p int, c Color) {
	pixels[p+0] = c.R
	pixels[p+1] = c.G
	pixels[p+2] = c.B
	pixels[p+3] = c.A
}

func screenCoordToField(x, y int) (xf, yf int) {
	return x / ratioW, y / ratioH
}

func getNeighbors(x, y int, field [][]uint8) uint8 {
	return field[minus(y-1, fieldHeight)][x] + // up
		field[plus(y+1, fieldHeight)][x] + // down
		field[y][minus(x-1, fieldWidth)] + // left
		field[y][plus(x+1, fieldWidth)] + // right
		field[minus(y-1, fieldHeight)][minus(x-1, fieldWidth)] + // up left
		field[minus(y-1, fieldHeight)][plus(x+1, fieldWidth)] + // up right
		field[plus(y+1, fieldHeight)][minus(x-1, fieldWidth)] + // down left
		field[plus(y+1, fieldHeight)][plus(x+1, fieldWidth)] // down right
}

func drawGrid(g *Game) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := (y*width + x) * 4
			if x%ratioW == 0 {
				setColor(g.pixels, p, gray)
			}
			if y%ratioH == 0 {
				setColor(g.pixels, p, gray)
			}
		}
	}
}

func minus(n, v int) int {
	return (n + v) % v
}

func plus(n, v int) int {
	return n % v
}

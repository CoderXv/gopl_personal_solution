package solution

import (
	"fmt"
)

const (
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
)

// 3.3
func Solution3() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	min, max := getMinMax()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, r, g, b := corner2(i+1, j, min, max)
			bx, by, r, g, b := corner2(i, j, min, max)
			cx, cy, r, g, b := corner2(i, j+1, min, max)
			dx, dy, r, g, b := corner2(i+1, j+1, min, max)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%x%x%x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Println("</svg>")
}

func getColor(min, max, current float64) (int, int, int) {
	step := (max - min) / 255
	v := int((current - min) / step)
	r := v
	g := 0
	b := 255 - v
	return r, g, b
}

func corner2(i, j int, min, max float64) (float64, float64, int, int, int) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	// 将(x,y,z)等角投射到二维SVG绘图平面上,坐标是(sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	r, g, b := getColor(min, max, z)
	return sx, sy, r, g, b
}

func getMinMax() (float64, float64) {
	var min, max float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)

			z := f(x, y)
			if z < min {
				min = z
			}
			if z > max {
				max = z
			}
		}
	}
	return min, max
}


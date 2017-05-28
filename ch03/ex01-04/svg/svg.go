package svg

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Config struct {
	Func          func(x, y float64) float64
	Width, Height int     // canvas size in pixels
	Cells         int     // number of grid cells
	XYRange       float64 // axis ranges (-xyrange..+xyrange)
	Angle         float64 // angle of x, y axes
	MaxColor      uint32
	MinColor      uint32

	xyscale  int // pixels per x or y unit
	zscale   int // pixels per z unit
	sinRange float64
	cosRange float64
}

func (c *Config) Validate() error {
	if c.Func == nil || c.Width < 1 || c.Height < 1 || c.Cells < 1 || c.XYRange == 0.0 {
		return fmt.Errorf("invalid config: %v", c)
	}
	c.xyscale = int(float64(c.Width) / 2 / c.XYRange)
	c.zscale = int(float64(c.Height) * 0.4)

	c.sinRange = math.Sin(c.Angle)
	c.cosRange = math.Cos(c.Angle)

	return nil
}

func SVG(w io.Writer, c *Config) error {
	if err := c.Validate(); err != nil {
		return err
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", c.Width, c.Height)

	calc := NewCalc(c)
	for i := 0; i < c.Cells; i++ {
		for j := 0; j < c.Cells; j++ {
			ax, ay, bx, by, cx, cy, dx, dy, color, err := calc.Polygon(i, j)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(w, "</svg>")
	return nil
}

type Calc struct {
	conf     *Config
	memo     [][]float64
	max, min float64
}

func NewCalc(conf *Config) *Calc {
	max := math.Inf(-1)
	min := math.Inf(1)
	memo := make([][]float64, conf.Width+1)
	for i := 0; i < conf.Cells+1; i++ {
		col := make([]float64, conf.Height+1)
		for j := 0; j < conf.Cells+1; j++ {
			// Find point (x,y) at corner of cell (i,j).
			x, y := toPoint(i, j, conf)

			// Compute surface height z.
			z := conf.Func(x, y)
			col[j] = z

			if z > max {
				max = z
			}
			if z < min {
				min = z
			}
		}
		memo[i] = col
	}
	return &Calc{memo: memo, max: max, min: min, conf: conf}
}
func (c *Calc) Get(i, j int) float64 {
	if i > c.conf.Cells || j > c.conf.Cells {
		return math.NaN()
	}
	return c.memo[i][j]
}
func (c *Calc) Corner(i, j int) (float64, float64, error) {
	x, y := toPoint(i, j, c.conf)

	z := c.Get(i, j)
	if math.IsInf(z, 0) {
		return 0, 0, fmt.Errorf("Func(%g, %g) is infinity", x, y)
	}
	if math.IsNaN(z) {
		return 0, 0, fmt.Errorf("Func(%g, %g) is NaN", x, y)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(c.conf.Width)/2 + (x-y)*c.conf.cosRange*float64(c.conf.xyscale)
	sy := float64(c.conf.Height)/2 + (x+y)*c.conf.sinRange*float64(c.conf.xyscale) - z*float64(c.conf.zscale)
	return sx, sy, nil
}
func (c *Calc) Polygon(i, j int) (ax, ay, bx, by, cx, cy, dx, dy float64, color string, err error) {
	ax, ay, err = c.Corner(i+1, j)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, "", err
	}
	bx, by, err = c.Corner(i, j)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, "", err
	}
	cx, cy, err = c.Corner(i, j+1)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, "", err
	}
	dx, dy, err = c.Corner(i+1, j+1)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, "", err
	}

	return ax, ay, bx, by, cx, cy, dx, dy, c.Color(i, j), nil
}

func (c *Calc) Color(i, j int) string {
	normal := (c.Get(i, j) - c.min) / (c.max - c.min)
	maxcol := normalizeColor(c.conf.MaxColor, normal)
	mincol := normalizeColor(c.conf.MinColor, 1-normal)
	return fmt.Sprintf("#%06x", maxcol+mincol)
}

func normalizeColor(color uint32, normal float64) uint32 {
	r := float64(0xff0000 & color)
	g := float64(0x00ff00 & color)
	b := float64(0x0000ff & color)
	return (uint32(r*normal))&0xff0000 + (uint32(g*normal))&0x00ff00 + (uint32(b*normal))&0x0000ff
}

func toPoint(i, j int, c *Config) (x, y float64) {
	// Find point (x,y) at corner of cell (i,j).
	x = c.XYRange * (float64(i)/float64(c.Cells) - 0.5)
	y = c.XYRange * (float64(j)/float64(c.Cells) - 0.5)
	return x, y
}

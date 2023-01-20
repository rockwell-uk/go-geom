package geom

import (
	"errors"
	"fmt"
	"math"

	geos "github.com/rockwell-uk/go-geos"
)

func Circle(origin []float64, radius float64, numPoints int) ([][]float64, error) {

	d := float64(360)

	var ai float64
	points := [][]float64{}
	ai = d / float64(numPoints) * math.Pi / 180

	c := 0.0
	for i := 0; i <= numPoints; i++ {
		c += ai
		if c > d {
			break
		}
		x := radius*math.Cos(c) + origin[0]
		y := radius*math.Sin(c) + origin[1]

		points = append(points, []float64{x, y})
	}

	n := len(points)
	if n == 0 {
		return [][]float64{}, errors.New("no points in circle")
	}

	return points, nil
}

func CircleWKT(origin []float64, radius float64, numPoints int) (string, error) {

	points, err := Circle(origin, radius, numPoints)
	if err != nil {
		return "LINESTRING EMPTY", err
	}

	n := len(points)
	s := "LINESTRING("
	for i, c := range points {
		s = fmt.Sprintf("%v%v %v", s, c[0], c[1])
		if i < n-1 {
			s = fmt.Sprintf("%v,", s)
		}
	}

	s = fmt.Sprintf("%v)", s)

	return s, nil
}

func CircleGeom(origin []float64, radius float64, numPoints int) (*geos.Geom, error) {

	e := geos.Geom{}

	r, err := CircleWKT(origin, radius, numPoints)
	if err != nil {
		return &e, err
	}

	g, err := gctx.NewGeomFromWKT(r)
	if err != nil {
		return &e, err
	}

	return g, nil
}

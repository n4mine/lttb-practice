package main

import (
	"./testdata"

	"github.com/devtoolkits/downsample"
	"github.com/dgryski/go-lttb"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const factor = 18

func main() {
	// downsmaple(avg)
	originPoints := testdata.OriginPoints2

	var ps = downsample.NewPoints()
	for _, p := range originPoints {
		ps = append(ps, downsample.Point{int64(p.X), p.Y})
	}

	avgPoints := ps.Downsample(10 * factor)

	// downsample(lttb)
	lttbPoints := lttb.LTTB(originPoints, len(originPoints)/factor)

	// plot
	p, _ := plot.New()

	// origin points
	var originPlotPoints plotter.XYs
	for _, point := range originPoints {
		originPlotPoints = append(originPlotPoints, plotter.XY{point.X, point.Y})
	}

	// downsample(avg) points
	var avgPlotPoints plotter.XYs
	for _, point := range avgPoints {
		avgPlotPoints = append(avgPlotPoints, plotter.XY{float64(point.Timestamp), point.Value})
	}

	// downsample(lttb) points
	var lttbPlotPoints plotter.XYs
	for _, point := range lttbPoints {
		lttbPlotPoints = append(lttbPlotPoints, plotter.XY{point.X, point.Y})
	}

	plotutil.AddLinePoints(p,
		"origin", originPlotPoints,
		//"avg", avgPlotPoints,
		"lttb", lttbPlotPoints,
	)

	p.Save(10*vg.Inch, 4*vg.Inch, "out.png")
}

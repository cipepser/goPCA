package singleRandNorm

import (
	"math/rand"
	"os/exec"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Srn() {
	rand.Seed(time.Now().UnixNano())

	v := make([]float64, 100000)
	for i := range v {
		v[i] = rand.NormFloat64()
	}

	// ここから下はヒストグラム表示
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	h, err := plotter.NewHist(plotter.Values(v), 16)
	if err != nil {
		panic(err)
	}

	h.Normalize(1)
	p.Add(h)

	file := "hist.png"
	if err := p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

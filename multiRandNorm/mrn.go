package multiRandomNorm

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/gonum/mat"
)

// MyScatter is a wrapper of Scatter of package plotter
// with slice of float64 x and y.
func MyScatter(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}

	s.Radius = vg.Length(2)

	p.Add(s)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// MultiNorm returns multi-dimension normally distributed VecDense
// with average vector u and covariance matrix S.
func MultiNorm(u *mat.VecDense, S *mat.SymDense) (*mat.VecDense, error) {
	rand.Seed(time.Now().UnixNano())

	n, _ := S.Dims()
	x := make([]float64, n)
	for i := range x {
		x[i] = rand.NormFloat64()
	}

	y := mat.NewVecDense(len(x), x)

	var chol mat.Cholesky
	if ok := chol.Factorize(S); !ok {
		return nil, fmt.Errorf("covariance matrix must be poositive defined")
	}

	var L mat.TriDense
	chol.LTo(&L)

	y.MulVec(&L, y)
	y.AddVec(y, u)

	return y, nil
}

func mrn() {
	N := 10000
	x1 := make([]float64, N)
	x2 := make([]float64, N)

	for i := 0; i < N; i++ {
		y, _ := MultiNorm(mat.NewVecDense(2, []float64{0.0, 0.0}),
			mat.NewSymDense(2, []float64{3.0, 0.5, 0.5, 1.0}),
		)

		x1[i] = y.At(0, 0)
		x2[i] = y.At(1, 0)
	}

	MyScatter(x1, x2)
}

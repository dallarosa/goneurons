package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	learnRate = 0.001
	featSize  = 2
	iterNum   = 15000
)

type Neuron struct {
	W    []float64
	b    float64
	X    Input
	y    []float64
	yHat []float64
}

type Input [][featSize]float64

func (n *Neuron) Initialize(input Input, output []float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n.W = make([]float64, featSize)
	n.yHat = make([]float64, len(output))
	for i, _ := range n.W {
		n.W[i] = r.Float64()
	}
	n.b = r.Float64()
	n.X = input
	n.y = output
}

func (n *Neuron) ForwardStep() {
	for i, v := range n.X {
		n.yHat[i] = 0
		for j := 0; j < featSize; j++ {
			n.yHat[i] += v[j] * n.W[j]
		}
		n.yHat[i] += n.b
	}
}

func (n *Neuron) Cost() float64 {
	var cost float64
	for i, v := range n.y {
		cost += math.Pow((v - n.yHat[i]), 2)
	}
	return cost / float64(len(n.y))
}

func (n *Neuron) BackwardStep() {
	for i, _ := range n.W {
		delta := 0.0
		for j, v := range n.y {
			delta += 2 * (v - n.yHat[j]) * n.X[j][i]
		}
		delta = delta / float64(len(n.y))
		n.W[i] += delta * learnRate
	}
	delta := 0.0
	for j, v := range n.y {
		delta += 2 * (v - n.yHat[j])
	}
	delta = delta / float64(len(n.y))
	n.b += delta * learnRate
}

func main() {
	n := new(Neuron)
	n.Initialize(Input{{1, 2}, {2, 3}, {5, 5}, {4, 3}, {2, 2}, {10, 10}, {6, 4}, {34, 12}}, []float64{3, 5, 10, 7, 4, 20, 10, 46})

	fmt.Println(n)

	for i := 0; i < iterNum; i++ {
		n.ForwardStep()
		n.BackwardStep()
		fmt.Println("Cost: ", n.Cost())
	}
	for i, v := range n.yHat {
		fmt.Println("yHat[", i, "]: ", v)
	}
	fmt.Println(n)

}

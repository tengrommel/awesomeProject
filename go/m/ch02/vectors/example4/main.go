package main

import (
	"fmt"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/matrix/mat64"
)

func main() {
	vectorA := mat64.NewVector(3, []float64{11.0, 5.2, -1.3})
	vectorB := mat64.NewVector(3, []float64{-7.2, 4.2, 5.1})

	dotProduct := mat64.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.f\n", dotProduct)
	vectorA.ScaleVec(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)
	normB := blas64.Nrm2(3, vectorB.RawVector())
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}

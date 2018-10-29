package testmatrix

import (
	"fmt"
	"lab1/algmatrix"
	"lab1/printmatrix"
	"lab1/randommatrix"
)

func UnitTests(ln int){
	var a [][] float64
	var b [] float64
	a = randommatrix.RandGenerationMatrix(a, ln)
	b = randommatrix.RandGenerationVect(b, ln)
	fmt.Println("Start matrixes A and B")
	printmatrix.PrintAwithB(a, b, ln)
	fmt.Println("Invert Matrix of A")
	printmatrix.PrintMatrix(algmatrix.InvertMatrix(a, ln), ln)
	fmt.Println("Multiply A to inv(A) and get 1")
	printmatrix.PrintMatrix(algmatrix.MultiplyMatrixMatrix(a, algmatrix.InvertMatrix(a, ln), ln), ln)
	fmt.Println("X vector")
	printmatrix.PrintVector(algmatrix.MultiplyMatrixVect(algmatrix.InvertMatrix(a, ln), b, ln), ln)
	fmt.Println("Multiply A to X and get B")
	printmatrix.PrintVector(algmatrix.MultiplyMatrixVect(a,
		algmatrix.MultiplyMatrixVect(algmatrix.InvertMatrix(a, ln), b, ln), ln), ln)
}

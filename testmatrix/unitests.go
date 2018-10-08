package testmatrix

import (
	"lab1/algmatrix"
	"lab1/printmatrix"
	"lab1/randommatrix"
)

func UnitTests(ln int){
	var a [][] float64
	var b [] float64
	a = randommatrix.RandGenerationMatrix(a, ln)
	b = randommatrix.RandGenerationVect(b, ln)
	printmatrix.PrintVector(b, ln)
	printmatrix.PrintVector(algmatrix.MultiplyMatrixVect(a, algmatrix.MultiplyVectMatrix(b, algmatrix.InvertMatrix(a, ln), ln), ln), ln)
}

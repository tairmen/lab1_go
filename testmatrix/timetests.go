package testmatrix

import (
	"fmt"
	"lab1/algmatrix"
	"lab1/randommatrix"
	"time"
)

func Testing(period, limit, ln, n int){
	fmt.Printf("%10s%20s%20s%20s%20s\n", "ln", "Multiply", "Invert", "Invert Method", "Gauss Method")
	for i:=ln; i <= limit; i+=period{
		fmt.Printf("%10d%20v%20v%20v%20v\n", i, TestMultiplyVectMatrix(n, i),
			TestInvertMatrix(n, i), TestInvertMatrixAlg(n, i), TestGaussAlg(n, i))
	}
}

func TestMultiplyVectMatrix(n int, ln int) time.Duration{
	var a [][] float64
	var b [] float64
	t1 := time.Now()
	for i:=0; i < n; i++{
		a = randommatrix.RandGenerationMatrix(a, ln)
		b = randommatrix.RandGenerationVect(b, ln)
		algmatrix.MultiplyVectMatrix(b, a, ln)
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	return  diff
}

func TestInvertMatrix(n int, ln int) time.Duration{
	var a [][] float64
	a = randommatrix.RandGenerationMatrix(a, ln)
	t1 := time.Now()
	for i:=0; i < n; i++{
		a = randommatrix.RandGenerationMatrix(a, ln)
		algmatrix.InvertMatrix(a, ln)
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	return  diff
}

func TestInvertMatrixAlg(n int, ln int) time.Duration{
	var a [][] float64
	var b [] float64
	t1 := time.Now()
	for i:=0; i < n; i++{
		a = randommatrix.RandGenerationMatrix(a, ln)
		b = randommatrix.RandGenerationVect(b, ln)
		algmatrix.MultiplyVectMatrix(b, algmatrix.InvertMatrix(a, ln), ln)
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	return  diff
}

func TestGaussAlg(n int, ln int) time.Duration{
	var a [][] float64
	var b [] float64
	t1 := time.Now()
	for i:=0; i < n; i++{
		a = randommatrix.RandGenerationMatrix(a, ln)
		b = randommatrix.RandGenerationVect(b, ln)
		algmatrix.GaussMethod(a, b, ln)
	}
	t2 := time.Now()
	diff := t2.Sub(t1)
	return  diff
}

package printmatrix

import "fmt"

func PrintAwithB(a[][] float64, b[] float64, ln int){
	for i:=0; i < ln; i++{
		for j:=0; j < ln; j++{
			fmt.Printf("%10.3f ", a[i][j])
		}
		fmt.Printf(" | %10.3f\n", b[i])
	}
	fmt.Printf("\n\n")
}

func PrintMatrix(a[][] float64, ln int){
	for i:=0; i < ln; i++{
		for j:=0; j < ln; j++{
			fmt.Printf("%10.3f ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("\n\n")
}

func PrintVector(b[] float64, ln int){
	for i:=0; i < ln; i++{
		fmt.Printf("%10.3f ", b[i])
	}
	fmt.Printf("\n\n")
}
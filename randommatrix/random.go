package randommatrix

import "math/rand"


func RandGenerationMatrix(a [][] float64, ln int) [][] float64{
	var buf [] float64
	for i:=0; i < ln; i++{
		for j:=0; j < ln; j++{
			if rand.Intn(2) == 1{
				buf = append(buf, rand.Float64())
			} else {
				buf = append(buf, -rand.Float64())
			}
		}
		a = append(a, buf)
	}
	return a
}

func RandGenerationVect(b [] float64, ln int) [] float64{
	for i:=0; i < ln; i++{
		if rand.Intn(2) == 1{
			b = append(b, rand.Float64())
		} else {
			b = append(b, -rand.Float64())
		}
	}
	return b
}

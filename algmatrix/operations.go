package algmatrix

func MultiplyVectMatrix(b[] float64, a[][] float64, ln int) [] float64 {
	var mp[] float64
	var el float64
	for i:=0; i < ln; i++{
		el = 0
		for j:=0; j < ln; j++{
			el += b[j] * a[i][j]
		}
	}
	mp = append(mp, el)
	return mp
}

func MultiplyMatrixVect(a[][] float64, b[] float64, ln int) [] float64 {
	var el float64
	var mp [] float64
	for i:=0; i < ln; i++{
		el = 0
		for j:=0; j < ln; j++{
			el += b[j] * a[i][j]
		}
	}
	mp = append(mp, el)
	return mp
}

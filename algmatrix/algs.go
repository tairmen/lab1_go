package algmatrix

func GaussMethod(a [][] float64, b[] float64, ln int) [] float64{
	var del float64
	var buf float64
	for k:=0; k < ln; k++{
		del = a[k][k]
		for j:=0; j < ln; j++{
			a[k][j] = a[k][j] / del
		}
		b[k] = b[k] / del
		for i:=0; i < ln; i++{
			if i != k{
				buf = a[i][k]
				for j:=0; j < ln; j++{
					a[i][j] = a[i][j] - a[k][j] * buf
				}
				b[i] = b[i]  - b[k] * buf
			}
		}
	}
	return b
}

func InvertMatrix(a[][] float64, ln int) [][] float64{
	var ia [][] float64
	var bufia [] float64
	for i:=0; i < ln; i++{
		for j:=0; j < ln; j++{
			if i == j{
				bufia = append(bufia, 1)
			} else {
				bufia = append(bufia, 0)
			}
		}
		ia = append(ia, bufia)
		bufia = bufia[:0]
	}
	// PrintMatrix(ia)
	var del float64
	var buf float64
	for k:=0; k < ln; k++{
		del = a[k][k]
		for j:=0; j < ln; j++{
			ia[k][j] = ia[k][j] / del
			a[k][j] = a[k][j] / del
		}
		for i:=0; i < ln; i++{
			buf = a[i][k]
			if i != k{
				for j:=0; j < ln; j++{
					ia[i][j] = ia[i][j] - ia[k][j] * buf
					a[i][j] = a[i][j]  - a[k][j] * buf
				}
			}
		}
		// PrintMatrix(ia)
	}
	return ia
}
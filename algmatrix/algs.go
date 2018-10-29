package algmatrix

func GaussMethod(a [][] float64, b[] float64, ln int) [] float64{
	var aa [][] float64
	var bufia [] float64
	for i:=0; i < ln; i++{
		for j:=0; j < ln; j++{
			bufia = append(bufia, a[i][j])
		}
		aa = append(aa, bufia)
		bufia = nil
	}
	var del float64
	var buf float64
	for k:=0; k < ln; k++{
		del = aa[k][k]
		for j:=0; j < ln; j++{
			aa[k][j] = aa[k][j] / del
		}
		b[k] = b[k] / del
		for i:=0; i < ln; i++{
			if i != k{
				buf = aa[i][k]
				for j:=0; j < ln; j++{
					aa[i][j] = aa[i][j] - aa[k][j] * buf
				}
				b[i] = b[i]  - b[k] * buf
			}
		}
	}
	return b
}

func InvertMatrix(a[][] float64, ln int) [][] float64{
	var aa [][] float64
	var bufia [] float64
	for i:=0; i < ln; i++{
		for j:=0; j < ln; j++{
			bufia = append(bufia, a[i][j])
		}
		aa = append(aa, bufia)
		bufia = nil
	}
	var ia [][] float64
	for i:=0; i < ln; i++{
		for j:=0; j < ln; j++{
			if i == j{
				bufia = append(bufia, 1)
			} else {
				bufia = append(bufia, 0)
			}
		}
		ia = append(ia, bufia)
		bufia = nil
	}
	// printmatrix.PrintMatrix(ia, ln)
	var del float64
	var buf float64
	for k:=0; k < ln; k++{
		del = aa[k][k]
		for j:=0; j < ln; j++{
			ia[k][j] = ia[k][j] / del
			aa[k][j] = aa[k][j] / del
		}
		for i:=0; i < ln; i++{
			buf = aa[i][k]
			if i != k{
				for j:=0; j < ln; j++{
					ia[i][j] = ia[i][j] - ia[k][j] * buf
					aa[i][j] = aa[i][j]  - aa[k][j] * buf
				}
			}
		}
		// printmatrix.PrintMatrix(ia, ln)
		// printmatrix.PrintMatrix(aa, ln)
	}
	return ia
}
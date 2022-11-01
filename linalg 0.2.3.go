package linalg

/*Version 0.2.3*/
/*More functions added, bug fixes*/

type Vector []complex128

type Matrix [][]complex128

type Matrix3D [3][3]complex128

type Matrix2D [2][2]complex128

type Vector3D [3]complex128

type Vector2D [2]complex128

func Add(v Vector, v2 Vector) Vector {
	if len(v) == len(v2) {
		var newVector Vector
		for i := 0; i < len(v); i += 1 {
			newVector = append(newVector, v[i]+v2[i])
		}
		return newVector
	} else {
		var newVector Vector
		return newVector
	}
}

func ConstMul(v Vector, constant complex128) Vector {
	for i := 0; i < len(v); i += 1 {
		v[i] *= constant
	}
	return v
}

func Dot(v Vector, v2 Vector) complex128 {
	sum := 0 + 0i
	for i := range v {
		sum += v[i] * v2[i]
	}
	return sum
}

func (v Vector3D) Cross(v2 Vector3D) Vector3D {
	var result Vector3D
	result[0] = v[1]*v2[2] - v[2]*v2[1]
	result[1] = v[2]*v2[0] - v[0]*v2[2]
	result[2] = v[0]*v2[1] - v[1]*v2[0]
	return result
}

func FCross(v Vector3D, v2 Vector3D) Vector3D {
	var result Vector3D
	result[0] = v[1]*v2[2] - v[2]*v2[1]
	result[1] = v[2]*v2[0] - v[0]*v2[2]
	result[2] = v[0]*v2[1] - v[1]*v2[0]
	return result
}

func Mul(m Matrix, v Vector) Vector {
	var newVector Vector
	for i := 0; i < len(v); i += 1 {
		newVector = append(newVector, 0+0i)
	}
	for j := 0; j < len(newVector); j += 1 {
		newVector[j] = Dot(m[j], v)
	}
	return newVector
}

func (m Matrix) Add(m2 Matrix) Matrix {
	for i := 0; i < len(m); i += 1 {
		for j := 0; j < len(m[i]); j += 1 {
			m[i][j] += m2[i][j]
		}
	}
	return m
}

func (m Matrix) Mul(m2 Matrix) Matrix { /*m is on left if expressed in ordinary notation*/
	var newMatrix Matrix = [][]complex128{{}}
	for i := 0; i < len(m2); i += 1 {
		newMatrix = append(newMatrix, []complex128{0 + 0i})
		for j := 0; j < len(m)-1; j += 1 {
			newMatrix[i] = append(newMatrix[i], 0+0i)
		}
	} /*Properly initialises result matrix*/
	for i := 0; i < len(m2); i += 1 {
		for j := 0; j < len(m); j += 1 {
			newMatrix[i][j] = 0
		}
	}

	for i := 0; i < len(m); i += 1 {
		for j := 0; j < len(m2); j += 1 {
			sum := 0 + 0i
			for k := 0; k < len(m2[j]); k += 1 {
				sum += m[i][k] * m2[k][j]
			}
			newMatrix[i][j] = sum
		}
	}
	return newMatrix
}

func Transpose(m Matrix) Matrix {
	var newMatrix Matrix = [][]complex128{{}}
	n := len(m)
	for i := 0; i < n; i += 1 {
		newMatrix = append(newMatrix, []complex128{0 + 0i})
		for j := 0; j < n-1; j += 1 {
			newMatrix[i] = append(newMatrix[i], 0+0i)
		}
	} /*Properly initialises result matrix*/
	for i := 0; i < len(m); i += 1 {
		for j := 0; j < len(m[i]); j += 1 {
			newMatrix[i][j] = m[j][i]
		}
	}
	return newMatrix
}

func Det3D(m Matrix3D) complex128 {
	a := m[0][0]
	b := m[1][0]
	c := m[2][0]
	d := m[0][1]
	e := m[1][1]
	f := m[2][1]
	g := m[0][2]
	h := m[1][2]
	i := m[2][2]
	return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i - a*f*h
}

func Det2D(m Matrix2D) complex128 {
	a := m[0][0]
	b := m[1][0]
	c := m[0][1]
	d := m[1][1]
	return a*d - b*c
}

func Det(m Matrix, n int) complex128 {
	det := 0 + 0i
	var h int
	var k int
	var i int
	var j int
	var temp Matrix = [][]complex128{{}}
	for i := 0; i < n; i += 1 {
		temp = append(temp, []complex128{0 + 0i})
		for j := 0; j < n-1; j += 1 {
			temp[i] = append(temp[i], 0+0i)
		}
	} /*Properly initialises result matrix*/
	if n == 1 {
		return m[0][0]
	} else if n == 2 {
		a := m[0][0]
		b := m[1][0]
		c := m[0][1]
		d := m[1][1]
		return a*d - b*c
	} else if n == 3 {
		a := m[0][0]
		b := m[1][0]
		c := m[2][0]
		d := m[0][1]
		e := m[1][1]
		f := m[2][1]
		g := m[0][2]
		h := m[1][2]
		i := m[2][2]
		return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i - a*f*h
	} else {
		for p := 0; p < n; p += 1 {
			h = 0
			k = 0
			for i = 1; i < n; i += 1 {
				for j = 0; j < n; j += 1 {
					if j == p {
						continue
					}
					temp[h][k] = m[i][j]
					k++
					if k == n-1 {
						h++
						k = 0
					}
				}
			}
			det = det + m[0][p]*pow(p)*Det(temp, n-1)
		}
		return det
	}
}

func Adj(m Matrix) {
	/*TODO*/
}

func Inv(m Matrix) {
	/*TODO*/
}

func Trace(m Matrix) complex128 {
	result := 0 + 0i
	for i := 0; i < len(m); i += 1 {
		result += m[i][i]
	}
	return result
}

func pow(n int) complex128 {
	if n%2 == 0 {
		return 1 + 0i
	} else {
		return -1 + 0i
	}
}

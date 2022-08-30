package linalg

/*Version 0.2.0*/
/*Type system changed (more concise) and Matrix Multiplication added*/

type Vector []complex128

type Matrix [][]complex128

type Matrix3D [3][3]complex128

type Matrix2D [2][2]complex128

type Vector3D [3]complex128

type Vector2D [2]complex128

func (v Vector) Add(v2 Vector) Vector {
	if len(v) == len(v2) {
		var newVector Vector
		for i := 0; i < len(v); i += 1 {
			newVector[i] = v[i] + v2[i]
		}
		return newVector
	} else {
		var newVector Vector
		return newVector
	}
}

func FAdd(v Vector, v2 Vector) Vector {
	if len(v) == len(v2) {
		var newVector Vector
		for i := 0; i < len(v); i += 1 {
			newVector[i] = v[i] + v2[i]
		}
		return newVector
	} else {
		var newVector Vector
		return newVector
	}
}

func (v Vector) ConstMul(constant complex128) Vector {
	for i := 0; i < len(v); i += 1 {
		v[i] *= constant
	}
	return v
}

func FConstMul(v Vector, constant complex128) Vector {
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

func Cross(v Vector3D, v2 Vector3D) Vector3D {
	var result Vector3D
	result[0] = v[1]*v2[2] - v[2]*v2[1]
	result[1] = v[2]*v2[0] - v[0]*v2[2]
	result[2] = v[0]*v2[1] - v[1]*v2[0]
	return result
}

func Mul(m Matrix, v Vector) Matrix {
	var newMatrix Matrix
	for i := range m {
		newMatrix[i] = FConstMul(newMatrix[i], v[i])
	}
	return newMatrix
}

func (m Matrix) Add(m2 Matrix) Matrix {
	for i := range m {
		for j := 0; j < len(m[i]); j += 1 {
			m[i][j] += m2[i][j]
		}
	}
	return m
}

func (m Matrix) Mul(m2 Matrix) Matrix { /*m2 is on left if expressed in ordinary notation*/
	var newMatrix Matrix
	for i := 0; i < len(m); i += 1 {
		for j := 0; j < len(m2); j += 1 {
			newMatrix[i][j] = 0
		}
	}

	for i := 0; i < len(m); i += 1 {
		for j := 0; j < len(m2); j += 1 {
			sum := 0 + 0i
			for k := 0; k < len(m2); k += 1 {
				sum += m[j][i] * m2[i][j]
			}
			newMatrix[i][j] = sum
		}
	}
	return newMatrix
}

func (m Matrix3D) Det3D() complex128 {
	a := m[0][0]
	b := m[0][1]
	c := m[0][2]
	d := m[1][0]
	e := m[1][1]
	f := m[1][2]
	g := m[2][0]
	h := m[2][1]
	i := m[2][2]
	return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i - a*f*h
}

func FDet3D(m Matrix3D) complex128 {
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

func (m Matrix2D) Det2D() complex128 {
	a := m[0][0]
	b := m[1][0]
	c := m[0][1]
	d := m[1][1]
	return a*d - b*c
}

func FDet2D(m Matrix2D) complex128 {
	a := m[0][0]
	b := m[1][0]
	c := m[0][1]
	d := m[1][1]
	return a*d - b*c
}

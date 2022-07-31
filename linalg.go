package linalg

/*version: 0.1.0*/

type Vector struct {
	coordinates []complex128
}

type Matrix struct {
	vectors []Vector
}

type Vector3D struct {
	coordinates [3]complex128
}

type Vector2D struct {
	coordinates [2]complex128
}

type Matrix3D struct {
	vectors [3]Vector3D
}

type Matrix2D struct {
	vectors [2]Vector2D
}

func (v Vector) Add(v2 Vector) Vector {
	if len(v.coordinates) == len(v2.coordinates) {
		var newVector Vector
		for i := 0; i < len(v.coordinates); i += 1 {
			newVector.coordinates[i] = v.coordinates[i] + v2.coordinates[i]
		}
		return newVector
	} else {
		var newVector Vector
		return newVector
	}
}

func FAdd(v Vector, v2 Vector) Vector {
	if len(v.coordinates) == len(v2.coordinates) {
		var newVector Vector
		for i := 0; i < len(v.coordinates); i += 1 {
			newVector.coordinates[i] = v.coordinates[i] + v2.coordinates[i]
		}
		return newVector
	} else {
		var newVector Vector
		return newVector
	}
}

func (v Vector) ConstMul(constant complex128) Vector {
	for i := 0; i < len(v.coordinates); i += 1 {
		v.coordinates[i] *= constant
	}
	return v
}

func FConstMul(v Vector, constant complex128) Vector {
	for i := 0; i < len(v.coordinates); i += 1 {
		v.coordinates[i] *= constant
	}
	return v
}

func Dot(v Vector, v2 Vector) complex128 {
	sum := 0 + 0i
	for i := range v.coordinates {
		sum += v.coordinates[i] * v2.coordinates[i]
	}
	return sum
}

func Cross(v Vector3D, v2 Vector3D) Vector3D {
	var result Vector3D
	result.coordinates[0] = v.coordinates[1]*v2.coordinates[2] - v.coordinates[2]*v2.coordinates[1]
	result.coordinates[1] = v.coordinates[2]*v2.coordinates[0] - v.coordinates[0]*v2.coordinates[2]
	result.coordinates[2] = v.coordinates[0]*v2.coordinates[1] - v.coordinates[1]*v2.coordinates[0]
	return result
}

func Mul(m Matrix, v Vector) Matrix {
	var newMatrix Matrix
	for i := range m.vectors {
		newMatrix.vectors[i] = FConstMul(newMatrix.vectors[i], v.coordinates[i])
	}
	return newMatrix
}

func (m Matrix) Add(m2 Matrix) Matrix {
	for i := range m.vectors {
		for j := 0; j < len(m.vectors[i].coordinates); j += 1 {
			m.vectors[i].coordinates[j] += m2.vectors[i].coordinates[j]
		}
	}
	return m
}

/*func (m Matrix) Mul(m2 Matrix) Matrix {
	var newMatrix Matrix 
	
}*/

func (m Matrix3D) Det3D() complex128 {
	a := m.vectors[0].coordinates[0]
	b := m.vectors[0].coordinates[1]
	c := m.vectors[0].coordinates[2]
	d := m.vectors[1].coordinates[0]
	e := m.vectors[1].coordinates[1]
	f := m.vectors[1].coordinates[2]
	g := m.vectors[2].coordinates[0]
	h := m.vectors[2].coordinates[1]
	i := m.vectors[2].coordinates[2]
	return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i -  a*f*h
}

func FDet3D(m Matrix3D) complex128 {
	a := m.vectors[0].coordinates[0]
	b := m.vectors[1].coordinates[0]
	c := m.vectors[2].coordinates[0]
	d := m.vectors[0].coordinates[1]
	e := m.vectors[1].coordinates[1]
	f := m.vectors[2].coordinates[1]
	g := m.vectors[0].coordinates[2]
	h := m.vectors[1].coordinates[2]
	i := m.vectors[2].coordinates[2]
	return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i -  a*f*h
}

func (m Matrix2D) Det2D() complex128 {
	a := m.vectors[0].coordinates[0]
	b := m.vectors[1].coordinates[0]
	c := m.vectors[0].coordinates[1]
	d := m.vectors[1].coordinates[1]
	return a*d - b*c
}

func FDet2D(m Matrix2D) complex128 {
	a := m.vectors[0].coordinates[0]
	b := m.vectors[1].coordinates[0]
	c := m.vectors[0].coordinates[1]
	d := m.vectors[1].coordinates[1]
	return a*d - b*c
}

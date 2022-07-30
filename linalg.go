package linalg

/*version: 0.0.0*/

type Vector struct {
	coordinates []complex128
}

type Matrix struct {
	vectors []Vector
}

type Vector3D struct {
	coordinates [3]complex128
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

func Mul(m Matrix, v Vector) Vector {
	var newVector Vector
	for i := range v.coordinates {
		newVector.Add(m.vectors[i].ConstMul(v.coordinates[i]))
	}
	return newVector
}

func (m Matrix) Add(m2 Matrix) Matrix {
	for i := range m.vectors {
		for j := 0; j < len(m.vectors[i].coordinates); j += 1 {
			m.vectors[i].coordinates[j] += m2.vectors[i].coordinates[j]
		}
	}
	return m
}

package swiftVisitor

import (
	"fmt"
)

type Matrix3D struct {
	data  [][][]*SwiftValue
	rows  int
	cols  int
	depth int
	tipo  string
}

func NewMatrix3D(rows, cols, depth int, tipo string) *Matrix3D {
	data := make([][][]*SwiftValue, rows)
	for i := range data {
		data[i] = make([][]*SwiftValue, cols)
		for j := range data[i] {
			data[i][j] = make([]*SwiftValue, depth)
		}
	}
	return &Matrix3D{
		data:  data,
		rows:  rows,
		cols:  cols,
		depth: depth,
		tipo: tipo,
	}
}

func (m *Matrix3D) SetValue(row, col, depth int, value *SwiftValue) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols && depth >= 0 && depth < m.depth {
		m.data[row][col][depth] = value
	}
}

func (m *Matrix3D) GetValue(row, col, depth int) *SwiftValue {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols && depth >= 0 && depth < m.depth {
		fmt.Println("si encontro")
		return m.data[row][col][depth]
	}
	fmt.Println("no encontro")
	return nil
}
func (m *Matrix3D) SetDataFromList(dataList []*SwiftValue) {
	index := 0
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			for k := 0; k < m.depth; k++ {
				if index < len(dataList) {
					m.SetValue(i, j, k, dataList[index])
					index++
				} else {
					return
				}
			}
		}
	}
}

func (m *Matrix3D) Print() {
	for _, row := range m.data {
		for _, col := range row {
			fmt.Print("[ ")
			for _, val := range col {
				fmt.Printf("%d ", val)
			}
			fmt.Print("] ")
		}
		fmt.Println()
	}
}

func (m *Matrix3D) setearDatos(data *Matrix3D) {
	m=data
}

func (m *Matrix3D) GetSubMatrixAt(row, col int) []*SwiftValue{
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		return m.data[row][col]
	}
	fmt.Println("Fila o columna fuera de rango")
	return nil
}
package swiftVisitor

import (
	"fmt"
)

type Matrix struct {
	data [][]*SwiftValue
	rows int
	cols int
	tipo string
}

func NewMatrix(rows, cols int, tipo string) *Matrix {
	data := make([][]*SwiftValue, rows)
	for i := range data {
		data[i] = make([]*SwiftValue, cols)
	}
	
	return &Matrix{
		data: data,
		rows: rows,
		cols: cols,
		tipo: tipo,
	}
}

func (m *Matrix) SetValue(row, col int, value *SwiftValue) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		m.data[row][col] = value
	}
}

func (m *Matrix) GetValue(row, col int) *SwiftValue {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		return m.data[row][col]
	}
	return NULL
}

func (m *Matrix) Impresion() {
	for _, row := range m.data {
		fmt.Print("[")
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Print("]")
		fmt.Println()
	}
}

func (m *Matrix) SetValuesFromList(data []*SwiftValue) {
	index := 0
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if index < len(data) {
				m.SetValue(i, j, data[index])
				index++
			} else {
				return
			}
		}
	}
}

func (m *Matrix) setearDatos(data *Matrix) {
	m=data
}

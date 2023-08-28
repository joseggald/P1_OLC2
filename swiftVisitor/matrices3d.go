package swiftVisitor

import (
	"fmt"
)

type Matrix3D struct {
	data  [][][]*SwiftValue
	rows  int
	cols  int
	depth int
}

func NewMatrix3D(rows, cols, depth int) *Matrix3D {
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
	}
}

func (m *Matrix3D) SetValue(row, col, depth int, value *SwiftValue) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols && depth >= 0 && depth < m.depth {
		m.data[row][col][depth] = value
	}
}

func (m *Matrix3D) GetValue(row, col, depth int) *SwiftValue {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols && depth >= 0 && depth < m.depth {
		return m.data[row][col][depth]
	}
	return nil // You might want to handle out-of-bounds cases differently
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

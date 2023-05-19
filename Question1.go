package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//(a)
	matrix := Matrix{Rows: 3, Columns: 3, Elements: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}
	//(d)
	rows := matrix.GetRows()
	fmt.Println("Number of rows: ", rows)

	//(e)
	col := matrix.GetColumns()
	fmt.Println("Number of columns: ", col)

	for _, row := range matrix.Elements {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//(f)
	matrix.SetElements(1, 3, 10)
	for _, row := range matrix.Elements {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
	fmt.Println()
	
	//(g)
	matrix2 := Matrix{
		Rows:    3,
		Columns: 3,
		Elements: [][]int{
			{11, 12, 13},
			{14, 15, 16},
			{17, 18, 19},
		},
	}
	
	result := matrix.AddMatrix(matrix2)
	for _, row := range result.Elements {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//(h)
	ConvertJson(matrix2)
	ConvertJson(result)
}

type Matrix struct {
	Rows     int
	Columns  int
	Elements [][]int
}

func (m Matrix) GetRows() int{
	return m.Rows
}

func (m Matrix) GetColumns() int{
	return m.Columns
}

func (m *Matrix) SetElements(i, j, value int) {
	if i >= 1 && i <= m.Rows && j >= 1 && j <= m.Columns {
		m.Elements[i-1][j-1] = value
	} else {
		fmt.Println("Position is invalid")
	}
}

func (m Matrix) AddMatrix(matrix2 Matrix) Matrix{
	result := m 
	for i := 0; i<m.Rows; i++{
		for j := 0; j<m.Columns; j++{
			result.Elements[i][j]+=matrix2.Elements[i][j]
		}
	}
	return result
}

func ConvertJson(m Matrix) {
	finalJson, err := json.Marshal(m.Elements)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
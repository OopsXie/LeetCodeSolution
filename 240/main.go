package main

import (
	"fmt"
	"time"
)

func findRowsCols(matrix [][]int, target int, rows int, cols int, x int, y int) bool {
	if rows == 1 {
		low := y
		high := y + cols - 1
		for low <= high {
			if matrix[x][(low+high)/2] == target {
				return true
			}
			if matrix[x][(low+high)/2] < target {
				low = (low+high)/2 + 1
			}
			if matrix[x][(low+high)/2] > target {
				high = (low+high)/2 - 1
			}
		}
		return false
	}
	if cols == 1 {
		low := x
		high := x + rows - 1

		for low <= high {
			if matrix[(low+high)/2][y] == target {
				return true
			}
			if matrix[(low+high)/2][y] < target {
				low = (low+high)/2 + 1
			}
			if matrix[(low+high)/2][y] > target {
				high = (low+high)/2 - 1
			}
		}
		return false
	}

	return false
}

func searchSubMatrix(matrix [][]int, target int, rows int, cols int, x int, y int) bool {
	if rows <= 0 || cols <= 0 {
		return false
	}

	if rows == 1 || cols == 1 {
		return findRowsCols(matrix, target, rows, cols, x, y)
	}

	low_rows := x
	low_cols := y

	high_rows := x + rows - 1
	high_cols := y + cols - 1

	for low_rows <= high_rows && low_cols <= high_cols && low_rows >= 0 && low_cols >= 0 && high_rows >= 0 && high_cols >= 0 {
		middle_rows := (low_rows + high_rows) / 2
		middle_cols := (low_cols + high_cols) / 2

		if matrix[middle_rows][middle_cols] == target || matrix[low_rows][low_cols] == target || matrix[high_rows][high_cols] == target {
			return true
		}
		if matrix[middle_rows][middle_cols] < target {
			low_rows = middle_rows + 1
			low_cols = middle_cols + 1
		}
		if matrix[middle_rows][middle_cols] > target {
			high_rows = middle_rows - 1
			high_cols = middle_cols - 1
		}
	}

	right_up_rows := low_rows - x
	right_up_cols := y + cols - low_cols
	right_up_start_x := x
	right_up_start_y := low_cols

	left_down_rows := x + rows - low_rows
	left_down_cols := low_cols - y
	left_down_start_x := low_rows
	left_down_start_y := y

	return searchSubMatrix(matrix, target, right_up_rows, right_up_cols, right_up_start_x, right_up_start_y) ||
		searchSubMatrix(matrix, target, left_down_rows, left_down_cols, left_down_start_x, left_down_start_y)
}

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	return searchSubMatrix(matrix, target, rows, cols, 0, 0)
}

func main() {
	start := time.Now()
	matrix := [][]int{
		{1, 3, 5, 7, 9},
		{2, 4, 6, 8, 10},
		{11, 13, 15, 17, 19},
		{12, 14, 16, 18, 20},
		{21, 22, 23, 24, 25},
	}

	target := 10

	isTure := (searchMatrix(matrix, target))
	elapsed := time.Since(start)
	fmt.Println(isTure)
	fmt.Println("耗时（微秒）：", elapsed.Microseconds(), "µs")

}

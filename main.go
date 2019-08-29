package main

import (
	"fmt"
	"strings"
	"math"
	"strconv"
)

func sudokuQuadrantChecker(strArr []string) string {
	ca := createSudokuCoordinateArray(strArr)
	duplicates := make(map[int]bool, 9)

	// check each horizontal row
	for i:=0; i<9; i++ {
		dupes := getDuplicatePositions(ca[i]);
		for j := range dupes {
			duplicates[(calculateQuadrent(i) - 1)*3 + j] = true
		}
	}

	// check each vertical row
	for i:=0; i<9; i++ {
		row := []string{ca[0][i], ca[1][i], ca[2][i], ca[3][i], ca[4][i], ca[5][i], ca[6][i], ca[7][i], ca[8][i]}
		dupes := getDuplicatePositions(row);
		for j := range dupes {
			fmt.Println("j", j, i)
			k := (j - 1) * 3
			duplicates[calculateQuadrent(i) + k] = true
		}
	}

	// check each quadrant
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			a := i * 3
			b := j * 3
			row := []string{ca[a][b], ca[a+1][b], ca[a+2][b], ca[a][b+1], ca[a+1][b+1], ca[a+2][b+1], ca[a][b+2], ca[a+1][b+2], ca[a+2][b+2]}
			dupes := getDuplicatePositions(row);
			if len(dupes) > 0 {
				k := (i * 2) + 1
				duplicates[i + j + k] = true
			}
		}
	}

	var result []string
	for k := range duplicates {
		result = append(result, strconv.Itoa(k))
	}

	return strings.Join(result, ",")
}

func createSudokuCoordinateArray(strArr []string) [][]string {
	s := make([][]string, 0, 9)
	for _, str := range strArr {
		s = append(s, strings.Split(str[1:len(str)-1], ","))
	}
	return s
}

func getDuplicatePositions(arr []string) map[int]bool {

	seen := make(map[string]int, len(arr))
	dupes := make(map[int]bool, 3)

	for i, s := range arr {
		if j, ok := seen[s]; ok {
			if (s != "x") {
				dupes[calculateQuadrent(i)] = true
				dupes[calculateQuadrent(j)] = true
			}
			continue
		}

		seen[s] = i
	}
	return dupes
}

func calculateQuadrent(i int) int {
	return int(math.Ceil(float64(i/3))) + 1
}

func main() {

	grid := []string{"(1,2,3,4,5,6,7,8,1)","(x,x,x,x,x,x,x,x,x)","(x,x,x,x,x,x,x,x,x)","(1,x,x,x,x,x,x,x,x)","(x,x,x,x,x,x,x,x,x)","(x,x,x,x,x,x,x,x,x)","(x,x,x,x,x,x,x,x,x)","(x,x,x,x,x,x,x,x,x)","(x,x,x,x,x,x,x,x,x)"};

	// returns list of quadrants (a 3x3 grid) with errors in
	result := sudokuQuadrantChecker(grid)

	fmt.Println("results:", result)
}

package main

import "math/rand"

func generateGameID() int64 {
	return (rand.Int63n(8)+1)*1e17 + rand.Int63n(1e17)
}

func checkFor(mat [6][7]int, x int) int {
	nrOcupate := 0
	dx := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if mat[i][j] != 0 {
				nrOcupate++
			}
			if mat[i][j] != x {
				continue
			}
			isValid := func(x, y int) bool {
				return (x >= 0 && x < 6 && y >= 0 && y < 7)
			}
			for d := 0; d < len(dx); d++ {
				var ok bool = true
				val := mat[i][j]
				for k := 1; k <= 3; k++ {

					nx := i + k*dx[d]
					ny := j + k*dy[d]
					if !isValid(nx, ny) {
						ok = false
						break
					}
					if val != mat[nx][ny] {
						ok = false
					}
				}
				if ok == true {
					return 1
				}
			}
		}
	}
	if nrOcupate == 42 {
		return -1
	}
	return 0
}

func resultFromGame(game Game) int {
	return 1*checkFor(game.Table, 1) + 2*checkFor(game.Table, 2)
}

func matrixToArray(mat [6][7]int) []int {
	nr := []int{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			nr = append(nr, mat[i][j])
		}
	}
	return nr
}

func arrayToMatrix(nrx []int64) [6][7]int {
	n := 0
	nr := [6][7]int{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			nr[i][j] = int(nrx[n])
			n++
		}
	}
	return nr
}

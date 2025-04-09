package main

import (
	"fmt"
	"image"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func FindCoordAround(nbDimentions int, nbAround int) *[][]int {
	find := make([][]int, 0)
	for i := 1; i <= nbDimentions; i++ {
		GetNCoordDifferent(nbDimentions, i, nbAround, &find)
	}
	return &find
}

func GetNCoordDifferent(nbDimentions int, nbDifferent int, nbAround int, find *[][]int) {
	coord := make([]int, nbDifferent)
	for i := 0; i < nbDifferent; i++ {
		coord[i] = -nbAround + i
	}
	GetCompletVariationOfCoord(nbDimentions, find, &coord)

	for ;;{
		modified := false
		cursor := len(coord)-1
		for ;!modified; {
			if cursor < 0 || cursor == len(coord) {
				return
			}
			if coord[cursor] + len(coord) -1 - cursor < nbAround {
				modified = true
				coord[cursor]++
				val := coord[cursor]+1
				cursor++
				for ;cursor < len(coord); {
					coord[cursor] = val
					cursor++
					val++
				}
			} else {
				cursor--
			}
		}
		GetCompletVariationOfCoord(nbDimentions, find, &coord)
	}
}

func GetCompletVariationOfCoord(nbDimentions int, find *[][]int, partialCoord *[]int) {
	coord := make([]int, nbDimentions)
	for i := 0; i < nbDimentions; i++ {
		if i < len(*partialCoord) {
			coord[i] = (*partialCoord)[i]
		} else {
			coord[i] = (*partialCoord)[0]
		}
	}
	getPermutationOfCoord(find, &coord)
	for ;;{
		modified := false
		cursor := len(coord)-1
		for ;!modified; {
			if cursor == len(*partialCoord) -1{
				return
			}
			if coord[cursor] != (*partialCoord)[len(*partialCoord)-1] {
				modified = true
				afterCursor := 0
				for ;afterCursor < len(*partialCoord) && coord[cursor] != (*partialCoord)[afterCursor]; {
					afterCursor++
				}
				afterCursor++
				for ;cursor < len(coord); {
					coord[cursor] = (*partialCoord)[afterCursor]
					cursor ++
				}
			} else {
				cursor--
			}
		}
		getPermutationOfCoord(find, &coord)
	}
}

func getPermutationOfCoord(find *[][]int, coord *[]int) {
	//https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
	var helper func([]int, int)
	tempFind := make([][]int, 0)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			if !isDoublon(&tempFind, &tmp) {
				tempFind = append(tempFind, tmp)
			}
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	tcoord := make([]int, len(*coord))
	copy(tcoord, *coord)
	helper(tcoord, len(tcoord))
	*find = append(*find, tempFind...)
}

func isDoublon(find *[][]int, coord *[]int) bool {
	for i := 0; i < len(*find); i++ {
		for j := 0; j < len(*coord); j++ {
			if (*find)[i][j] != (*coord)[j] {
				break
			}
			if j == len(*coord)-1 {
				return true
			}
		}
	}
	return false
}

func getCursorPosition(win *pixelgl.Window) (int,int) {
	return int(win.MousePosition().X), 280- int(win.MousePosition().Y)
}

func drawPixel(img *image.RGBA, x int, y int) {
	fmt.Println("début de la fonction drawPixel")
	drawAtDist := func(dist int, x int, y int, color float64) {
		for i := 0; i < 2*dist+1; i++ {
			img.Set(x-dist+i, y-dist, pixel.RGB(color, color, color))
			img.Set(x-dist, y-dist+i, pixel.RGB(color, color, color))
			img.Set(x+dist-i, y+dist, pixel.RGB(color, color, color))
			img.Set(x+dist, y+dist-i, pixel.RGB(color, color, color))
		}
	}

	for i := 1; i < 10; i++ {
		c := 255*float64(i)/9
		fmt.Println(c,9-i)
		drawAtDist(9-i, x, y, c)
	}
}
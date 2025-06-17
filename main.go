package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const Size = 6

var counter uint64 = 0

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(Size * Size)
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			go possibilities(i, j, wg)
		}
	}
	wg.Wait()
	fmt.Println(atomic.LoadUint64(&counter))
}

func possibilities(row int, col int, wg *sync.WaitGroup) {
	defer wg.Done()
	var field = [Size][Size]bool{}
	atomic.AddUint64(&counter, possibilitiesInternal(field, row, col, 1))
}

func possibilitiesInternal(field [Size][Size]bool, row int, col int, count uint8) uint64 {
	if row < 0 || row >= Size || col < 0 || col >= Size || field[row][col] {
		return 0
	} else if count == Size*Size {
		return 1
	} else {
		field[row][col] = true
		return possibilitiesInternal(field, row, col-1, count+1) +
			possibilitiesInternal(field, row, col+1, count+1) +
			possibilitiesInternal(field, row+1, col, count+1) +
			possibilitiesInternal(field, row-1, col, count+1)
	}
}

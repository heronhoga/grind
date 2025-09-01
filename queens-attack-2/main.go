package main

import (
	"fmt"
	"math"
)

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	up := n - r_q
    down := r_q - 1
    right := n - c_q
    left := c_q - 1
    upRight := min(up, right)
    upLeft := min(up, left)
    downRight := min(down, right)
    downLeft := min(down, left)

	for _, obstacle := range obstacles {
        r, c := obstacle[0], obstacle[1]

        // same column
        if c == c_q {
            if r > r_q {
                up = min(up, r-r_q-1)
            } else {
                down = min(r_q-r-1, down)
            }
        }

        // same row
        if r == r_q {
            if c > c_q {
                right = min(right, c-c_q-1)
            } else {
                left = min(left, c_q-1-c)
            }
        }

        //hexagonal
        if math.Abs(float64(r-r_q)) == math.Abs(float64(c-c_q)) {
            //up right
            if r > r_q && c > c_q {
                upRight = min(upRight, r-r_q-1)
            //down right
            } else if r < r_q && c > c_q {
                downRight = min(downRight, r_q-r-1)
            //up left
            } else if r > r_q && c < c_q {
                upLeft = min(upLeft, r-r_q-1)
            //down left
            } else if r < r_q && c < c_q {
                downLeft = min(downLeft, r_q-r-1)
            }
        }
    }

    return up + down + right + left + upRight + downRight + upLeft + downLeft
}

func main() {
    n := int32(5)
    k := int32(3)
    r_q := int32(4)
    c_q := int32(3)

    obstacles := [][]int32{
        {5, 5},
        {4, 2},
        {2, 3},
    }

    result := queensAttack(n, k, r_q, c_q, obstacles)
    fmt.Println("hasil: ", result)
}
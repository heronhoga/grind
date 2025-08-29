package main

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
    minBlack := int64(b) * min(int64(bc), int64(wc)+int64(z))
    minWhite := int64(w) * min(int64(wc), int64(bc)+int64(z))

    return minBlack + minWhite
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}

func main() {
 taumBday(3, 5, 3, 4, 1)
}
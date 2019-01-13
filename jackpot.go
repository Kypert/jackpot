package main

import "time"
import "fmt"
import "math/rand"

func main() {
	src1 := rand.NewSource(time.Now().UnixNano())
	rdm1 := rand.New(src1)
	fmt.Print(rdm1.Intn(50) + 1, ", ")
	fmt.Print(rdm1.Intn(50) + 1, ", ")
	fmt.Print(rdm1.Intn(50) + 1, ", ")
	fmt.Print(rdm1.Intn(50) + 1, ", ")
	fmt.Print(rdm1.Intn(50) + 1, ", ")

	time.Sleep(3000 * time.Millisecond)

	src2 := rand.NewSource(time.Now().UnixNano())
	rdm2 := rand.New(src2)
	fmt.Print(rdm2.Intn(10) + 1, ", ")
	fmt.Print(rdm2.Intn(10) + 1)
	fmt.Println()
}

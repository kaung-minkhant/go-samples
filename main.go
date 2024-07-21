package main

import "fmt"

func main() {
	// var name string = "abc"
	// var nameTwo = "abcd"
	// var nameThree string
	// nameFour := "kmk"
	//
	// fmt.Println(name, nameTwo, nameThree, nameFour)
	// fmt.Println("Hello")
	//
	// fmt.Printf("my age is %v my name is %v. \n", 24, "shunn")
	// fmt.Printf("my age is %q my name is %q.\n", "24", "shunn")
	// fmt.Printf("Age is of type %T\n", 25)
	// fmt.Printf("Your score is %0.2f\n", 1.255)
	//
	// // save printf
	// printString := fmt.Sprintf("Your score is %0.2f\n", 1.255)
	// fmt.Println("Saved string is: ", printString)
	// var ages [3]int = [3]int{1, 2, 3}
	// ages[1] = 4
	// fmt.Println(ages, len(ages))
	// names := [2]string{"kmk", "shunn"}
	// fmt.Println(names, len(names))
	//
	// // slices
	// var scores = []int{1, 2, 3}
	// scores[2] = 5
	// scores = append(scores, 4)
	// sort.Ints(scores)
	// fmt.Println(scores, len(scores))
	//
	// // slice ranges
	// rangeOne := names[0:]
	// fmt.Println(rangeOne)
	//
	// var name string = "shunn le yee"
	// fmt.Println(strings.Split(name, " "))
	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Printf("I'm a bool %v \n", t)
	// 	case int:
	// 		fmt.Printf("I'm a int with value %v with type %T \n", t, t)
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }
	// whatAmI(5)
	// mapOne := make(map[int]string)
	// mapOne := new(map[int]string)
	// (*mapOne)[1] = "polar"
	// mapOne[1] = "polar"
	// mapOneValue := mapOne[1]
	// mapTwoValue, ok := mapOne[2]
	// fmt.Println("Sample Map is ", mapOne)
	// fmt.Println("Sample Map One Value is ", mapOneValue)
	// fmt.Println("Sample Map One Value is ", mapTwoValue, "and it is ok?", ok)
	// append := func(slice, data []byte) []byte {
	// 	// l := len(slice)
	// 	// if l+len(data) > cap(slice) {
	// 	// 	newSlice := make([]byte, l+len(data)*2)
	// 	// 	copy(newSlice, slice)
	// 	// 	slice = newSlice
	// 	// }
	// 	// slice = slice[0 : l+len(data)]
	// 	// fmt.Printf("New capacity is %v\n", cap(slice))
	// 	// fmt.Printf("New length is %v\n", len(slice))
	// 	// copy(slice[l:], data)
	// 	slice[1] = 100
	// 	return slice
	// }
	// slice1 := make([]byte, 10)
	// slice2 := make([]byte, 5)
	// for i := range len(slice1) {
	// 	slice1[i] = byte(i)
	// }
	// for i := range len(slice2) {
	// 	slice2[i] = byte(i)
	// }
	// slice3 := append(slice1, slice2)
	// append(slice1, slice1)
	// fmt.Println("Slice1", slice1)
	// fmt.Print(slice1, "\n")
	// fmt.Println("Slice3", slice3)
	// fmt.Println("Slice1 == Slice3", slices.Equal(slice1, slice3))
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println("Fib of 9 is", fib(9))
}

//use go run . to execute program

package main

//if lines are commented out, the imported package (such as strings) will be removed
import (
	"fmt"
	"sort"
)

func main() {

	// greeting := "Hello there friends"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// //replaces Hello with Hi
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))

	//searches for the index of a letter in the greeting
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println((strings.Split(greeting, " ")))

	//the original value is unchanged
	// fmt.Println("original string value=", greeting)

	ages := []int{45, 55, 65, 75, 20, 40, 19}
	sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 75)
	fmt.Println(index)

	names := []string{"Annika", "Elliot", "Mario", "Zelda"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Mario"))

}

//strings
// fmt.Println("Hello, Annika!")

// var nameOne string = "mario"
// var nameTwo = "luigi"
// var nameThree string

// fmt.Println(nameOne, nameTwo, nameThree)

// nameOne = "peach"
// nameThree = "elliot"
// nameFour :="yoshi"

// fmt.Println(nameFour)

// //ints

// var ageOne int = 20
// var ageTwo = 30
// ageThree := 40

// fmt.Println(ageOne, ageTwo, ageThree)

// //bits & memory

// // var numOne int8 = 25
// // var numTwo int8= -128
// // var numThree uint16 = 256

// var scoreOne float32 = 25.98
// var scoreTwo float64 = 9423456654.4
// scoreThree := 1.5

// age := "39"
// name := "annika"

//Print
// fmt.Print("hello, ")
// fmt.Print("world! \n")
// fmt.Print("new line \n")

//Println
// fmt.Println("hello annika")
// fmt.Println("goodbye annika")
// fmt.Println("my age is", age, "and my name is", name)

//Printf (formatted strings)
// //%v searches for the first value by default when added to line - %v format specifier
// fmt.Printf("my age is %v and my name is %v \n", age, name)
// //%q puts double quotes around a string, does not work for age as it's an integer
// fmt.Printf("my age is %q and my name is %q \n", age, name)
// fmt.Printf("age is of type %T", age)
// fmt.Printf("you scored %f points \n", 255.55)
// //output for line 29 is "you scored 255.6 points"
// fmt.Printf("you scored %0.1f points \n", 255.55)

// //Sprintf (save formatted strings)
// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
// fmt.Println("the saved string is:", str)

// var ages [3] int = [3] int{20, 25, 30}
// var ages = [3]int{20, 25, 30}

// names := [4]string{"Annika", "Elliot", "Mario", "Zelda"}
// names[1] = "Luigi"
// //elements in arrays are fixed
// fmt.Println(ages, len(ages))
// fmt.Println(names, len(names))

// //slices(use arrays under the hood)
// var scores = []int{100, 50, 60}
// scores[2] = 33
// scores = append(scores, 85)

// fmt.Println(scores, len(scores))

// //sclice range prints everything from 1st element (index 0) to 3rd element
// rangeOne := names[1:3]
// rangeTwo := names[2:]
// rangeThree := names[:3]
// fmt.Println(rangeOne, rangeTwo, rangeThree)

// rangeOne = append(rangeOne, "Koopa")
// fmt.Println(rangeOne)

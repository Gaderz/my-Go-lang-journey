package main

import (
	"fmt"
	"math"
)

//  func main(){
// 	fmt.Println("hello world")
//  }
//  //PRIMITIVE TYPES

//  //BOOLEAN

//   func main(){
// 		var n bool = true
// 		fmt.Printf("%v, %T\n", n, n)
// 	}
//how to identify boolean using logic
//  func main(){
// 	n := 1==1
// 	m := 1==2
// 	fmt.Printf("%v, %T", n, n)
// 	fmt.Printf("%v, %T", m, m)

// }
//  everytime you initialize a variable it has a 0 value which is false boolean

// #NUMERIC type

// # INTEGER type
//  REGARDLESS of your environemnt an int will be atleast 64bit or 32bit or 132 bit depending on the system you are working
// signed and unsigned integers
// 	 func main(){
// a :=10
// b :=5
// fmt.Println(a+b)
// fmt.Println(a-b)
// fmt.Println(a*b)
// fmt.Println(a/b)
// fmt.Println(a%b)

// }
// bit operators

// 				func main(){
// a :=10
// b :=3
// fmt.Println(a & b) //0010
// fmt.Println(a | b)//1011
// fmt.Println(a ^ b)//1001
// fmt.Println(a &^ b) //0100
// }

// bit shifting

// func main(){
// a :=8

// fmt.Println(a << 3) //2^3 * 2^3 = 2^6 ==64
// fmt.Println(a >> 3)///2^3 /2^3  = 2^0 == 1

// }

//FLOATING POINT NUMBERS

// func main() {
// 	n := 3.14
// 	n = 13.7e72
// 	n = 2.1E14
// 	fmt.Printf("%v, %T", n, n)
// }

// //  COMPLEX DATA TYPE

// func main() {
// 	var n complex64 = 1 + 2i

// 	fmt.Printf("%v, %T", n, n)
// }
// func main() {
// 	a := 1 +2i
// 	b := 2 + 5.2i

// 	fmt.Println(a+b)
// 		fmt.Println(a-b)
// 			fmt.Println(a*b)
// 				fmt.Println(a/b)

// // }
//  func main() {
// 	var n complex128 = 1 + 2i

// 	fmt.Printf("%v, %T\n", real(n), real(n))
// 	fmt.Printf("%v, %T\n", imag(n), imag(n))
// //  }
//  func main() {
// 	var n complex128 = complex(5, 12)

// 	fmt.Printf("%v, %T\n",n, n)
//  }

//  TEXT DATA type
//  2 CATEGORIES
//  #1 STRINGS
// THIS REPERESENT ANT UTF 32 CHARACTERS

// func main(){
// 	s := "this is a string"
// s2 := "this  is also a string"
//concatinate 2 strings
// fmt.Printf("%v, %T\n", s + s2, s+ s2)

// slice of byte oe collection of byte
// 	b := []byte(s)
// 	fmt.Printf("%v, %T\n", b,b)
// }

// RUNE DATA TYPE
// THIS REPERESENT ANT UTF 32 CHARACTERS
// func main(){
// 	r := 'a'
// 	fmt.Printf("%v, %T\n", r,r)
// }

//  CONSTANT
// const (
// a =iota
// b = iota
// c=iota
// )
// func main(){
// 	const d  = 42
// 	var b int16 = 33
// 	// adding variable and const together
// 	fmt.Printf("%V\n", a)
// }

// VALUE UNCREASES AS THE CODE KEEO EXECUTING
// const (
// a =iota
// b = iota
// c=iota
// )
// const (
// 	catSpecialist = iota
// 	dogSpecialist
// 	snakeSpecialist
// )
// func main(){
// 	// fmt.Printf("%v\n", a)
// 	// fmt.Printf("%v\n", b)
// 	// fmt.Printf("%v\n", c)
// 	 var specialitType int = catSpecialist
// 	 fmt.Printf("%v\n", specialitType ==  catSpecialist)
// }
//  const (
// 	_ = iota
// 	kP = 1 <<(10 * iota)
// 	MP
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
//  )

//  func main() {
// 	fileSize := 400000000000.
// 	fmt.Printf("%.2fGB", fileSize/GB)
//  }
// IOTA IS SCOPED TO A CONSTANT BLOCK

// const (
// 	isAdmin = 1 << iota
// 	isHeadQuaters
// 	canSeeFinance
// 	cabSeeAfrica
// 	canSeeAsia
// 	canSeeEuropa
// 	canSEENoethAmerica
// 	canSeeSouthAmerica
// )
// //  constant are immutable but cant be shadowed
// func main(){
// 	var roles byte = isAdmin | canSeeFinance | canSeeEuropa
// 	fmt.Printf("%b\n", roles)
// 	fmt.Printf("is Admin? %v\n", isAdmin & roles == isAdmin)
// 	fmt.Printf("is HQ? %v\n", isHeadQuaters & roles == isHeadQuaters)
// }

// COLLECTION TYPES
// ARRAYS AND SLICES

//  func 	main(){
// 	// grades := [3]int{36,87,98}
// 	// fmt.Printf("grades: %v", grades)
// 	var students  [3]string
// 	fmt.Printf("students: %v\n", students)
// 	students[0] = "lisa"
// 	students[1] = "Nelson"
// 	students[2] = "Elviz"
// 	fmt.Printf("students: %v\n", students)
// 	fmt.Printf("number of students: %v\n", len(students))
// 	var numberMtrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,1}, [3]int{1,0,1}}
// 	fmt.Printf("numberMatrix: %v", numberMtrix)
//  }

// ARRAY SLICE METHOD

// func main(){
// 	a := []int{1,2,3,4,5,6,7,8,9,10}
// 	b :=a[:] //SLICE ALL ELEMENT
// 	c :=a[3:] //SLICE FROM 4TH ELEMENT
// 	d :=a[:6]//SLICE FIRST 6 ELEMTS
// 	e :=a[3:6] //CLICE 4HT, 5TH AND 6TH ELEMENT
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// }

// MAKE FUNCTIONS
// func main() {
// // x := make([]int, 3)
// x := []int{}
// fmt.Println(x)
// fmt.Printf("lenght: %v\n", len(x))
// fmt.Printf("capacityt: %v\n", cap(x))
// x = append(x, 1)
// fmt.Println(x)
// fmt.Printf("lenght: %v\n", len(x))
// fmt.Printf("capacityt: %v\n", cap(x))
// // SPREAD OPERATORA
// x = append(x, []int{2, 4, 5, 6, 8}...)
// fmt.Println(x)
// fmt.Printf("lenght: %v\n", len(x))
// fmt.Printf("capacityt: %v\n", cap(x))

//STACK OPPERATION
// shifting >>>>> removing element at the miffle
// 	a := []int{3, 5, 6, 7, 8}
// 	b := append(a[:2], a[3:]...)
// 	fmt.Println(b)
// 	fmt.Println(a)
// 	// SUMMARY

// }

//  SUMMARY
//  ARRAYS
// 	COLLEXTION OF ITEMS WITH SAME TYPE
// 	fixe size
// 	declaration of array
// 	eg ---a := [3]int{2,4,5}
// 				b := [...]int{1,4,5}
// 				var a [3]int
// 				access to zero-base index
// 				a := [3]int{1,3,2} ...a[1] ==3
// 				len function returns rhe size of the array
// 				copies refer to df underlying data

// 				SLICES

// 				backed by array
// 				creation styles
// slices existing arrays or slice
// literal style
// via make function
// eg ....a := make([]int, 10) create slice with capacity and lenght
//        b := make([]int,10, 100) slice with lenght==10 and capacity ==100
// 			 len function returns the lenght of slice
// 			 cap funtion return the capacity of slice
// 			 append function to add element to slice
// 			 make,, causes expensive copy operations if underlying arrays are too small
// 			 copy refes to underlying arrays

// MAOS AND STRUCTS COLLECTIONS TYPES
// func main() {
//literal synthax
// statePopulation := map[string]int{
// 	"Enugu":  3687846,
// 	"Abujs":  7478775,
// 	"Awka":   475587,
// 	"Kaduna": 9058934774,
// 	"Kano":   84745468,
// }
// fmt.Println(statePopulation)
// WITH MAKE FUNCTION
// 	statePopulation := make(map[string]int)
// 	statePopulation = map[string]int{
// 		"Enugu":  3687846,
// 		"Abujs":  7478775,
// 		"Awka":   475587,
// 		"Kaduna": 9058934774,
// 		"Kano":   84745468,
// 	}
// 	fmt.Println(statePopulation)
// }

// STRUCT DATA TYPE

// type doctor struct {
// 	number    int
// 	actorNmae string
// 	companion []string
// }

// func main() {
// aDoctor := doctor{
// 	number:    3,
// 	actorNmae: "John Legend",
// 	companion: []string{
// 		"lizzy Benson", "Nelson Gaderz", "chika ugo",
// 	},
// }
// fmt.Println(aDoctor.actorNmae, aDoctor.companion[2], aDoctor.number)

// passing data in struct
// 	aMan := struct{ name string }{name: "Nleson William"}
// 	anotherMan := aMan
// 	anotherMan.name = "John Dili"
// 	fmt.Println(aMan)
// 	fmt.Println(anotherMan)

// // }

// type Animal struct {
// 	Name   string
// 	Origin string
// }
// type Bird struct {
// 	Animal
// 	Speed  float32
// 	CanFly bool
// }

// func main() {
// 	b := Bird{}
// 	b.Name = "Parrot"
// 	b.Origin = "Austrlia"
// 	b.Speed = 48
// 	b.CanFly = false
// 	fmt.Println(b)
// }

// SUMMARY
// MAPS ANS=D STRUCT
// MAP:
// collection of values types that accessed via keys
// created via literal or via make function
// members accessed via [key] synthax
// .myMap["key"] = "value"
// check for presence with "value, ok" form of result
// mltiple assignemnt refers to a,e undeerlying data.

// // STRUCT
//  collection of disparate data types that decsribe single concept
//  keyed by names field
//  normally created as types, but anonymous struct are allowed
//  structs are value type
//  no inheritance , but use composition via embending
//  tags can be added to struct fields to dcribes field

//  NEXT IS CONTORL FLOW
// if and else statement?

func main() {
	// literal synthax
	statePopulation := map[string]int{
		"Enugu":  3687846,
		"Abujs":  7478775,
		"Awka":   475587,
		"Kaduna": 9058934774,
		"Kano":   84745468,
	}
	if pop, ok := statePopulation["Kaduna"]; ok {
		fmt.Println(pop)
	}

	// guessing game
	number := 50
	guess := -5
	// USING || AND && OPERATORS
	// if guess < 1 || guess > 100 {
	// 	println("guess must be between 1 and 100")
	// }
	// if guess >= 1 && guess <= 100 {
	// 	if guess < number {
	// 		println("too low")
	// 	}
	// 	if guess > number {
	// 		println("too high")
	// 	}
	// 	if guess == number {
	// 		println("you got it right")
	// 	}
	// 	if guess != number {
	// 		println("not a number")
	// 	}

	// else statement
	if guess < 1 {
		println("guess must be between 1 and 100")
	} else if guess > 100 {
		println("guess must be less than 100")
	} else {
		if guess < number {
			println("too low")
		}
		if guess > number {
			println("too high")
		}
		if guess == number {
			println("you got it right")
		}

	}

	fmt.Println(number <= guess, number >= guess, number != guess)

	// ??working with numberS
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		println("they are thesame number")

	} else {
		println(myNum, "is a different digit from")
	}

	//  SWITCH STATEMENT
	switch i := 2 + 4; i {
	case 1, 5, 10:
		fmt.Println("one , five, ten")

	case 2, 4, 6:
		fmt.Println("two, four  or six")
	default:
		fmt.Println("not case 1 or case 2")
	}
	// TAGLIST SYNTHAX

	i := 21
	switch {
	case i < 10:
		fmt.Println(i, "less than or 10")
	case i == 10:
		fmt.Println(i, " is equal to 10")
	case i <= 20:
		fmt.Println(i, "is less than 20")
	default:
		fmt.Println(i, "is greater than 20")
	}

}

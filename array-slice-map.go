package main

// keep the memory you use stay loaded within CPU caches longer
// Random access
// Type inforamation for the array provides distance in memory you have to move to find each element
// Once an array is declared, neither the type of data being stored nor its length can b changed

// decare an integer array, set to its zero value
// Passing arrays between functions
// BY VALUE!! expensive!!
var array [1e6]int

func foo(array [1e6]int) {

}

func foo2(array *[1e6]int) {

}

type Vector struct {
	addr     *int
	length   int
	capacity int
}

func slice() {
	// a slice of strings, length and capacity are both 5
	slice := make([]string, 5)

	// length is 3, capacity is 5
	slice2 := make([]int, 3, 5)

	// nothing inside [], creating a slice of length/capacity as 5
	slice3 := []string{"Red", "Blue", "Green", "Yellow"}

	// a slice of length/capacity as 3
	slice4 := []int{10, 20, 30}

	// a slice of strings
	// init the 100th element as empty string
	slice := []string{99: ""}
}

func main() {
	// copy the whole array!!
	foo(array)

	// copy only the pointer, 8 bytes under x64 platform
	foo2(&array)

	array := [5]int{10, 20, 30, 40, 50}
	array2 := [...]int{10, 20, 30, 40, 50}

	// the rest of elementsis zero
	array3 := [5]int{1: 10, 2: 20}

	array[0] = -10

	// an integer pointer array
	array4 := [5]*int{0: new(int), 1: new(int)}
	// assign integer pointers to array element
	*array[0] = 10
	*array[1] = 20

	var array5 [5]string
	array6 := [5]string{"Red", "Green", "Blue"}
	// copy the values from array to array1
	array5 = array6
}

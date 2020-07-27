/*
Topics Covered
- Variable Declaration
- Conditional Statements
- For-Loops
- Array
- Functions
	- Anonymous Functions
- Classes
*/


// function without argument
fun getValueComaprison1(): String {
	var count: Int = 100
	var answerString: String = if (count < 50) {
		"Number is less"
	} else if (count > 50) {
		"Number is greater"
	} else {
		"Number is equal"
	}
	return answerString
}

// function with argument
fun getValueComaprison2(count: Int): String {
	var answerString: String = if (count < 50) {
		"Number is less"
	} else if (count > 50) {
		"Number is greater"
	} else {
		"Number is equal"
	}
	return answerString
}

// Another Declaration format of functions
fun getValueComaprison3(count: Int): String = if (count < 50) {
	"Number is less"
} else if (count > 50) {
	"Number is greater"
} else {
	"Number is equal"
}

// Anonymous Functions
val stringLengthFunc: (String) -> Int = { input -> input.length}


// Classes in Kotlin
class Person (Fname: String = "Shivam" , Lname: String = "Pal", var numOne: Int = 12, private var numTwo: Int = 13) {
	var first_name: String = Fname
	var last_name: String = Lname
	var full_name: String = "${first_name} ${last_name}"		// string template

	// init block of Kotlin class, called when this class will initiated
	init {
		println(full_name)
	}

	fun addTwoNumbers() : Int {
		return numOne + numTwo
	}
}


// Kotlin Main function
fun main() {

    println("Hello, World!") // Simple print command

	// 				1. Variable Declaration

	// Variable Declaration
	// var - value can be reassign and changed later
	// val - value can't be reassign

	var count1: Int = 10 // Byte, Long, Short, Float, Double
	println("Number of Count is ${count1}")	// Printing with variable

	val languageName: String = "Kotlin"
	println(languageName)

	// Kotlin variables can't hold null values by default
	// var name: String = null // Statement will throw error
	var name: String? = null // Statement is correct
	println(name)


	// 				2. Conditional Statements

	// Conditionals Statement in Kotlin

	// if-else expression for Conditionals
	var count: Int = 50
	if (count < 50) {
		println("Number is less")
	} else if (count > 50) {
		println("Number is greater")
	} else {
		println("Number is equal")
	}

	// Another expression for Conditionals
	var answerString: String = if (count < 50) {
		"Number is less"
	} else if (count > 50) {
		"Number is greater"
	} else {
		"Number is equal"
	}
	println(answerString)

	// when expression for Conditionals
	var answer: String = when {
		count == 50 -> "Number is equal"
		count < 50 -> "Number is less"
		else -> "Number is greater"
	}
	println(answer)

	// 				3. For Loops
	// Simple for Loop
	for (i in 1..10) {
		println(i)
	}

	// Another for loop
	for (i in 6 downTo 0 step 2) println(i)

	// for (item in collection) print(item)

	// 				4. Array Declaration

	var array1 = arrayOf(1, 2, 3, 4)	// Implicit Declaration
	for (i in 0..array1.size-1) println(array1[i])

	var array2 = arrayOf<Int>(1, 2, 3, 4)	// Explicit Declaration
	for (i in 0..array2.size-1) println(array2[i])

	var array3 = Array(5, {i -> i*10})	// Declaration using constructor
	for (i in 0..array3.size-1) println(array3[i])

	// 				5. Functions
	var ans1 = getValueComaprison1()		// call non-argument function
	println(ans1)

	var ans2 = getValueComaprison2(48)		// call argument function
	println(ans2)

	var ans3 = getValueComaprison3(123)		// call argument function
	println(ans3)

	var ans4 = stringLengthFunc("This is Shivam")
	println(ans4)

	// 				6. Classes

	val person = Person("Deepak", "Shakyawar")		// Define object of class
	print(person.last_name)
	print(person.addTwoNumbers())


}

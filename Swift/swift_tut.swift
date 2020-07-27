
print("This is Shivam") // Simple print statement in Swift

//--------------------------------- Variable Declaration ---------------------------------//

var myVariable = "This is Shivam"   // var is used to define variables
let apple = "Apple"                 // let is used to define constants
let mango = "Mango"                 // constant

let myAge = 24
let months = 0.5
print("My age is \(Double(myAge) + months)")         // Prinitng variable in template
print("My age is " + String(myAge)) // Implicit conversion in Swift

let quotation = """
Finish this tutorial as soon as possible.
So that you can move to another tutorial.
"""
print(quotation)

//--------------------------------- Arrays Declaration ---------------------------------//
var names = [String]() // initialising empty array
names.append("Shivam")  // appending entry in array
names.append("Somveer")
print(names)

var names2 = ["Shivam", "Somveer"]
names2.append("Lakhan")
print(names2)

//--------------------------------- Dictionary Declaration ---------------------------------//

var emptyDictionary = [String : String]()
emptyDictionary["employee1"] = "Shivam Pal"
emptyDictionary["employee2"] = "Somveer Sharma"
print(emptyDictionary)

let employees = [
    "employee1": "Shivam Pal",
    "employee2": "Somveer Sharma"
]
print(employees)

//--------------------------------- Control Flow ---------------------------------//

let individualScores = [12, 14, 56, 32, 57, 78, 29]
var teamScore = 0

for score in individualScores {
    if score > 50 {
        teamScore += 30
    } else {
        teamScore += 10
    }
}

for i in 0..<4 {
    print(i)
}

print(teamScore)

// nil Statements
var optionalString: String? = nil
print(optionalString == nil)

// If-else Statement
if let name = optionalString {
    print("There is a name: \(name)")
} else {
    print("string is empty")
}

// Switch statement
let marks = 40
switch marks {
case 30:
    print("Student is fail")
case 40:
    print("Student is passed")
default:
    print("Can't comment")
}

// While statement
var n = 2
while n < 100 {
    n *= 2
}
print(n)


// Repeat-while statement
repeat {
    n *= 3
} while n < 1000
print(n)


//--------------------------------- Functions and Closures ---------------------------------//

func greet(person: String, date: Int) -> String {
    return "Hello \(person), today is \(date)"
}
var str = greet(person: "Shivam", date: 21)
print(str)

func get_stats(scores: [Int]) -> (min: Int, max: Int, sum: Int) {
    var min = scores[0]
    var max = scores[0]
    var sum = 0
    
    for score in scores {
        sum += score
        if min < score {
            min = score
        }
        if max > score {
            max = score
        }
    }
    return (min, max, sum)
}

var stats_arr = [5, 3, 100, 3, 9]
let stats = get_stats(scores: stats_arr)
print(stats, stats.min, stats.max, stats.sum)


//--------------------------------- Objects and Classes ---------------------------------//

class NamedShape {
    var numberOfSides = 0
    var name: String
    
    init (name: String) {
        self.name = name
        numberOfSides = 12
    }
    
    func getDescription () -> String{
        return "\(name) has total \(numberOfSides) slides"
    }
}

var shape = NamedShape(name: "Shivam")
shape.numberOfSides = 100
print(shape.getDescription())


class Square: NamedShape {
    var sideLength: Double

    init (sideLength: Double, name: String) {
        self.sideLength = sideLength
        super.init(name: name)
        numberOfSides = 50
    }
    
    func area() -> Double {
        return sideLength * Double(numberOfSides)
    }
    
    override func getDescription() -> String {
        return "Hi \(name), there are \(numberOfSides) sides and sidelength is \(sideLength)"
    }
}

var quad = Square(sideLength: 4.0, name: "Shivam")
print(quad.area())
print(quad.getDescription())


//--------------------------------- Strucutures ---------------------------------//

struct Card {
    var rank: Int
    var suit: String
    
    func cardDescription () -> String {
        return "The \(rank) of \(suit)"
    }
}

let threeOfSpades = Card(rank: 3, suit: "spades")
let threeOfSpadesDescription = threeOfSpades.cardDescription()
print(threeOfSpadesDescription)

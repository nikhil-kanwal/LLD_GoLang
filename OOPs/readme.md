# **What is Object-Oriented Programming (OOP)?**

Object-Oriented Programming (OOP) is a programming paradigm based on the concept of "objects", which can contain data (fields or attributes) and code (methods or functions). 
Data in an object refers to the characteristics or properties of the object. These are typically referred to as fields, attributes, or variables within the object.
For example, if you have a Car object, the data might include fields like make, model, color, and year.

Code in an object refers to the behavior or actions that the object can perform. These are called methods (or functions in some languages) and are defined within the object to interact with the data.
Methods represent the behavior of the object. For example, a Car object might have methods like start(), stop(), and accelerate() that define how the car behaves when those actions are called.

## Core Principles of OOP:

**Encapsulation:**

Encapsulation is the bundling of data and methods that manipulate that data into a single unit, known as an object. It restricts direct access to some of the object’s components, ensuring that data is manipulated only through the defined methods. This helps protect the internal state of an object and ensures control over how the data is accessed and modified.

-> Grouping of fields and methods into a single unit which is an Object.

-> Manipulating data through Getters and Setters

```python
class Car:
    def __init__(self, make):
        self.__make = make  # Private attribute (name-mangled)
    
    def get_make(self):
        return self.__make  # Controlled access via getter
    
    def set_make(self, make):
        self.__make = make  # Controlled modification via setter

my_car = Car("Toyota")
print(my_car.get_make())  # Accessing via getter method

```

**Abstraction:**

Abstraction involves hiding the complex implementation details and exposing only the necessary and relevant parts of an object or method. It helps in reducing complexity by allowing the programmer to focus on interactions with simpler interfaces.

```python
class Animal:
    def speak(self):
        pass  # Abstract method to be implemented by subclasses
    
class Dog(Animal):
    def speak(self):
        return "Woof!"

class Cat(Animal):
    def speak(self):
        return "Meow!"
```

**Inheritance:**

Inheritance allows one class (child or derived class) to inherit properties and behaviors (methods) from another class (parent or base class). This promotes code reusability and establishes a relationship between the parent and child classes.

```python
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        return "Some generic sound"

class Dog(Animal):
    def speak(self):
        return "Woof!"

dog = Dog("Buddy")
print(dog.speak())  # Output: Woof!
```

**Polymorphism:**

Polymorphism describes the concept that you can access objects of different types through the same interface. Each type can provide its own independent implementation of this interface. Here by interface we don't mean GoLang interface but the defination of the methods.
There are two types of Polymorphism :
1. Compile-time(Method Overloading) - Allows you to implement multiple methods within the same class that use the same name. But the parameters should be different.

2. Run-time(Method OverRiding) - This happens during inheritance when within an inheritance hierarchy, a subclass can override a method of its superclass, enabling the developer of the subclass to customize or completely replace the behavior of that method.
Doing so also creates a form of polymorphism. Both methods implemented by the super- and subclasses share the same name and parameters. However, they provide different functionality.

Static(compile-time) polymorphism
```c++
#include <iostream>
using namespace std;

class Calculator {
public:
    // Method overloading: Same method name with different signatures (number/type of parameters)
    int add(int a, int b) {
        return a + b;
    }
    
    double add(double a, double b) {
        return a + b;
    }
    
    int add(int a, int b, int c) {
        return a + b + c;
    }
};

int main() {
    Calculator calc;

    // Static/Compile-time polymorphism: the appropriate add method is selected based on the arguments
    cout << calc.add(3, 4) << endl;        // Calls add(int, int) → Output: 7
    cout << calc.add(2.5, 3.5) << endl;    // Calls add(double, double) → Output: 6.0
    cout << calc.add(1, 2, 3) << endl;     // Calls add(int, int, int) → Output: 6

    return 0;
}

```

Run-time(Method OverRiding)

```python
class Animal:
    def speak(self):
        return "Some generic animal sound"

class Dog(Animal):
    def speak(self):
        return "Bark!"

class Cat(Animal):
    def speak(self):
        return "Meow!"

# Dynamic polymorphism in action
def animal_speak(animal):
    print(animal.speak())

# Creating objects of Dog and Cat classes
dog = Dog()
cat = Cat()

# Depending on the object type, the overridden method is called
animal_speak(dog)  # Output: Bark!
animal_speak(cat)  # Output: Meow!
```


------------------------------
# OOPs in goLang

Go is not a class-based object-oriented language; it does not have classes, inheritance, or explicit support for some of the typical OOP features found in languages like Java or C++. Instead, Go uses a combination of structs and interfaces to achieve some of the principles of OOP.

Here’s how Go approaches key OOP concepts:

**Encapsulation:**
In Go, encapsulation is achieved through structs. A struct is a composite data type that groups together variables (fields) under a single name. These fields can have different access levels, depending on whether their names start with an uppercase or lowercase letter (exported or unexported).
Methods: Go allows methods to be defined on structs, which provides a way to define behavior associated with a particular type.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Method on Person
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Greet())
}
```


**Inheritance:**
Go does not support inheritance directly. Instead, Go promotes composition over inheritance. You can achieve similar results using embedded structs. An embedded struct allows you to include one struct within another, giving you access to the methods and fields of the embedded struct.
Anonymous Fields: Go supports anonymous fields (also called embedded fields) which can be used to compose structs and leverage their functionality.

```go
package main

import "fmt"

type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Age     int
    Address // Embedded struct
}

func main() {
    p := Person{
        Name:    "Alice",
        Age:     30,
        Address: Address{City: "Bangalore", State: "KA"},
    }
    fmt.Println(p.Name, p.Address.City)
}
```


**Polymorphism:**
Go uses interfaces to provide polymorphic behavior. An interface defines a contract with methods but does not provide implementations. Any type that implements all the methods of an interface is said to satisfy that interface. This allows for a high degree of flexibility and abstraction.


```go

package main

import "fmt"

type Greeter interface {
    Greet() string
}

type Person struct {
    Name string
}

func (p Person) Greet() string { // different implementation of same method
    return "Hello, my name is " + p.Name
}

type Animal struct {
    Species string
}

func (a Animal) Greet() string { // different implementation of same method
    return "I am a " + a.Species
}

func greetAll(g Greeter) { // Since the parameter of this function is an interface any method which implements this interface can be sent as an argumnet to this method.
    fmt.Println(g.Greet())
}

func main() {
    p := Person{Name: "Alice"}
    a := Animal{Species: "Dog"}

    greetAll(p)
    greetAll(a)
}


```
**Abstraction:**
Abstraction is a core concept in object-oriented programming (OOP) that focuses on hiding the implementation details of a system and exposing only the essential features to the user. It allows you to interact with objects or systems in a more simplified and generalized manner, without needing to understand the complexities of their implementation.

In Go, abstraction is achieved primarily through interfaces. Interfaces allow you to define a contract (a set of methods) without specifying how these methods are implemented. This allows different types to be used interchangeably as long as they adhere to the same interface, which helps in reducing the dependency on concrete implementations and promotes flexibility

```go

package main

import "fmt"

// Define the Shape interface
type Shape interface {
    Area() float64
}

// Define the Rectangle struct
type Rectangle struct {
    Width, Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Define the Circle struct
type Circle struct {
    Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shape

    s = Rectangle{Width: 10, Height: 5}
    fmt.Println("Rectangle Area:", s.Area())

    s = Circle{Radius: 7}
    fmt.Println("Circle Area:", s.Area())
}

```




**What is concurrency**

* Concurrency is the capability to deal with lots of things at once.
* Let's consider a person jogging. 
* During his morning jog, lets say his shoe laces become untied. 
* Now the person stops running, ties his shoe laces and then starts running again. 
* This is a classic example of concurrency.
* The person is capable of handling both running and tying shoe laces, that is the person is able to deal with lots of things at once

**What is parallelism and how is it different from concurrency?**

* Parallelism is doing lots of things at the same time.
* lets assume that the person is jogging and also listening to music in his iPod.
* In this case the person is jogging and listening to music at the same time, that is he is doing lots of things at the same time. 
* This is called parallelism.

**What is an interface?**

* When a type provides definition for all the methods in the interface, it is said to implement the interface
* WashingMachine can be an interface with method signatures Cleaning() and Drying(). 
* Any type which provides definition for Cleaning() and Drying() methods is said to implement the WashingMachine interface.

**Empty interface**

* An interface that has zero methods is called an empty interface. 
* It is represented as interface{}. 
* Since the empty interface has zero methods, all types implement the empty interface.

**Methods in Go**

* A method is just a function with a special receiver type between the func keyword and the method name
* Methods with the same name can be defined on different types whereas functions with the same names are not allowed.
* Let's assume that we have a Square and Circle structure. 
* It's possible to define a method named Area on both Square and Circle. 

**rune**

* Rune literals are just 32-bit integer values (however they're untyped constants, so their type can change). 
* They represent unicode codepoints. 
* For example, the rune literal 'a' is actually the number 97



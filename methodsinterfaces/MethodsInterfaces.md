
- A method is a function with a receiver argument
    - func(recv<type>)fn()
    - the type with the method belongs to is the receiver. 

- Type has to be in the same package

- Pointer receiveres use dwht the method needs to change the values t which th reciever points
- Pointers save memory because the pass the memory addresses around. 


Interfaces

- A set of method signatures
- Any type that implemnets interface methods is considiered to be a child of the interface
- A value of inerface type can hold any value that implements these methods. 
- Pointers and functions implement interfaces too since they are all types. 
- Go expects you to use built-in interfaces to expend functionality : Example Stringers
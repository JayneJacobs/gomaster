# Effective Go

- Cross-compiling
- Go tools


https://golang.org/doc/effective_go.html

func init: 
   Used to insure the program state is good before running. 

   Use before main function


 _  net/http/pprof // to import a package only for side effects like init

 type switch; 
    Used to discover dynamic type of a variable; 

    var t interface{}

    switch t := t.(type) {
        default:
            fmt.printf("unexpected type %T\n", t) /// prints the type
        case bool:
            fmt.printf("boolean %t\n", t)
        case int:
            fmt.printf("boolean %t\n", t)
        case *bool:
            fmt.printf("pointer to boolean %t\n", *t)
        case *int:
            fmt.printf("pointer to boolean %d\n", *t)
        
    }


    compare types Example
    str := value.(string)// true or false but will panic

      str, ok := value.(string)
      if ok {
          fmt.Printf("string value is %q\n", str)
      } else {
          fmt.Printf(" not a string value \n", str)
      }
          // true or false


    Recover

    Equivalent to try catch

    if err := recover()





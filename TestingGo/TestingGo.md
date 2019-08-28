# Testing
- Unit testing 
- Benchmarking
- Go tools


##  Go Unit Test Steps

1. Create files appended with _test.go 
2. Import test package
3. Write functions with signature TestXxxx(*test.T) Uppercase T
4. Run Tests with go test tools


## Benchmarking

- Measure code performance
- Run code N times, obtain average run time


func Benchmarkxxx(b *testing.B){
    for i := 0; i< b.N; i++ {
        fmt.Sprintf(hello)
    }
}


### Benchmarking Tests
- Create files appended with _test.go for testing
- Import test package
- Write functions with signature BenchmarkXxxx(*test.B)
- Use go test -bench={wildcard}
    o -bench=. //runs all menchmarks


##Median calculation benchmark
Run the benchmark with to get memory too

    go test -benchtime=10000x -bench=. -test.benchmem=true
    
*not specifying **benchtime** can cause the test to run very long due to an issue with go benchmarking [issue](https://github.com/golang/go/issues/27217#issuecomment-453829330)

    
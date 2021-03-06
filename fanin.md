[Home/](https://github.com/harryosmar/go-playground/blob/master/concurrency.md) [Prev/](https://github.com/harryosmar/go-playground/blob/master/pipeline.md)

![fanin](https://github.com/harryosmar/go-playground/blob/master/resources/fanin.png)

> A function can **read from multiple inputs** and proceed until all are closed. This is called fan-in.

> How to achieve that : By multiplexing the input channels onto a single channel that's closed when all the inputs are closed.


Source codes :
- [Fain Code](https://github.com/harryosmar/go-playground/blob/master/actions/fanin.go)
- Without fan in
    ```
    curl --location --request GET 'http://localhost:9091/api/routine/fan/in/no'
    
    // output 
    // the output is ordered Joe first then Ann. Because there is channel block on receive
    Joe 0
    Ann 0
    Joe 1
    Ann 1
    Joe 2
    Ann 2
    Joe 3
    Ann 3
    Joe 4
    Ann 4
    
    // time 5.01 seconds
    ```

- With fan in
    ```
    curl --location --request GET 'http://localhost:9091/api/routine/fan/in/yes'
    
    // output is random. Who's come first .
    Ann 0
    Joe 0
    Joe 1
    Ann 1
    Ann 2
    Joe 2
    Ann 3
    Joe 3
    Ann 4
    Joe 4
    
    // time 5.01 seconds
    ```
-  [With fan in + parallel](https://github.com/harryosmar/go-playground/blob/master/actions/fanin_parallel.go)
    - time 1.08 seconds

## links
- [pipeline](https://blog.golang.org/pipelines)
- [Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs&t=938s)
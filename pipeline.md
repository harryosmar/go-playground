[Home/](https://github.com/harryosmar/go-playground/blob/master/concurrency.md) [Prev/](https://github.com/harryosmar/go-playground/blob/master/with_concurrency_and_parallel.md) [Next](https://github.com/harryosmar/go-playground/blob/master/fanin.md)

> a pipeline is a series of stages connected by channels

In each stage, the goroutines :
- **receive** values from upstream via **inbound channels**
- perform some function on that data, usually producing **new values**
- **send** values downstream via **outbound** channels

Source codes :
```
curl --location --request GET 'http://localhost:9091/api/routine/pipeline'
```

- [utils.Gen](https://github.com/harryosmar/go-playground/blob/master/utils/gen.go) execution time is 1 second 
- [utils.Sq](https://github.com/harryosmar/go-playground/blob/master/utils/sq.go) execution time is 1 second
- [Pipeline Code](https://github.com/harryosmar/go-playground/blob/master/actions/pipeline.go)
    ```go
      for n := range utils.Sq(utils.Gen(2, 3, 4, 5, 6)) {
      		fmt.Println(n)
      }
    ```
- Time : 6.03 Secons
- Console Output
    ```
    4
    9
    16
    25
    36
    ```
  
## links
- [pipeline](https://blog.golang.org/pipelines)
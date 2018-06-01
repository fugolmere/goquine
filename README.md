# goquine
Self-Replicating Program

This repo provides a solution to the Igneous [golang-contest](https://www.igneous.io/golang-contest)

Complete the following [quine](https://en.wikipedia.org/wiki/Quine_(computing)):
```go
package main
import quine "fmt"
var q = "package main\n\nimport quine \"fmt\"\n\n[...?]
```

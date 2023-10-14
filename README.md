# Go Concurrency Patterns and Usage Guide

This repository contains code examples and explanations of how to effectively use Go's concurrency features, including channels, goroutines, and the `select` statement. It also covers some common concurrency patterns.

## Table of Contents

- [Go Concurrency Patterns and Usage Guide](#go-concurrency-patterns-and-usage-guide)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Goroutines](#goroutines)
  - [Channels](#channels)
  - [Select Statement](#select-statement)
  - [Common Concurrency Patterns](#common-concurrency-patterns)
    - [Fan-Out, Fan-In](#fan-out-fan-in)
    - [Done Channel](#done-channel)
    - [Pipelines](#pipelines)

## Introduction

Go is known for its powerful concurrency support, and it's essential to understand how to use goroutines, channels, and the `select` statement effectively to harness the full power of Go's concurrent programming capabilities.

## Goroutines

Goroutines are lightweight, concurrent functions that can be executed concurrently with other goroutines. To create a goroutine, use the `go` keyword.

```go
package main

import (
    "fmt"
)

func main() {
    go func() {
        fmt.Println("Hello from a goroutine!")
    }()

    // Wait for the goroutine to complete
    fmt.Println("Main function")
}
```

## Channels

Channels are a way to communicate and synchronize goroutines. They provide a safe and efficient way to exchange data.

```go
package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)

    go func() {
        ch <- 42 // Send a value into the channel
    }()

    value := <-ch // Receive a value from the channel
    fmt.Println("Received:", value)
}
```

## Select Statement

The `select` statement allows you to wait on multiple channel operations simultaneously. It is often used to build non-blocking communication between goroutines.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- 42
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- 100
    }()

    select {
    case value := <-ch1:
        fmt.Println("Received from ch1:", value)
    case value := <-ch2:
        fmt.Println("Received from ch2:", value)
    }
}
```

## Common Concurrency Patterns

### Fan-Out, Fan-In

Fan-Out, Fan-In is a pattern where you have multiple goroutines reading from a channel (Fan-Out) and one goroutine writing to a channel (Fan-In). This is often used to process data concurrently.

```go
package main

import (
    "fmt"
)

func main() {
    data := []int{1, 2, 3, 4, 5}

    // Fan-Out
    ch := make(chan int)
    for _, d := range data {
        go func(d int) {
            ch <- d
        }(d)
    }

    // Fan-In
    go func() {
        for val := range ch {
            fmt.Println("Received:", val)
        }
    }()

    // Allow some time for goroutines to finish
    time.Sleep(2 * time.Second)
}
```

### Done Channel

A "Done" channel is used to signal the completion of multiple goroutines. It ensures that you wait for all goroutines to finish before exiting the program.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    done := make(chan struct{})

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            defer fmt.Println("Goroutine", id, "done")
            // Your goroutine logic here
        }(i)
    }

    go func() {
        wg.Wait()
        close(done)
    }()

    <-done // Wait for all goroutines to complete
}
```

### Pipelines

Pipelines are a way to chain multiple processing stages using channels. Data flows through a series of processing steps, making it easy to build complex data processing pipelines.

```go
package main

import (
    "fmt"
)

func main() {
    data := []int{1, 2, 3, 4, 5}

    // Stage 1
    in := make(chan int)
    out := make(chan int)

    go func() {
        for d := range in {
            out <- d * 2
        }
        close(out)
    }()

    // Stage 2
    go func() {
        for d := range out {
            fmt.Println("Processed value:", d)
        }
    }()

    // Feed data into the pipeline
    for _, d := range data {
        in <- d
    }
    close(in)

    // Wait for the pipeline to finish
    // ...
}
```

These are some of the fundamental concepts and patterns for working with concurrency in Go. Experiment with these examples and explore more advanced patterns and use cases to unlock the full potential of Go's concurrency features.

Happy coding!

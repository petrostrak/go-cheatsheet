### Race Condition

A race condition occurs when two or more operations must execute in the correct order, but the programm has not been written so that this order is guaranteed to be maintained.

Most of the time, this shows up in hat's called a data race; when one concurrent operation attempts to read a variable while at some undetemined time another concurrent operation is attempting to write to the same variable.

In this example, if we build the project with the `-race` paramenter and run it, we have the following output:
```
==================
WARNING: DATA RACE
Write at 0x00c00019c018 by goroutine 7:
  main.main.func1()
      /Users/petrostrak/Documents/github.com/petrostrak/go-cheatsheet/Concurrency/01-An-Introduction-to-Concurrency/01-Race-Condition/main.go:10 +0x44

Previous read at 0x00c00019c018 by main goroutine:
  main.main()
      /Users/petrostrak/Documents/github.com/petrostrak/go-cheatsheet/Concurrency/01-An-Introduction-to-Concurrency/01-Race-Condition/main.go:13 +0xb8

Goroutine 7 (running) created at:
  main.main()
      /Users/petrostrak/Documents/github.com/petrostrak/go-cheatsheet/Concurrency/01-An-Introduction-to-Concurrency/01-Race-Condition/main.go:9 +0xae
==================

```
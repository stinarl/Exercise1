Exercise 1 - Theory questions
-----------------------------
 
 ### What is the difference between concurrency and parallelism?
 > Concurrency is when several processes makes progress at the same time, not necessarily executing at the same time. Parallellism is when two or more processes are working in parallell. Actually executing at the same time.
 
 ### Why have machines become increasingly multicore in the past decade?
 > several cores improves the working speed, since it means that the machine can work on several datathreads at the same time.

 
 ### Why do we divide software (programs) into concurrently executing parts (like threads or processes)?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Concurrency is used to make it possible to split programs into several smaller programs/functions and have them working together through a main program.
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > It is easier for testing and finding flaws, since each part of the program can be tested alone. 
 
 ### What is the conceptual difference between threads and processes?
 > What is the conceptual difference between threads and processes?
 > A thread is a datastream executing something. A process can consist of multiple threads working together. (Connected in a main thread)

 
 ### Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they?
 > A fibre is a way of multitasking. A running fibre must allow another fibre to run.
 A coroutine is a routine that has the ability to stop executing in the middle of a function and then return to it later. It can also call other coroutines.

 
 ### What is the Go-language's "goroutine"? A C/POSIX "pthread"?
 > *In go goroutines is organized through channels.
 In c multithreading is not supported. A pthread is language independent  and gives the opportunity to run several threads at he same time.

 
 ### In Go, what does `func GOMAXPROCS(n int) int` change? 
 > It deefines how many prosesses that will be working at the same time. at the same time.



 
 
 
 

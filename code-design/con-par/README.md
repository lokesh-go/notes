#### Concurrency

- Concurrency is about multiple tasks which start, run, and complete in overlapping time periods, in no specific order.
- Concurrency is when two or more tasks can start, run, and complete in overlapping time periods. It doesn't necessarily mean they'll ever both be running at the same instant. For example, multitasking on a single-core machine.
- Concurrency is about dealing with lots of things at once.


#### Parallelism

- Parallelism is about multiple tasks or subtasks of the same task that literally run at the same time on a hardware with multiple computing resources like multi-core processor.
- Parallelism is when tasks literally run at the same time, e.g. on a multicore processor.


<p align="center"><img src="https://github.com/lokesh-go/notes/assets/31778886/aaa0cf1f-79a1-4993-adee-d568b46904f1" alt="Con-vs-Par" width="600px"/></p>

---

#### Concurrency in Golang

- Golang achieves concurrency by means of Goroutines.


#### Goroutines

- It's a special inbuilt feature in Golang which helps run two or more methods independently and simultaneously in connection with any other Goroutines present in a program.
- Every concurrent process in a Golang application can be called a Goroutine.
- Goroutine is a function that runs simultaneously with other functions.
- The cost of creating Goroutines is very small compared to the creation of an operating system thread.
- Every Golang Application has at least one Goroutine and that Goroutine is known as the main Goroutine.
- All the Goroutines are working under the main Goroutines if the main Goroutine is terminated, then all the Goroutines present in the program are also terminated.


##### There are two important Concurrency Golang Concepts

- ###### Goroutines
- ###### Channels
  - A channel is a highway used for sending and receiving data between Goroutines. Channels provide a way for one goroutine to send structured data to another.

---

##### How Goroutines Work?

Goroutines allow an application to become asynchronous in nature. Goroutines are different from threads and are managed by Go runtime.

When there is more than one Goroutine running in a Golang application, These Goroutines are propagated to the Go runtime scheduler which manages the lifecycle and allocates the Goroutine with an OS thread and memory.

When a Goroutine blocks the Go scheduler it performs a **context-switch**. It is the process of **storing the state of the Goroutine** so that it can be restored and execution can start at a later point. So while one goroutine is stored in the background the Go scheduler gives the thread to another goroutine for execution.

---

##### Locking

Anytime two or more Goroutines are dealing with the same data, and that data changes frequently, it can cause a condition called **"racing"**.

The problem arises when two Goroutine tries to update the value of shared resource at the same time. This can create unwanted scenarios. To avoid such problems we use locking with the help of Mutex.

A **Mutex** is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at a given point in time.

---

##### Waiting

In Golang if the execution of the main Goroutine is completed all the other Goroutines in the programs are also stopped at the same time. To Avoid this scenario we can use waiting with the help of a Waitgroup. 

A **WaitGroup** is actually a type of counter which blocks the execution of Goroutines, until its internal counter becomes 0.


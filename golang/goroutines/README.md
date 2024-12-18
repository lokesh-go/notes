# Goroutines

### What is Goroutines?
- A goroutine is a lightweight execution thread in the Go programming language and a function that executes concurrently with the rest of the program.
- Goroutines are incredibly cheap when compared to traditional threads as the overhead of creating a goroutine is very low. Therefore, they are widely used in Go for concurrent programming.
- A goroutine is a function that executes simultaneously with other goroutines in a program and are lightweight threads managed by Go.

#### Some Points
- Goroutines enable concurrent execution of functions or methods in Go.
- Goroutines are lightweight, independently scheduled functions in Go.
- Goroutines are created using the go keyword before a function call.
- Goroutines are asynchronous and run concurrently without blocking other Goroutines.
- Goroutines are managed by the Go runtime.
- Goroutines communicate and synchronize through channels for coordination.
- The go runtime manages lightweight threads called goroutines.


#### Some Details
- A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space.
- It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.
- Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run.
- Prefix a function or method call with the go keyword to run the call in a new goroutine. When the call completes, the goroutine exits, silently. (The effect is similar to the Unix shell's & notation for running a command in the background.)


### What is Thread?

- A thread is a flow of execution through the process code, with its own **program counter** that keeps track of which instruction to execute next, **system registers** which hold its current working variables, and a **stack** which contains the execution history.
- Each thread belongs to exactly one process and no thread can exist outside a process. Each thread represents a separate flow of control.

### Goroutines Vs Thread
- Goroutines are cheaper than threads.
- Threads consume a lot of memory due to their large stack size (≥ 1MB). So creating 1000s of threads means you already need 1GB of memory.
- Goroutine created with initial only 2KB of stack size.

### When to use Goroutines?
Goroutines are useful when you want to do multiple things **simultaneously**. For example, if you have ten things you want to do at the same time, you can do each one on a separate goroutine, and wait for all of them to finish.

### Why goroutines are lightweight?
- It's because a goroutine starts with a stack space of 2KB which is extremely smaller and more compact than OS thread's fixed-size stack space of 2MB. However,  Go runtime dynamically manages the stack size of Goroutines based on the program's needs.
- Goroutines are managed by the Go runtime and not by the underlying operating system. It allows Go runtime to have more control over scheduling Goroutines and enables efficient concurrency management without relying solely on operating system threads.

### What happens when context switching will be in Goroutines

When a Goroutine is preempted, its execution is temporarily suspended by the Go scheduler to allow other Goroutines to run. To ensure that the Goroutine can resume its execution correctly after being preempted, the Go runtime **stores** its state, including the completion task process, in the **Goroutine's context**.

The key components of a Goroutine's Context are:
- **Stack**
  - Each Goroutine has its own stack, which is a fixed-size memory region used for storing local variables and function call frames.
  - When a Goroutine is preempted, the current state of its stack, including the function call frames, local variables, and the point of execution, is saved.
- **Program Counter (PC)**
  - The program counter is a register that keeps track of the current instruction being executed by the Goroutine.
  - When a Goroutine is preempted, the value of the program counter is saved so that the Goroutine can resume execution from the correct point later.
- **Registers**
  - Goroutines also have various registers that hold intermediate results and other important information during their execution.
  - When preempted, the values of these registers are saved so that the Goroutine can pick up where it left off.
- **Goroutine State**
  - The Go runtime also stores information about the Goroutine's state, such as whether it is blocked on a channel operation, waiting for a lock (mutex), or runnable.

When a preempted Goroutine is rescheduled to run again, the Go scheduler restores the Goroutine's context from where it was saved, allowing it to continue its execution as if it was never interrupted. This process of saving and restoring a Goroutine's context is known as **"context switching."**

Context switching in Goroutines is very efficient compared to traditional operating system threads because Goroutines have small fixed-size stacks and the Go scheduler is designed to minimize the overhead of context switches. The Go runtime can efficiently manage thousands or even millions of Goroutines concurrently, making Go an excellent choice for writing highly concurrent applications.

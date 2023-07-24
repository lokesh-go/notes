# Memory Management

Local variables are stored in two primary locations: the stack and the heap:

- **Stack**
  - The stack is a region of memory that is used to store local variables and function call frames for each Goroutine.
  - Local variables declared in functions have their memory allocated on the stack.
  - When a function is called, a new stack frame is created for that function, and the function's local variables are stored in this frame.
  - The stack operates in a last-in-first-out (LIFO) manner, meaning that the most recently declared local variables are removed first when the function completes its execution.
  - Since the stack allocation and deallocation are fast and efficient, it is the default storage location for most local variables in Go.
- **Heap**
  - The heap is another region of memory used to store data that needs to have a longer lifetime than the current function call.
  - Variables allocated on the heap persist beyond the scope of the function that created them.
  - Variables stored on the heap must be explicitly managed (allocated and deallocated) by the programmer, as Go does not have automatic garbage collection for variables created on the heap.
  - Some examples of variables stored on the heap include objects created with new() or make(), and any variables that escape from the scope of a function (i.e., they are referenced by variables outside of the function).

Go's automatic memory management and efficient stack allocation for local variables contribute to the language's simplicity and performance. In general, it is recommended to use the stack for most local variables as it allows for faster memory access and automatic deallocation. Only use the heap when necessary, such as when you need to share data across functions or Goroutines or when the data's lifetime extends beyond the function's scope.


### Garbage Collector

The garbage collector (GC) is responsible for managing memory allocated on the heap.
Its primary role is to identify and reclaim memory that is no longer in use, which helps prevent memory leaks and ensures efficient memory utilization.
The garbage collector in Go employs a concurrent, tri-color, mark-and-sweep algorithm. 
Let's see how it works for both the stack and heap:
- **Stack Memory Management**
  - The stack memory management in Go is automatic and efficient. When a Goroutine's function is called, a new stack frame is created to store local variables and function call information
  - As soon as the function returns, the stack frame is automatically deallocated, and the memory is reclaimed. This process is known as **"automatic stack deallocation"** and is handled by the **Go runtime**.
  - Since the stack operates in a last-in-first-out (LIFO) manner, memory deallocation for local variables is very fast and does not require the garbage collector's intervention.
- **Heap Memory Management**
  - Memory allocated on the heap requires explicit management by the garbage collector because objects on the heap may have longer lifetimes and may be referenced by multiple parts of the program.
  - The garbage collector uses a concurrent mark-and-sweep algorithm to determine which objects on the heap are still in use and which can be reclaimed.
  - The garbage collector works in several phases:
    - **Mark Phase**
      - The garbage collector traverses the object graph, starting from known root objects (e.g., global variables, Goroutine stacks, etc.), and marks all reachable objects as **"live."**
    - **Sweep Phase**
      - After the marking phase, the garbage collector sweeps through the entire heap, identifying and freeing memory that was not marked as "live." These are the objects that are no longer in use and can be safely reclaimed.
  - The garbage collector in Go runs concurrently with the application, meaning it runs in the background while Goroutines continue executing. This minimizes the impact of garbage collection on the overall performance of the program.

Go's garbage collector is designed to strike a balance between efficiency and responsiveness. The concurrent and incremental nature of the garbage collector helps reduce pause times and ensures that the application's performance remains smooth even during garbage collection cycles. However, developers should still be mindful of creating unnecessary objects on the heap to avoid undue pressure on the garbage collector and ensure optimal memory management in their Go programs.

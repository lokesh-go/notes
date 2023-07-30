# Variables

##### Go manages variables in different ways depending on their type and scope:
- ###### Local Variables
  - Local variables are declared within a specific block or function scope.
  - They are stored on the stack, which is automatically managed by the Go runtime.
  - Local variables have limited lifetimes and are automatically deallocated when they go out of scope (i.e., when the function exits).
  - Each Goroutine has its own stack, allowing local variables to be independent and not affected by other Goroutines.
- ###### Package-Level Variables:
  - Package-level variables are declared at the package level (outside any function) and have a global scope within that package.
  - They are stored in the data segment of the program's memory.
  - Package-level variables have a lifetime equal to the entire program's execution.
- ###### Shared Variables (Heap-Allocated):
  - Shared variables are variables that need to be accessed and modified by multiple Goroutines.
  - They are usually allocated on the heap using **new()**, **make()**, or composite literals (e.g., slices, maps, channels) to create new instances.
  - The memory for shared variables on the heap is managed by the Go runtime, and it is automatically garbage collected when there are no more references to the variable (i.e., when it is no longer in use).
- ###### Function Parameters:
  - Function parameters are passed by value, meaning a copy of the value is made and used within the function.
  - If the value is a primitive type (int, float, bool, etc.), the copy is independent and will not affect the original variable outside the function.
  - If the value is a reference type (slice, map, channel, pointer, etc.), the copy will point to the same underlying data, and changes made within the function will affect the original variable.

Go manages memory allocation and deallocation for local variables on the stack, while it handles the memory allocation and garbage collection for shared variables on the heap. By managing memory in this way, Go allows for efficient and safe handling of variables in both concurrent and non-concurrent scenarios. Developers can focus on writing their code without worrying about manual memory management, making Go a user-friendly language with robust memory handling capabilities.
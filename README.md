# mempool

## Introduction

mempool is a simple memory pool implementation, which is used to improve memory utilization and performance. It provides a configurable-size cache pool for repeatedly allocating and recycling memory blocks of the same size.

The memory pool pre-allocates a certain number of memory blocks and stores them in a stack, avoiding frequent memory allocation and deallocation operations. When memory needs to be allocated, a usable memory block is directly taken from the pool; when the memory block is no longer needed, it is returned to the pool for future use.

mempool uses a lock-free structure to achieve concurrent safety of data access, avoiding the performance overhead of traditional lock mechanisms. It utilizes atomic operations and lock-free algorithms to ensure the correctness and efficiency of concurrent access.

The memory pool supports custom memory block generation functions fn, which can create memory blocks suitable for actual needs.

## Usage

The following is a basic example of using mempool:

Import the required package:

```go
import "github.com/lambertxiao/go-mempool"
```

1. Create a new memory pool:

```go
pool := mempool.NewGoMemPool(uint, func() interface{}) *GoMemPool // Pass the appropriate size and generation function according to actual needs
```

2. Retrieve a memory block from the memory pool:

```go
item := pool.Get() // If the method is called when the memory pool is empty, it will wait indefinitely until a memory block is returned to the pool.

or

item := pool.GetByTime(time.Second) // If an available memory block is not retrieved within the specified timeout time, the method returns a null value (zero value of the corresponding type).
```

3. Return the memory block to the memory pool:

```go
pool.Put(item) // Return item to the pool for future use
```

4. Check the size or capacity of the memory pool:

```go
capacity := pool.Cap() // Get the capacity of the memory pool
```

5. Destroy the memory pool (empty all memory blocks):

```go
pool.Destory() // Clear all memory blocks in the memory pool
```

### Pool
Pool is a concurrent-safe implementation of the object pool pattern. At a high level, the pool pattern is a way to create and make available a fixed number, or pool, of things for use. It's commonly used to constrain the creation of things that are expensive (e.g. DB connections) so that only a fixed number of them are ever created, but an indeterminate number of operations can still request access to these things. In the case of Go's `sync.Pool`, this data type can be safely used by multiple goroutines.

Pool's primary interface is its `Get` method. When called, Get will first check whether there are any available instances within the pool to return to the caller, and if not, call its `New` member variable to create a new one. When finished, callers call `Put` to place the instance they were working with back in the pool for use by other precesses.

When working with Pool, remember the following points:
* When instantiating `sync.Pool`, give it a New member variable that is thread-safe when called.
* When you receive an instance from `Get`, make no assumptions regarding the state of the object you receive back.
* Make sure to call `Put` when you're finished with the object you pulled out of the pool. Otherwise, the Pool is useless. Usually this is done with `defer`.
* Objects in the pool must be roughly uniform in makeup.
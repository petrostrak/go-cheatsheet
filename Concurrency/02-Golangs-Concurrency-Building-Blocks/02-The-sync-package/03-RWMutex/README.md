### RWMutex
The `sync.RWMutex` is conceptually the same thing as a Mutex: it guards access to memory; however, RWMutex gives you a little bit more control over the memory. You can request a lock for reading, in which case you will be granted access unless the lock is being held for writting. This means that an arbitrary number of readers can hold a reader lock so long as nothing else is holding a writer lock.
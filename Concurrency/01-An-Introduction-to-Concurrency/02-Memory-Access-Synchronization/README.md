### Memory Access Synchronization

In this example we've created a convention for developers to follow. Anytime developers want to access the data variable's memory, they must first call `Lock`, and when they're finished they must call `Unlock`. Code between those two statements can then assume it has exclusive access to data; we have successfully synchronized access to the memory.
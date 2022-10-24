### Deata protected by Confinement
In some sense, immutable data is ideal because it is implicitly concurrent safe. Each concurrent process may operate on the same data, but it may not modify it. If it wants to create new data, it must create a new copy of the data with the desired modifications. In Go, you can achieve this by writing code that utilizes copies of values instead of pointers to values in memory.

The techniques to confine concurrent values are a bit more involved than simply passing copies of values.

Confinement is the simple yet powerful idea of ensuring information is only ever available from one concurrent process. When this is achieved, a concurrent program is implicitly safe and no synchronization is needed. There are two kinds of confinement possible: 
* ad hoc and 
* lexical.

#### Ad hoc
Ad hoc confinement is when you achieve confinement through a convention— whether it be set by the languages community, the group you work within, or the codebase you work within.

#### Lexical
Lexical confinement involves using lexical scope to expose only the correct data and concurrency primitives for multiple concurrent processes to use. It makes it impossi‐ ble to do the wrong thing.

Why pursue confinement if we have synchronization available to us? The answer is improved performance and reduced cognitive load on developers. Synchronization comes with a cost, and if you can avoid it you won’t have any critical sections, and therefore you won’t have to pay the cost of synchronizing them.
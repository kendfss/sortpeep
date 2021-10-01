sortpeep
---

This package demonstrates a number of approaches to concurrency by way of sorting analysis.  
Specifically, it implements several approaches to synchronization and communication using golang's built in tools. Each one enables the user to "watch" an array being sorted and/or count the number of operations 
It is intended as a reference for those exceptionally dark times where internet access is painfully slow or altogether unavailable.


### Implementations
- Coordinate program exit
    - Channels
    - WaitGroups
- Counting
    - Mutual Exclusion
    - Atomic Operations

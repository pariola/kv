### kv  
  
> Build an interactive shell that allows access to a "transactional in-memory key/value store".  
  
Idea from: https://www.freecodecamp.org/news/design-a-key-value-store-in-go/  
  
|Command|Description|  
|--- |--- |  
|SET|Sets the given key to the specified value. A key can also be updated.|  
|GET|Prints out the current value of the specified key.|  
|DELETE|Deletes the given key. If the key has not been set, ignore.|  
|COUNT|Returns the number of keys that have been set to the specified value. If no keys have been set to that value, prints 0.|  
|BEGIN|Starts a transaction. These transactions allow you to modify the state of the system and commit or rollback your changes.|  
|END|Ends a transaction. Everything done within the "active" transaction is lost.|  
|ROLLBACK|Throws away changes made within the context of the active transaction. If no transaction is active, prints "No Active Transaction".|  
|COMMIT|Commits the changes made within the context of the active transaction and ends the active transaction.|  
  
  
Sample:
```shell  
➜ go run kv.go
   
kv> GET x  
x not set  
kv> SET x 123  
kv> GET x  
123  
kv> BEGIN  
kv> GET x x not set  
kv> SET x 111  
kv> GET x  
111  
kv> COMMIT  
kv> GET x  
111  
kv> BEGIN  
kv> SET x 000  
kv> GET x  
000  
kv> ROLLBACK  
kv> GET x  
x not set  
kv>   
```  
  
A `Transaction` simply isolates changes from the global store.  
  
  
Yada.
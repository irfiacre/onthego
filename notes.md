Base tutorial - https://youtu.be/8uiZC0l4Ajw?si=f1fqvkl0H5c0amqd

### Six main points about GO

1. Is statistically typed language:

This is how to declare a variable

```:go

var myVariable string = "some string"

# or

var myVariable = "some string";
```
2. Strongly typed
    - The operation you can perform on its variables depend on the type of the variable.
    - You can not make an array of different value types.

3. GO is compiled.
4. Fast time to compile
5. Builtin concurrency.
6. Simplicity


### Packages vs Modules
1. Package: a collection of GO files.
2. Module: a collection of packages.

Note: when you use the "package main", we are telling the compiler that the entry point is here.



### Understanding Integer types:

#### Integers
- int8:
    - Uses 1 byte of memory.
    - Stores [-128, 127] | unassigned store [0, 255]

- int16:
    - Uses 2 byte of memory.
    - Stores [-32,768, 32,767]

- int32:
    - Uses 3 byte of memory.
    - Stores [-2,147,483,648, 2,147,483,647]

- int64:
    - Uses 1 byte of memory.
    - Stores [-9,223,372,036,854,775,808,  -9,223,372,036,854,775,807]

** Unassigned integers (unint) use the same memory, but start from 0 (basically natural numbers).

#### Float

Go does not have a default float type, one has to specify whether you are using float32 or float64 (float64 is more consisce).


#### Arthimetic operations
- You can not add numbers of different types, say int64 and int32. of float and int.
- Division of ints results in and int round down. you can get the remainder using the "%" symbol.


#### Strings
- If you use `len(stringVar)` you get the byte length, not the number of characters. 
- To count the actual length, use hte unicode package
```:go
import "unicode/utf8"
....

var stringLen int = utf8.RuneCountInString(string)
```

- In GO, you can use `const myVar <varType> = <constant value>` the same way you use var, with one exception as that the value of a const can not be changed.
 

#### Arrays
- They have a fixed length.
- Always the same type.
- Indexable
- Stored Contiguous in Memory.

#### Slices
Just wrappers under an array.

#### Maps
- A set of key value pairs, essentially a dictionary in Python or an object in Typescript.
- When looking for a non existent key, the map return the default value for the type of that map.
- A map always returns something.

#### Strings & Runes
- Strings are immutable in GO.
- 

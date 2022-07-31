## Struct
> A data structure. Collection of properties that are related together

**Definition**
```go
type person struct {
    firstName string
    lastName string
    contact contactInfo
}

type person struct {
    firstName string
    lastName string
    contactInfo
}
```
> When embedding structs, if you leave off the property name of the embedded struct, then the property name will be the same as the struct name

**Instantiation**
```go
var alex person
alex.firstName = "Alex"
alex.lastName = "Anderson"

alex := person{"Alex", "Anderson"}

alex := person{firstName: "Alex", lastName: "Anderson"}

alex := person{
    firstName: "Alex", 
    lastName: "Anderson",
    contact: contactInfo{
        email: "alex@gmail.com",
        zip: 12345,
    },
}
```
**Printing**
```go
fmt.Printf("%+v", alex) // {firstName:Alex lastName:Anderson}
```

## Type Zero Values (default values)
| Type | Zero Value |
|------|------------|
| `string` | "" |
| `int` | 0 |
| `float` | 0 |
| `boolean` | false |

## Pointers
> Go is a pass by value language

```go
jim.updateName("Bob")
func (p person) updateName() {}
```
> `jim` and `p` have different memory addresses

> `p` is just a copy of `jim`

`&variable` Gives the memory address of the value this variable is pointing at
```go
jimPointer := &jim
```

`*pointer` Gives the value this memory address is pointing at
```go
jimValue := *jimPointer
```

```go
func (pointerToPerson *person) updateName(newName string) {
    (*pointerToPerson).firstName = newName 
}
```
`*person` is a type description, which means we are working with a pointer to a person

`*pointerToPerson` is an operator, which means we want to manipulate the value the pointer is referencing

```go
jimPointer := &jim
jimPointer.updateName("Bob")

jim.updateName("Bob")

// the two options above are equivalent when using the func with the receiver below
func (p *person) updateName() {}
```

**Gotchas**
-  When creating a slice, two data structures are created
    - A Slice, which contains
        - pointer to head of array
        - capacity: how many it can contain at present
        - length: current number of elements inside the slice
    - An Array, which contains the values
- Since Go is a pass by value language, a copy of the slice data structure is passed to a function, which has the same pointer to the array as the original slice data structure

| Value Types | Reference Types |
|-------------|-----------------|
| int | slices|
| float | maps |
| string | channels |
| bool | pointers |
| structs | functions |

`Value Types:` Use pointers to change these things in a function

`Reference Types:` Don't worry about points with these 

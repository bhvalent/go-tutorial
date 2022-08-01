## Maps
**Instantiation**
```go
colors := map[string]string{
    "red":   "#ff0000",
    "green": "#00ff00",
}

var colors map[string]string

colors := make(map[string]string)
```
**Modification**
```go
colors["white"] = "#ffffff"

delete(colors, "white")
```

**Iteration**
```go
for color, hex := range c {
    fmt.Println("Color: ", color, " Hex Code: ", hex)
}
```
> `color` would be the `key` and `hex` would be the `value` of the map

## Structs vs Maps
- Structs
    - Values can be of different types
    - Keys don't support indexing
    - Value type
    - You need to know all the different fields at compile time
    - Use to represent a "thing" with a lot of different properties
- Maps
    - All keys must be the same type
    - All values must be the same type
    - Keys are indexed - we can iterate over them
    - Reference type
    - Use to represent a collection of related properties
    - Don't need to know all the keys at compile time
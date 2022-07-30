## Basic Go Types
| Type | Example |
|------|---------|
| `bool` | true, false |
| `string` | "hi", "how's it going?" |
| `int` | 0, -1000, 99999 |
| `float64` | 10.0000001, 0.00009, -100.003 |

## Variable Creation
**Longform:**
```
var card string = "Ace of Spades"
```
**Inferenced:**
```
card := "Ace of Spades"
```
- `:=` is only used on initialization

## Function with return type
```
func newCard() string {
    return "Five of Diamonds"
}
```

## Lists
- **Array:** Fixed length list of things
- **Slice:** An array that can grow or shrink 
> Both must have the same data type for the entire list
## Slice
**Initialization**
```
cards := []string{"Ace of Diamonds"}
```
**Adding**
```
cards = append(cards, "Six of Spades")
```
> `append` returns a new splice

**Iteration**
```
for i, card := range cards {
    fmt.Println(i, card)
}
```

## Type declaration
```
type deck []string
```

## Receiver
```
func (d deck) print() {}
```
> Any variable of type `deck` now gets access to the `print` method
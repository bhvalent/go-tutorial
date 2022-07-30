## Types
**Basic Types (nonexlusive list)**
| Type | Example |
|------|---------|
| `bool` | true, false |
| `string` | "hi", "how's it going?" |
| `int` | 0, -1000, 99999 |
| `float64` | 10.0000001, 0.00009, -100.003 |

**Type Conversion**
```
[]byte("Hi there")
```
> `typeWeWant` ( `currentValue` )

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

## Functions
**with return type**
```
func newCard() string {
    return "Five of Diamonds"
}
```
**with multiple return types**
```
func methodName(paramName paramType) (returnType1, returnType2) {
	return valueWithReturnType1, valueWithReturnType2
}
```
**with receiver**
```
func (d deck) print() {}
```
> Any variable of type `deck` now gets access to the `print` method

## Lists
- **Array:** Fixed length list of things
- **Slice:** An array that can grow or shrink 
> Both must have the same data type for the entire list
## Slice
- are zero-indexed

**Initialization**
```
cards := []string{"Ace of Diamonds", "Five of Hearts", "Seven of Spades", "Queen of Clubs"}
```
**Accessing**
```
cards[1] // "Five of Hearts"
cards[0:2] // "Ace of Diamonds", "Five of Hearts"
cards[:2] // "Ace of Diamonds", "Five of Hearts"
cards[1:] // "Five of Hearts", "Seven of Spades", "Queen of Clubs"
```
> cards[`startIndexIncluding` : `upToNotIncluding`]

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

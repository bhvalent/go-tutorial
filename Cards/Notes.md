## Types
**Basic Types (nonexlusive list)**
| Type | Example |
|------|---------|
| `bool` | true, false |
| `string` | "hi", "how's it going?" |
| `int` | 0, -1000, 99999 |
| `float64` | 10.0000001, 0.00009, -100.003 |

**Type Conversion**
```go
[]byte("Hi there")
```
> `typeWeWant` ( `currentValue` )

## Variable Creation
**Longform:**
```go
var card string = "Ace of Spades"
```
**Inferenced:**
```go
card := "Ace of Spades"
```
- `:=` is only used on initialization

## Functions
**with return type**
```go
func newCard() string {
    return "Five of Diamonds"
}
```
**with multiple return types**
```go
func methodName(paramName paramType) (returnType1, returnType2) {
	return valueWithReturnType1, valueWithReturnType2
}
```
**with receiver**
```go
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
```go
cards := []string{"Ace of Diamonds", "Five of Hearts", "Seven of Spades", "Queen of Clubs"}
```
**Accessing**
```go
cards[1] // "Five of Hearts"
cards[0:2] // "Ace of Diamonds", "Five of Hearts"
cards[:2] // "Ace of Diamonds", "Five of Hearts"
cards[1:] // "Five of Hearts", "Seven of Spades", "Queen of Clubs"
```
> cards[`startIndexIncluding` : `upToNotIncluding`]

**Adding**
```go
cards = append(cards, "Six of Spades")
```
> `append` returns a new splice

**Iteration**
```go
for i, card := range cards {
    fmt.Println(i, card)
}
```

## Type declaration
```go
type deck []string
```

## Testing
> filename must end in `_test.go`


**To run tests:**
```go
go test
```
**This automates the execution of any function of the form:**
```go
func TestXxx(t *testing.T)
```
> where `Xxx` does not start with a lowercase letter

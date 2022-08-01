## Interfaces
- are not `generic` types
- are implicit
- are a contract to help us manage types

```go
type bot interface {
	getGreeting() string
}
```
> Any type in the program with a function called `getGreeting` that returns a `string`, is also type `bot`

| Concrete Type | Interface Type |
|---------------|----------------|
| map, struct, int, string, englishBot | bot

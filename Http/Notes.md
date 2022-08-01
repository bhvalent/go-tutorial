## Interfaces
> You can nest interfaces
```go
type io.ReadCloser interface {
    Reader
    Closer
}

type Reader interface {
    Read([]byte) (int, error)
}

type Closer interface {
    Close() (error)
}
```
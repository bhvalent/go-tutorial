## How to run
> Run the following command in the HelloWorld directory
```
> go run main.go
```

## Go CLI
| Command | Description |
|---------|-------------|
| `go build` | compiles a bunch of go source code files |
| `go run` | compiles and executes one or two files |
| `go fmt` | Formats all the code in each file in the current directory |
| `go install` | Compiles and installs a package |
| `go get` | downloads the raw source code of someone else's package |
| `go test` | runs any tests associated with the current project |

## Terms
**package:**
-  like a project or workspace
- a collection of common source code files
- the first line of a file must declare what package it belongs to
- two types of packages
    - **Executable:** Generates a file that we can run
        - has to be named `main` in order for the build command to generate an executable
        - must have a func called `main`
    - **Reusable:** Code used as helpers. Good place to put resusable logic. Like a shared library
```go
package main 
```

**import:** used to gain access to another package inside the one we are authoring
```go
import "fmt"
```

**func:** tells we are about to declare a function
```go
func main() {
    // function body
}
```

## File Oranization
1. Package Declaration
2. Import other packages needed
3. Declare functions, tell Go to do things
# tpl ✍️

Generate golang templates with JSON data 💪.

Using `go 1.16.3`.

Installation:

```bash
$ go install github.com/alextanhongpin/tpl
```

Usage:
```bash
$ tpl -in=in.tpl -out=out.go -data data.json

# To print to stdout without creating the files.
$ tpl -in=in.tpl -out=out.go -data data.json -dry-run
```



Content of `in.tpl`:
```
package main

func main() {
	fmt.Println("hello {{.msg}}")
}
```

The input file extension can also be `.txt`, `.tmpl`, it is up to you :smile:


Content of `data.json`:
```json
{
	"msg": "world"
}
```

Content of `out.go`:
```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

Note the `import "fmt"` - files with `.go` extensions are automatically formatted with packages imported.

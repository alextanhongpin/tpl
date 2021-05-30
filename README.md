# tpl

Generate golang templates with json data.


Usage:
```bash
$ tpl -in=in.tpl -out=out.go -data data.json
```



Content of `in.tpl` (the naming can be `.txt`, `.tmpl`, it is up to you):
```
package main

func main() {
	fmt.Println("hello {{.msg}}")
}
```

Content of `data.json`:
```json
{
	"msg": "world"
}
```

Note the `pascalcase` naming convention to allow golang templates to read the data.


Content of `out.go`:
```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

Note the `import "fmt"` - files with `.go` extensions are automatically formatted with packages imported.

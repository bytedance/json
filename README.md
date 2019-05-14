# json

JSON standard package extension library.

## Extension feature

1. In the case of a non-cyclic combination, using zero value('[]' or '{}') instead of null.

```go
func ExampleMarshalNotnull() {
	var a *A
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
	// Output:
	// {"List":[],"Map":{},"Children":[],"B":{"ListB":null,"MapB":null,"A":null,"B":null},"MapPtr":{}}
}

type A struct {
	List     []string
	Map      map[string]string
	Children []*A
	B        *B
	MapPtr   *map[string]string
}

type B struct {
	ListB []string
	MapB  map[string]string
	A     *A
	B     *B
}
```
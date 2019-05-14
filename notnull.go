package json

var defalutNotnull = true

// SetNotnull specifies whether in the case of a non-cyclic combination,
// using zero value('[]' or '{}') instead of null.
// NOTE:
//  defalut true
func SetNotnull(on bool) {
	defalutNotnull = on
}

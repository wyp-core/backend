package general

// InArr checks for value present in a slice / array
func InArr[T string | int | int16 | int32 | int64 | bool](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
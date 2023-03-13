package base

// GetMapKeys 获得map的所有key
func GetMapKeys[T comparable, U any](m map[T]U) []T {
	if len(m) == 0 {
		return nil
	}
	allKeys := make([]T, 0, len(m))
	for k := range m {
		allKeys = append(allKeys, k)
	}
	return allKeys
}

// GetMapValues 获得map的所有value
func GetMapValues[T comparable, U any](m map[T]U) []U {
	if len(m) == 0 {
		return nil
	}
	allKeys := make([]U, 0, len(m))
	for _, v := range m {
		allKeys = append(allKeys, v)
	}
	return allKeys
}

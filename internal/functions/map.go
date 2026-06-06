package functions

func Map[T any, R any](items []T, fn func(T) R) []R {
	result := make([]R, len(items))

	for i, item := range items {
		result[i] = fn(item)
	}

	return result
}

package strain

func Keep[T any](list []T, pred func(T) bool) []T {
	var keep []T

	for _, val := range list {
		if pred(val) {
			keep = append(keep, val)
		}
	}

	return keep
}

func Discard[T any](list []T, pred func(T) bool) []T {
	var keep []T

	for _, val := range list {
		if !pred(val) {
			keep = append(keep, val)
		}
	}

	return keep
}

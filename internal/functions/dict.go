package functions

func Dict(values ...any) map[string]any {
	if len(values)%2 != 0 {
		panic("invalid dict call: must have even number of arguments")
	}
	dict := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			panic("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict
}

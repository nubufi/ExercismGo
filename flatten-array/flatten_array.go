package flatten

func Flatten(nested interface{}) []interface{} {
	result := []interface{}{}

	for _, v := range nested.([]interface{}) {
		switch v.(type) {
		case []interface{}:
			result = append(result, Flatten(v)...)
		case nil:
			continue
		default:
			result = append(result, v)
		}
	}

	return result
}

package main

import ("encoding/json"; "fmt"; "strconv")

func findOrderArray(value any) []any {
	if rows, ok := value.([]any); ok {
		return rows
	}
	if object, ok := value.(map[string]any); ok {
		for _, key := range []string{"data", "pedidos", "orders", "results"} {
			if nested, ok := object[key]; ok {
				if rows := findOrderArray(nested); rows != nil {
					return rows
				}
			}
		}
	}
	return nil
}

func findOccurrenceArray(value any) []any {
	if rows, ok := value.([]any); ok {
		return rows
	}
	if object, ok := value.(map[string]any); ok {
		for _, key := range []string{"data", "ocorrencias", "occurrences", "results"} {
			if nested, ok := object[key]; ok {
				if rows := findOccurrenceArray(nested); rows != nil {
					return rows
				}
			}
		}
	}
	return nil
}

func findLastPage(value any) int {
	object, ok := value.(map[string]any)
	if !ok {
		return 0
	}
	for _, key := range []string{"last_page", "lastPage"} {
		if page, err := strconv.Atoi(valueString(object[key])); err == nil && page > 0 {
			return page
		}
	}
	for _, key := range []string{"meta", "pagination", "data", "ocorrencias", "occurrences", "results"} {
		if page := findLastPage(object[key]); page > 0 {
			return page
		}
	}
	return 0
}

func objectValue(value any, keys ...string) any {
	object, ok := value.(map[string]any)
	if !ok {
		return nil
	}
	for _, key := range keys {
		if value, ok := object[key]; ok {
			return value
		}
	}
	return nil
}

func valueString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case json.Number:
		return v.String()
	case float64:
		return fmt.Sprintf("%.0f", v)
	default:
		return ""
	}
}



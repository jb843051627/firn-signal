package model

func CloneStrings(values []string) []string {
	if values == nil {
		return nil
	}
	copyOf := make([]string, len(values))
	copy(copyOf, values)
	return copyOf
}

func CloneLabels(values map[string]string) map[string]string {
	if values == nil {
		return nil
	}
	return values
}

func ClonePoints(values []ProfilePoint) []ProfilePoint {
	if values == nil {
		return nil
	}
	copyOf := make([]ProfilePoint, len(values))
	copy(copyOf, values)
	return copyOf
}

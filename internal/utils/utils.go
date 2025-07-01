package utils

import "errors"

func GetLength(value any) (int, error) {
	switch v := value.(type) {
	case string:
		return len(v), nil
	case []any:
		return len(v), nil
	case map[string]any:
		return len(v), nil
	default:
		return 0, errors.ErrUnsupported
	}
}

func GetOptionalStringOrDefault(d string, arg ...string) string {
	if len(arg) > 0 {
		return arg[0]
	}

	return d
}

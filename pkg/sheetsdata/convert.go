package sheetsdata

import "strconv"

func at(row []interface{}, idx int) interface{} {
	if idx < len(row) {
		return row[idx]
	}
	return nil
}

func toFloat(v interface{}) float64 {
	switch t := v.(type) {
	case float64:
		return t
	case string:
		f, _ := strconv.ParseFloat(t, 64)
		return f
	default:
		return 0
	}
}

func toInt(v interface{}) int {
	return int(toFloat(v))
}

func toStr(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

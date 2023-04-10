package regex

const (
	FILTERS_REGEX = "^([a-zA-Z_]+) (\\S+) ('(?:[^']|\\\\')*?'|[0-9.]+)(?: ((?i)AND|OR) ([a-zA-Z_]+) (\\S+) ('(?:[^']|\\').*?'|[0-9.]+))?$"
)

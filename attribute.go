package htmlselector

import "strings"

// attribute will be returned by parseAttributes
type attribute struct {
	key string
	value string
}

// parseAttributes takes a attribute line for example abc=123,b,c and parses it into attribute structs
func parseAttributes(attributeLine string) []attribute {
	attributeList := strings.Split(attributeLine, ",")
	attributes := make([]attribute, len(attributeList))

	for index, attr := range attributeList {
		splitAttribute := strings.Split(attr, "=")

		if len(splitAttribute) == 2 {
			attributes[index] = attribute{key: splitAttribute[0], value: strings.Trim(splitAttribute[1], "\"")}
		} else {
			attributes[index] = attribute{key: splitAttribute[0]}
		}
	}

	return attributes
}

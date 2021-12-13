package htmlselector

import (
	"fmt"
	"strings"
)

type Selector struct {
	RawQuery string
	Rows     []SelectorRow
}

type SelectorRow struct {
	RawQuery   string
	IsFirst    bool
	NthChild   int
	Type       string
	Attributes map[string]string
}

func NewSelector(query string) (Selector, error) {
	rows := strings.Fields(query)
	var selectorRows []SelectorRow
	var nthChild int

	for index, row := range rows {
		if row == ">" {
			nthChild = 1
			continue
		}

		selectorRow, err := newSelectorRow(row, index)
		if err != nil {
			return Selector{}, err
		}

		if nthChild > 0 {
			selectorRow.NthChild = nthChild
			nthChild = 0
		}
		selectorRows = append(selectorRows, selectorRow)
	}

	return Selector{RawQuery: query, Rows: selectorRows}, nil
}

func newSelectorRow(rowQuery string, index int) (SelectorRow, error) {
	selectorRow := SelectorRow{RawQuery: rowQuery, IsFirst: (index == 0), Attributes: make(map[string]string)}

	var parseCache string
	var firstKeyFound bool
	var parsingType rune

	for _, char := range rowQuery {
		if char == '#' || char == '.' || char == '[' || char == ']' {
			if firstKeyFound == false {
				selectorRow.Type = parseCache
				firstKeyFound = true
			}

			if parsingType == '#' {
				// there can only be one id
				selectorRow.Attributes["id"] = parseCache
			} else if parsingType == '.' {
				// append with space to the class attribute
				if _, ok := selectorRow.Attributes["class"]; ok {
					selectorRow.Attributes["class"] += " " + parseCache
				} else {
					selectorRow.Attributes["class"] = parseCache
				}
			} else if parsingType == '[' && char == ']' {
				attributes := parseAttributes(parseCache)

				for _, attribute := range attributes {
					selectorRow.Attributes[attribute.key] = attribute.value
				}
			}

			parseCache = ""
		}

		// always set the current parsing type
		if char == '#' || char == '.' || char == '[' {
			parsingType = char
		}

		// only add the current char to the parsingCache when we are not dealing with a
		// basic selector char
		if char != '#' && char != '.' && char != '[' && char != ']' {
			parseCache = fmt.Sprintf("%s%c", parseCache, char)
		}
	}

	if firstKeyFound == false {
		selectorRow.Type = parseCache
	}

	// if we were parsing a basic selector and there are no chars left
	// finish parsing the basic selectors
	if parsingType == '#' {
		// there can only be one id
		selectorRow.Attributes["id"] = parseCache
	} else if parsingType == '.' {
		// append with space to the class attribute
		if _, ok := selectorRow.Attributes["class"]; ok {
			selectorRow.Attributes["class"] += " " + parseCache
		} else {
			selectorRow.Attributes["class"] = parseCache
		}
	}

	return selectorRow, nil
}

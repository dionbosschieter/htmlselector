package htmlselector

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

type matcher struct {
	selector   Selector
	node       *html.Node
	ChildIndex int
}

func Match(query string, htmlData []byte) ([]*html.Node, error) {
	reader := bytes.NewReader(htmlData)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	selector, err := NewSelector(query)
	if err != nil {
		return nil, err
	}

	htmlSelector := matcher{selector: selector, node: doc}

	return htmlSelector.execute(), nil
}

func (m *matcher) execute() []*html.Node {
	return m.findNode(m.node)
}

// for selector rows find matches in entire doc
// - for next selector find matches in previous matches
func (m *matcher) findNode(n *html.Node) []*html.Node {
	matches := m.findMatchesInNode(n, m.selector.Rows[0], false)

	for _, selectorRow := range m.selector.Rows[1:] {
		var subMatches []*html.Node
		for _, match := range matches {
			m.ChildIndex = 0
			subMatches = append(subMatches, m.findMatchesInNode(match, selectorRow, true)...)
		}
		matches = subMatches
	}

	return matches
}

func (m *matcher) matchesSelectorRow(selector SelectorRow, n *html.Node) bool {
	if selector.Type != "" && n.Data != selector.Type {
		return false
	}

	// check if all selector attributes exist and match
	for attrKey, attrValue := range selector.Attributes {
		matches := false

		for _, elemAttr := range n.Attr {
			if elemAttr.Key == attrKey && isInList(strings.Fields(attrValue), strings.Fields(elemAttr.Val)) {
				matches = true
				break
			}
		}

		if matches == false {
			return false
		}
	}

	return true
}

func isInList(needles []string, list []string) bool {
	for _, needle := range needles {
		matches := false
		for _, listItem := range list {
			if listItem == needle {
				matches = true
				break
			}
		}
		if matches == false {
			return false
		}
	}

	return true
}

func (m *matcher) findMatchesInNode(n *html.Node, row SelectorRow, skipParent bool) (matches []*html.Node) {
	if row.NthChild != 0 && m.ChildIndex > row.NthChild {
		return
	}
	if skipParent != true && n.Type == html.ElementNode && m.matchesSelectorRow(row, n) {
		matches = append(matches, n)
	}

	// todo: replace with iteration
	// traverse through child nodes
	m.ChildIndex++
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		matches = append(matches, m.findMatchesInNode(c, row, false)...)
	}
	m.ChildIndex--

	return
}

package htmlselector

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestSingleMatch(t *testing.T) {
	nodes, err := Match("div#test", []byte("<html><body><div id=\"test\"></div></body></html>"))

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.Len(t, nodes, 1)
}

func TestSubSelectorMatch(t *testing.T) {
	nodes, err := Match("div#test a.test123", []byte("<html><body><div id=\"test\"><a class=\"test123\"></a></div></body></html>"))

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.Len(t, nodes, 1)
}

func TestSubSelectorMatchMultiple(t *testing.T) {
	nodes, err := Match("div#test a", []byte("<html><body><div id=\"test\"><a></a><a class=\"test123\"></a></div></body></html>"))

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.Len(t, nodes, 2)
}

func TestSubSelectorReturnsNothingMatch(t *testing.T) {
	nodes, err := Match("div#test a.test1234", []byte("<html><body><div id=\"test\"><a class=\"test123\"></a></div></body></html>"))

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.Len(t, nodes, 0)
}

func TestEmptyMatch(t *testing.T) {
	nodes, err := Match("div#test a.link", []byte("<html></html>"))

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.Equal(t, len(nodes), 0)
}

func TestDirectChildSelectorMatch(t *testing.T) {
	nodes, err := Match("div#test > a", []byte("<html><body><div id=\"test\"><a>link1</a><a>link2</a><p><a>link3</a></p></div></body></html>"))

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.Len(t, nodes, 2)
}

func TestMultipleDirectChildSelectorMatch(t *testing.T) {
	nodes, err := Match("div#test > div > p > a", []byte("<html><body><div id=\"test\"><div><p><a>link1</a></p></div><a>link2</a><p><a>link3</a></p></div></body></html>"))

	require.NoError(t, err, "Error was thrown on a non error query")
	assert.Len(t, nodes, 1)
}

func TestCombinationDirectNonDirectChildSelectorMatch(t *testing.T) {
	nodes, err := Match("div#test > div a", []byte("<html><body><div id=\"test\"><div><p><a>link1</a></p></div><a>link2</a><p><a>link3</a></p></div></body></html>"))

	require.NoError(t, err, "No error was thrown on a non error query")
	require.Len(t, nodes, 1)
	assert.Equal(t, "link1", nodes[0].FirstChild.Data)
}

func TestFullWebsiteFindListHrefs(t *testing.T) {
	file, err := os.Open("./fixtures/test.html")
	require.NoError(t, err)
	buffer, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	nodes, err := Match("tr.lista2 > td.lista > a[title=\"\"]", buffer)
	require.NoError(t, err, "Error was thrown on a non error query")
	require.Len(t, nodes, 6)

	assert.Equal(t, nodes[0].Attr[2].Val, "/torrent/exq6sn7")
}

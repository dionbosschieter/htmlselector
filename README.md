# HtmlSelector
Documentation of dependency used: https://pkg.go.dev/golang.org/x/net/html  
This package tries to provide you with a html selector, without using any external dependencies other than the golang net/html package.  

```golang
nodes, err := htmlselector.Match("div#test a.link", []byte("<html></html>"))
nodes, err := htmlselector.MatchNode("div#test a.link", *node) (*Node, err)

for node, _ := range nodes {
	fmt.Println(node)
}
```

## Pre Compiling
If you don't want the css selector parser to parse the same kind of query everytime you want to match a piece of html  
you can choose to precompile it.  
This would give you a selector which you can either use to match against raw html or parsed html  

```golang
selector, err := htmlselector.NewSelector("div#test a.link")
matchedNodes, err := selector.Match([]byte("<html></html>"))
matchedNodes := selector.MatchNode(*node))
```

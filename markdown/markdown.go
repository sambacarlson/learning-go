package markdown

import (
	"fmt"

	"github.com/russross/blackfriday/v2"
)

// takes markdown as string, and nodetype to check. returns string array of requested prop.
func ProcessMd(mkdwn string, nodetype string) ([]string, error) {
	if err := validateNode(nodetype); err != nil {
		return nil, err
	}
	// create abstract syntax tree
	ast := blackfriday.New().Parse([]byte(mkdwn))

	// get possible values for nodetype
	//nolint:varnamelen
	var ty []string
	switch nodetype {
	case "Document":
		ty = getNode(ast, blackfriday.Document)
	case "Heading":
		ty = getNode(ast, blackfriday.Heading)
	case "Paragraph":
		ty = getNode(ast, blackfriday.Paragraph)
	case "Blockquote":
		ty = getNode(ast, blackfriday.BlockQuote)
	case "List":
		ty = getNode(ast, blackfriday.List)
	case "Item":
		ty = getNode(ast, blackfriday.Item)
	case "Text":
		ty = getNode(ast, blackfriday.Text)
	case "Strong":
		ty = getNode(ast, blackfriday.Strong)
	case "Emph":
		ty = getNode(ast, blackfriday.Emph)
	default:
		//nolint:goerr113
		return nil, fmt.Errorf("node type %v currently unsupported", nodetype)
	}
	return ty, nil
}

func getNode(node *blackfriday.Node, tpe blackfriday.NodeType) []string {
	var nArr []string
	// traverse tree and return value of nodes whose type equals tpe
	node.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if entering && node.Type == tpe {
			if string(node.FirstChild.Literal) != "" {
				nArr = append(nArr, string(node.FirstChild.Literal))
			}
		}
		return blackfriday.GoToNext
	})
	return nArr
}

func validateNode(node string) error {
	nodetypes := []string{"Document", "Heading", "Paragraph", "Blockquote", "List", "Item", "Text", "Emph", "Strong", "Code", "Link", "Image"}
	for _, notetype := range nodetypes {
		if notetype == node {
			return nil
		}
	}
	//nolint:goerr113
	return fmt.Errorf("invalid node type: %v", node)
}
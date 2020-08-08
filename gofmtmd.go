package gofmtmd

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

// FmtGoCodeInMarkdown formats go code in Markdown.
// returns error if code has syntax error.
func FmtGoCodeInMarkdown(md []byte) ([]byte, error) {
	var err error
	n := blackfriday.New(blackfriday.WithExtensions(blackfriday.FencedCode)).Parse(md)
	n.Walk(genFmtWalker(&md, &err))
	if err != nil {
		return nil, fmt.Errorf("[gofmtmd] %w", err)
	}
	return md, nil
}

func isGoCodeBlock(node *blackfriday.Node) bool {
	return node.Type == blackfriday.CodeBlock && string(node.CodeBlockData.Info) == "go"
}

func genFmtWalker(md *[]byte, fmterr *error) blackfriday.NodeVisitor {
	return func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if isGoCodeBlock(node) {
			fmted, err := format.Source(node.Literal)
			if err != nil {
				line := strings.Split(err.Error(), ":")[0]
				i, cvtErr := strconv.Atoi(line)
				if cvtErr != nil {
					*fmterr = fmt.Errorf("failed to parse error msg: %w", cvtErr)
				} else {
					code := bytes.Split(node.Literal, []byte("\n"))
					*fmterr = fmt.Errorf("line='%s', msg='%w'", string(code[i+1]), err)
				}
				return blackfriday.Terminate
			}
			*md = bytes.ReplaceAll(*md, bytes.TrimRight(node.Literal, "\n"), bytes.TrimRight(fmted, "\n"))
		}
		return blackfriday.GoToNext
	}
}

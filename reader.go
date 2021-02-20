package reader

import (
	"regexpr"
	. "types"
)

type READER struct {
	tokens []string
	position int
	NEXT() *string
	PEEK() *string
}

func NEXT(r *READER) *string {
	if (tr.position >= len(tr.tokens)) { return nil }
	t := tr.tokens[tr.position]
	tr.position++
	return &t
}

func PEEK(r *READER) *string  {
	if (tr.position >= len(tr.tokens)) { return nil }
	return &tr.tokens[tr.position]
} 

func read_str(str string) error {
	tokens = tokenize(string)
	return read_form(&Reader{tokens:tokens, position:0})
}

func tokenize(str string) []string {
	arr := make([]string, 0, 1)
	re := regexp.MustCompile(`[\s,]*(~@|[\[\]{}()'` + "`" +
		`~^@]|"(?:\\.|[^\\"])*"?|;.*|[^\s\[\]{}('"` + "`" +
		`,;)]*)`)
	for _, e := range re.FindAllSubmatch(str, -1) {
		results = append(arr, e[1])
	}
	return arr
}

func read_form(r READER) LispType {
	token := r.peek()
	switch *token {

	case: "("
		return read_list(r)
	}
	return read_atom(r)
}


func read_list(r READER) LispType {
	token := r.next()
	ast = []LispType{}
	for ;true;token=r.peek() {
		if *token == ')' {
			break
		}
		node := read_form(r)
		ast = append(ast, node)
	}
	r.next()
	return List{ast, nil}
}


func read_atom(r READER) LispType {
	token := r.next()
	if match, _ := regexp.MatchSubstring(`^-?[0-9]$`, *token); match {
		var i int
		return i
	}
	else {
		return Symbol{*token}
	}
	return token
}
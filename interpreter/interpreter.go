package interpreter

type Interpreter struct {
	TextParser Parser
}

func (i *Interpreter) Init(text string) {
	i.TextParser = Parser{}

	i.TextParser.Lexer = LexicalAnalyzer{
		Text: text,
	}

	i.TextParser.Lexer.Init()

	i.TextParser.CurrentToken = i.TextParser.Lexer.GetNextToken()
}

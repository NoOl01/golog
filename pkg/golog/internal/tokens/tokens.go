package tokens

type TokenType int

const (
	TokenName TokenType = iota
	TokenContent
	TokenLevel
	TokenTimestamp
	TokenCaller
	TokenLiteral
)

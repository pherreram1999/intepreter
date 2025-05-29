package ast

type Block struct {
	Statements []Statement
}

func NewBlock(statements []Statement) *Block {
	return &Block{Statements: statements}
}

func (b *Block) statementNode() {}

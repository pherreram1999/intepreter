package ast

import "pahm/intepreter/compiler/environment"

type Block struct {
	Statements []Statement
}

func NewBlock(statements []Statement) *Block {
	return &Block{Statements: statements}
}

func (b *Block) statementNode() {}

func (b *Block) Env(env *environment.Environment) (any, error) {
	// Crear nuevo ambiente para el bloque
	blockEnv := environment.NewWithEnclosing(env)

	for _, stmt := range b.Statements {
		_, err := stmt.Env(blockEnv)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

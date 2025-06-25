package ast

import (
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
	"pahm/intepreter/compiler/validators"
)

type ForStatement struct {
	Initializer Statement  //Inicia el contador
	Condition   Expression //Condicion de paro
	Increment   Expression
	Body        Statement //Bloque/ statement
}

func NewForStatement(initializer Statement, condition Expression, increment Expression, body Statement) *ForStatement {
	return &ForStatement{initializer, condition, increment, body}
}
func (p *ForStatement) statementNode() {}
func (p *ForStatement) GetToken() *token.Token {
	return nil
}

func (f *ForStatement) Env(env *environment.Environment) (any, error) {
	// Crear nuevo ambiente para el for
	forEnv := environment.NewWithEnclosing(env)

	// Ejecutar inicializador si existe
	if f.Initializer != nil {
		_, err := f.Initializer.Env(forEnv)
		if err != nil {
			return nil, err
		}
	}

	// Loop principal
	for {
		// Evaluar condición si existe
		if f.Condition != nil {
			condition, err := f.Condition.Env(forEnv)
			if err != nil {
				return nil, err
			}
			if !validators.IsTruthy(condition) {
				break
			}
		}

		// Ejecutar cuerpo
		_, err := f.Body.Env(forEnv)
		if err != nil {
			return nil, err
		}

		// Ejecutar incremento si existe
		if f.Increment != nil {
			_, err = f.Increment.Env(forEnv)
			if err != nil {
				return nil, err
			}
		}
	}

	return nil, nil
}

package environment

import (
	"fmt"
	"pahm/intepreter/compiler/token"
)

type Environment struct {
	enclosing *Environment
	values    map[string]any
}

func New() *Environment {
	return &Environment{
		enclosing: nil,
		values:    make(map[string]any),
	}
}

func NewWithEnclosing(env *Environment) *Environment {
	return &Environment{
		enclosing: env,
		values:    make(map[string]any),
	}
}

// Assing agrega el valor de variables definidas
func (env *Environment) Assing(token *token.Token, value any) error {
	if _, ok := env.values[token.Lexema]; ok {
		// si encuentra en ambito actual
		env.values[token.Lexema] = value
		return nil
	}

	// si no existe y no tenemos cierre
	if env.enclosing == nil {
		return fmt.Errorf("variable indefinida '%s' en línea %d",
			token.Lexema, token.Linea)
	}

	return env.enclosing.Assing(token, value) // recorremos hacia arriba en busca del ambito
}

func (env *Environment) Define(key string, value any) {
	env.values[key] = value
}

func (env *Environment) Get(token *token.Token) (any, error) {
	if _, ok := env.values[token.Lexema]; ok {
		return env.values[token.Lexema], nil
	}
	if env.enclosing == nil {
		return nil, fmt.Errorf("variable indefinida '%s' en línea %d",
			token.Lexema, token.Linea)
	}

	return env.enclosing.Get(token) // buscamos hacia arriba
}

func (env *Environment) String() string {
	res := fmt.Sprintf("%v", env.values)
	if env.enclosing != nil {
		res += fmt.Sprintf(" -> %s", env.enclosing)
	}
	return res
}

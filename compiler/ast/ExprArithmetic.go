package ast

import (
	"errors"
	"fmt"
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
	"pahm/intepreter/compiler/validators"
)

type Arithmetic struct {
	left     Expression
	operator *token.Token
	right    Expression
}

func NewArithmetic(left Expression, operator *token.Token, right Expression) *Arithmetic {
	return &Arithmetic{left, operator, right}
}

func (a *Arithmetic) expressionNode()        {}
func (a *Arithmetic) GetToken() *token.Token { return a.operator }

func (a *Arithmetic) Env(env *environment.Environment) (any, error) {
	left, err := a.left.Env(env)
	if err != nil {
		return nil, err
	}
	right, err := a.right.Env(env)
	if err != nil {
		return nil, err
	}

	switch a.operator.Tipo {
	case token.Plus:
		return a.EvalPlus(left, right)
	case token.Minus:
		return a.EvalMinus(left, right)
	case token.Star:
		return a.EvalProduct(left, right)
	case token.Slash:
		return a.EvalDivision(left, right)
	default:
		return nil, fmt.Errorf("operador aritmético desconocido: %s en línea %d",
			a.operator.Lexema, a.operator.Linea)
	}
}

func (a *Arithmetic) EvalPlus(left, rigth any) (any, error) {
	if validators.IsNumberV(left) && validators.IsNumberV(rigth) {
		return left.(float64) + rigth.(float64), nil // damos por convención la conversion a flotante
	}

	if validators.IsString(left) && validators.IsString(rigth) {
		return left.(string) + rigth.(string), nil // retornamos la concatenacion
	}

	return nil, errors.New("operandos incompatibles para +")
}

func (a *Arithmetic) EvalMinus(left, rigth any) (any, error) {
	if validators.IsNumberV(left) && validators.IsNumberV(rigth) {
		return left.(float64) - rigth.(float64), nil
	}
	return nil, errors.New("operandos incompatibles para -")
}

func (a *Arithmetic) EvalProduct(left, rigth any) (any, error) {
	if validators.IsNumberV(left) && validators.IsNumberV(rigth) {
		return left.(float64) * rigth.(float64), nil
	}
	return nil, errors.New("operandos incompatibles para *")
}

func (a *Arithmetic) EvalDivision(left, rigth any) (any, error) {
	if validators.IsNumberV(left) && validators.IsNumberV(rigth) {
		b := rigth.(float64)
		if b == 0 {
			return nil, errors.New("no es posible dividir entre cero, aun no aplicamos limites")
		}
		return left.(float64) / b, nil
	}
	return nil, errors.New("operandos incompatibles para /")
}

package ast

import (
	"bufio"
	"fmt"
	"os"
	"pahm/intepreter/compiler/environment"
	"pahm/intepreter/compiler/token"
	"strconv"
	"strings"
)

type InputExpression struct {
	Name  *token.Token
	Value Expression
}

func NewInputExpression(name *token.Token, value Expression) *InputExpression {
	return &InputExpression{name, value}
}
func (ie *InputExpression) expressionNode() {}
func (ie *InputExpression) GetToken() *token.Token {
	return ie.Name
}
func (ie *InputExpression) Env(env *environment.Environment) (any, error) {
	// Evaluar el prompt si existe
	var prompt string
	if ie.Value != nil {
		promptValue, err := ie.Value.Env(env)
		if err != nil {
			return nil, err
		}
		prompt = fmt.Sprintf("%v", promptValue)
	}

	if prompt != "" {
		fmt.Print(prompt)
	}

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("error leyendo entrada: %v", err)
	}

	inputStr := strings.TrimSpace(string(input))

	if num, err := strconv.ParseFloat(inputStr, 64); err == nil {
		return num, nil
	}

	return inputStr, nil
}

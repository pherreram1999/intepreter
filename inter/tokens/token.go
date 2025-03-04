package tokens

type Token struct {
	Tipo    TipoToken
	Lexema  string
	Literal interface{}
	Linea   uint
}

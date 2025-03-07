package token

const (
	Minus        TipoToken = iota
	Plus         TipoToken = iota
	Slash        TipoToken = iota //Division
	Star         TipoToken = iota //Multiplicacion
	Bang         TipoToken = iota //Diferente
	BangEqual    TipoToken = iota
	Equal        TipoToken = iota
	EqualEqual   TipoToken = iota
	Greater      TipoToken = iota
	GreaterEqual TipoToken = iota
	Less         TipoToken = iota
	LessEqual    TipoToken = iota
)

package tokens

type TipoToken uint

const (
	// reserved keys
	Else  TipoToken = iota // 0
	Fun   TipoToken = iota // 1
	Print TipoToken = iota
	Var   TipoToken = iota
	False TipoToken = iota
)

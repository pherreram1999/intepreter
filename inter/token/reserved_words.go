package token

type TipoToken uint

const (
	// reserved keys
	Else   TipoToken = iota // 0
	Fun    TipoToken = iota // 1
	Print  TipoToken = iota
	Var    TipoToken = iota
	False  TipoToken = iota
	If     TipoToken = iota
	Return TipoToken = iota
	While  TipoToken = iota
	For    TipoToken = iota
	Null   TipoToken = iota
	True   TipoToken = iota
)

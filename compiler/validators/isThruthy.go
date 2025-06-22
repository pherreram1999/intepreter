package validators

func IsTruthy(value any) bool {
	if value == nil {
		return false // cualquier cosa es falsa, pero sera negada
	}

	if b, ok := value.(bool); ok {
		return b // niega el booleano como tal
	}

	return true // por defecto es falso
}

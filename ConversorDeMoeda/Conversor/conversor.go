package conversor

func LogDoConversor(valor float64, moedaDeOrigem, moedaDeDestino string) float64 {
	moedas := map[string]map[string]float64{
		"BRL": {
			"USD": 0.19,
			"EUR": 0.16,
		},
		"USD": {
			"BRL": 5.36,
			"EUR": 0.86,
		},
		"EUR": {
			"BRL": 6.22,
			"USD": 1.66,
		},
	}

	return valor * moedas[moedaDeOrigem][moedaDeDestino]
}

package cruder

// uri меняет endpoint для запроса
type uri struct {
	uri string
}

// Option имплементация интерфеса types.ICruderOption
func (o *uri) Option(optionSetIn iCruderOptionHelper) {
	optionSetIn.SetUri(o.uri)
}

// Uri меняет endpoint для запроса
func Uri(uriIn string) (out ICruderOption) {
	return &uri{uri: uriIn}
}

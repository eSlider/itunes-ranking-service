package itunes

type Country string

const (
	USA         Country = "us"
	Deutschland Country = "de"
	Spanien     Country = "es"
	Frankreich  Country = "fr"
	Italien     Country = "it"
)

var Countries = []Country{USA, Deutschland, Spanien, Frankreich, Italien}

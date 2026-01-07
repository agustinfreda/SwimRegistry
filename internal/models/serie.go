package models

type Serie struct {
	Cantidad int `json:"cantidad"`
	Ejercicio []Ejercicio json:"ejercicio"`
}

package models

type Ejercicio struct {
	Repeticiones int      `json:"repeticiones"`
	Distancia    int      `json:"distancia"`
	Estilo       string   `json:"estilo"`
	Ejercicio    string   `json:"ejercicio"`
	Tiempo		 	 *string 	`json:"tiempo"`
	Material     []string `json:"material"`
}

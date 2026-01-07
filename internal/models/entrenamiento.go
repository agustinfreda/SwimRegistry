package models

type Entrenamiento struct {
	ID             string  `json:"id"`
	Fecha          string  `json:"fecha"`
	Comentario     string  `json:"comentario"`
	DuracionTotal  *string `json:"duracion_total,omitempty"`
	Series         []Serie `json:"series"`
}

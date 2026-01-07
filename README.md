# SwimRegistry


# Invariantes de Ejercicio

## Repeticiones > 0 -> El valor de Repeticiones debe ser mayor a 0.
## Distancia > 0 && % 25 -> El valor de Distancia debe ser mayor a 0 y multiplo de 25.
## Estilo = ['Crol', 'Mariposa', 'Espalda', 'Pecho'] -> El estilo debe ser Crol, Mariposa, Espalda o Pecho
## Material = ['Nada', 'Manopla', 'Aleta', 'Pullbouy', 'Tabla', 'Drill'] -> El material debe ser Nada, Manopla, Aleta, Pullbouy, Tabla o Drill
## Ejercicio = ['Patada', 'Libre', 'Drill'] -> Ejercicio debe ser Patada, Libre o Drill
# Invarientes de Entrenamiento

## Fecha formato fecha -> La fecha debe tener formato fecha (dd/mm/aaaa)
## Comentario personal de como senti el entrenamiento, debe ser string maximo 50 caracteres -> Debe contener máximo 50 caracteres.
## Duracion total en minutos y luego se convierte a horas, por ejemplo 1h35m, 1h46m, si pasa de 60min se convierte en h -> La duracion total debe ser mayor a 0 y minutos.


# Invariantes de Serie

## Cantidad > 0 -> Cantidad debe ser un valor mayor a 0.
## len(Ejercicios) > 0 -> Una serie debe tener al menos un ejercicio.



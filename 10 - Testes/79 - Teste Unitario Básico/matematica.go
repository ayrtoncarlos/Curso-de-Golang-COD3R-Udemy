package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular a média das notas passados.
func Media(notas ...float64) float64 {
	total := 0.0
	for _, nota := range notas {
		total += nota
	}
	media := total / float64(len(notas))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}

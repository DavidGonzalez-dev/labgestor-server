package verificationCodes

import (
	"math/rand"
	"fmt"
)

func GenerarCodigoVerificacion() string {
	codigo := fmt.Sprintf("%06d", rand.Intn(100000))
	return codigo
}
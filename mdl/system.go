package mdl

import (
	"crypto/sha256"
	"encoding/hex"
)

var (
	URL              = "http://localhost/devel/api/"
	MT_ACTIVE_COMMIT = false
	MT_COMMIT        = "motfa-api"

	FNX_GetAPI = "_SYS_ListarAPI"
)

// SetHash Asignar un hash de 256 para definir una cadena
func SetHash(code []byte) string {
	h := sha256.New()
	h.Write(code)
	return hex.EncodeToString(h.Sum(nil))
}

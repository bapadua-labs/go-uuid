package gouuid

import (
	"crypto/rand"
	"fmt"
)

type UUID struct {
	UUID string
}

// NewUUID cria um UUID v4 e retorna o struct.
func NewUUID() UUID {
	return UUID{UUID: NewUUIDv4()}
}

// NewUUIDv4 gera um UUID v4 (RFC 4122) e retorna como string.
func NewUUIDv4() string {
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		// Em um ambiente real, o rand.Read raramente falha, a menos que
		// o sistema operacional fique sem entropia.
		panic(fmt.Errorf("falha ao gerar bytes aleatórios: %w", err))
	}

	// Define a versão do UUID para 4 no byte 6
	bytes[6] = (bytes[6] & 0x0f) | 0x40

	// Define a variante do UUID para a RFC 4122 no byte 8
	bytes[8] = (bytes[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
}

func (u UUID) String() string {
	return u.UUID
}

package gouuid

import (
	"regexp"
	"testing"
)

var uuidV4Pattern = regexp.MustCompile(
	`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`,
)

func TestNewUUID_Format(t *testing.T) {
	u := NewUUID()

	if u.UUID == "" {
		t.Fatal("UUID gerado está vazio")
	}

	if !uuidV4Pattern.MatchString(u.UUID) {
		t.Errorf("UUID inválido: %q (esperado formato RFC 4122 versão 4)", u.UUID)
	}
}

func TestNewUUID_VersionAndVariant(t *testing.T) {
	u := NewUUID()
	s := u.UUID

	if s[14] != '4' {
		t.Errorf("versão esperada 4, obtida %c em %q", s[14], s)
	}

	switch s[19] {
	case '8', '9', 'a', 'b':
		// ok
	default:
		t.Errorf("variante RFC 4122 inválida: %c em %q", s[19], s)
	}
}

func TestNewUUID_Uniqueness(t *testing.T) {
	const n = 1000
	seen := make(map[string]struct{}, n)

	for i := 0; i < n; i++ {
		u := NewUUID()
		if _, exists := seen[u.UUID]; exists {
			t.Fatalf("UUID duplicado gerado: %q", u.UUID)
		}
		seen[u.UUID] = struct{}{}
	}
}

func TestUUID_String_ReturnsStoredValue(t *testing.T) {
	u := NewUUID()

	if u.String() != u.UUID {
		t.Errorf("String() = %q, esperado %q", u.String(), u.UUID)
	}

	if !uuidV4Pattern.MatchString(u.String()) {
		t.Errorf("String() retornou UUID inválido: %q", u.String())
	}
}

func TestNewUUIDv4_Format(t *testing.T) {
	s := NewUUIDv4()

	if s == "" {
		t.Fatal("UUID string gerado está vazio")
	}

	if !uuidV4Pattern.MatchString(s) {
		t.Errorf("UUID string inválido: %q (esperado formato RFC 4122 versão 4)", s)
	}
}

func TestNewUUIDv4_VersionAndVariant(t *testing.T) {
	s := NewUUIDv4()

	if s[14] != '4' {
		t.Errorf("versão esperada 4, obtida %c em %q", s[14], s)
	}

	switch s[19] {
	case '8', '9', 'a', 'b':
		// ok
	default:
		t.Errorf("variante RFC 4122 inválida: %c em %q", s[19], s)
	}
}

func TestNewUUIDv4_Uniqueness(t *testing.T) {
	const n = 1000
	seen := make(map[string]struct{}, n)

	for i := 0; i < n; i++ {
		s := NewUUIDv4()
		if _, exists := seen[s]; exists {
			t.Fatalf("UUID string duplicado gerado: %q", s)
		}
		seen[s] = struct{}{}
	}
}

func TestNewUUIDv4_Length(t *testing.T) {
	s := NewUUIDv4()

	// 8-4-4-4-12 + 4 hífens = 36 caracteres
	if len(s) != 36 {
		t.Errorf("tamanho esperado 36, obtido %d em %q", len(s), s)
	}
}

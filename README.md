# go-uuid

Biblioteca leve em Go para geração de **UUID versão 4** (aleatório), conforme a [RFC 4122](https://datatracker.ietf.org/doc/html/rfc4122) (atualizada pela [RFC 9562](https://www.rfc-editor.org/rfc/rfc9562)).

```go
import gouuid "github.com/bapadua/go-uuid"

id := gouuid.NewUUIDv4()
// exemplo: "550e8400-e29b-41d4-a716-446655440000"
```

---

## O que é um UUID?

**UUID** (*Universally Unique Identifier*) é um identificador de **128 bits** (16 bytes) projetado para ser único entre sistemas, sem precisar de um servidor central de IDs.

Na forma textual canônica, um UUID tem **36 caracteres**:

```text
xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
│         │    │    │    │
│         │    │    │    └─ 12 hex (48 bits)
│         │    │    └────── 3 hex + variante (16 bits)
│         │    └─────────── versão M + 3 hex (16 bits)
│         └──────────────── 4 hex (16 bits)
└────────────────────────── 8 hex (32 bits)
```

| Parte | Tamanho | Papel |
|-------|---------|--------|
| `xxxxxxxx` | 8 hex | Bits aleatórios |
| `xxxx` | 4 hex | Bits aleatórios |
| `Mxxx` | 4 hex | **M** = versão (em v4 vale `4`) |
| `Nxxx` | 4 hex | **N** = variante RFC (`8`, `9`, `a` ou `b`) |
| `xxxxxxxxxxxx` | 12 hex | Bits aleatórios |

Exemplo válido de UUID v4:

```text
f47ac10b-58cc-4372-a567-0e02b2c3d479
              ↑    ↑
           versão  variante
              4      a
```

---

## Por que UUID versão 4?

Existem várias versões de UUID. As mais comuns:

| Versão | Base | Quando usar |
|--------|------|-------------|
| **1** | Tempo + MAC | Ordenação temporal, mas vaza info de rede/hora |
| **3** | Hash MD5 de um nome | ID determinístico a partir de um nome |
| **4** | Aleatório | IDs opacos, tokens, chaves primárias sem sequência |
| **5** | Hash SHA-1 de um nome | Como a v3, com hash mais forte |
| **7** (RFC 9562) | Tempo + aleatório | Boa para índices em banco (ordenável no tempo) |

A **versão 4** é a mais usada no dia a dia porque:

- Não depende de relógio, hostname ou MAC
- Não revela ordem de criação
- É simples de gerar em qualquer ambiente
- Tem espaço enorme: **122 bits** aleatórios (128 − 4 de versão − 2 de variante)

A chance de colisão é astronômica: mesmo gerando bilhões de UUIDs, a probabilidade de duplicata continua desprezível para a maioria das aplicações.

---

## Como a RFC 4122 define o UUID v4

O algoritmo é:

1. Gerar **16 bytes** aleatórios (criptograficamente seguros).
2. Forçar a **versão 4** nos 4 bits mais significativos do **byte 6**:
   ```text
   bytes[6] = (bytes[6] & 0x0F) | 0x40
   ```
   Isso deixa o byte no formato `0100xxxx`.
3. Forçar a **variante RFC 4122** nos 2 bits mais significativos do **byte 8**:
   ```text
   bytes[8] = (bytes[8] & 0x3F) | 0x80
   ```
   Isso deixa o byte no formato `10xxxxxx` (dígito hex `8`, `9`, `a` ou `b`).
4. Formatar como hexadecimal com hífens: `8-4-4-4-12`.

```text
Byte:  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
                    │     │
                    │     └─ variante: 10xxxxxx
                    └─────── versão:   0100xxxx
```

Esta biblioteca segue exatamente esse algoritmo, usando `crypto/rand` do Go.

---

## Instalação

```bash
go get github.com/bapadua/go-uuid
```

Requisito: **Go 1.26+** (conforme o `go.mod` do projeto).

---

## Uso

### String direto

```go
package main

import (
	"fmt"

	gouuid "github.com/bapadua/go-uuid"
)

func main() {
	id := gouuid.NewUUIDv4()
	fmt.Println(id)
}
```

### Struct `UUID`

```go
u := gouuid.NewUUID()

fmt.Println(u.UUID)    // campo string
fmt.Println(u.String()) // mesmo valor (estilo toString)
fmt.Println(u)          // fmt usa String() automaticamente
```

---

## API

| Função / método | Retorno | Descrição |
|-----------------|---------|-----------|
| `NewUUIDv4()` | `string` | Gera um UUID v4 e devolve a representação textual |
| `NewUUID()` | `UUID` | Gera um UUID v4 encapsulado no struct |
| `(UUID).String()` | `string` | Devolve o valor já gerado (não gera outro) |

```go
type UUID struct {
	UUID string
}
```

---

## Exemplos de saída

Toda chamada produz um valor **novo** e **único**:

```text
3f2a9c1e-8b4d-4e71-9c2a-1f6b8d0e3a5c
a1b2c3d4-e5f6-4789-8abc-def012345678
9e8d7c6b-5a49-4f3e-b2a1-0987654321fe
```

Garantias visíveis no texto:

- 3º grupo **sempre** começa com `4` (versão)
- 4º grupo **sempre** começa com `8`, `9`, `a` ou `b` (variante)
- Comprimento **sempre** 36 caracteres

---

## Testes

```bash
go test ./...
```

A suíte valida formato RFC 4122 v4, bits de versão/variante, tamanho e unicidade em milhares de gerações.

---

## Quando (não) usar UUID v4

**Bom para:**

- Chaves primárias / IDs de recurso em APIs
- Correlation IDs em logs e tracing
- Tokens opacos sem sequência previsível
- Sistemas distribuídos sem coordenação central de IDs

**Considere alternativas se:**

- Precisar de IDs **ordenáveis no tempo** para índices de banco → veja UUID v7 (RFC 9562)
- Precisar de ID **determinístico** a partir de um nome → UUID v3/v5
- Quiser IDs curtos e legíveis → ULID, KSUID, nanoid, etc.

---

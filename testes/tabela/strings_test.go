package strings

import (
	"strings"
	"testing"
)

const msgErro = "%s (parte: %s) - índices: esperando (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Texto bala", "abc", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Ok Ratinho", "r", 2},
	}
	for _, teste := range testes {
		t.Logf("Show: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgErro, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}

package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // executar em paralelo
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em amd64")
	}
	t.Fail()
}

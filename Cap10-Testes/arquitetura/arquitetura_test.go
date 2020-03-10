package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // Executa o teste de forma concorrente/paralela
	if runtime.GOARCH != "amd64" {
		t.Skip("Apenas funciona em arquitetura amd64") //pula o teste nesse caso
	}
	arq := runtime.GOARCH
	t.Logf("Log: %v", arq)
	t.Fail()
}

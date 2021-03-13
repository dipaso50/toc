package domain

import "testing"

func Test_createFile(t *testing.T) {
	origen := "origen"
	name := "name"
	destino := "destino"

	n, err := NewFileOrganizer(origen, destino, name)

	if err != nil {
		t.Errorf("Error creando fichero %v", err)
	}

	if n.Name != name || n.Destination != destino || n.Origin != origen {
		t.Error("Objeto creado no es el esperado")
	}
}

package domain

import "fmt"

type FileOrganizer struct {
	Origin      string
	Destination string
	Name        string
}

func NewFileOrganizer(origin, destination, name string) (FileOrganizer, error) {
	if len(destination) == 0 || len(origin) == 0 || len(name) == 0 {
		return FileOrganizer{}, fmt.Errorf("Error creando Objecto File, hace falta un origen, un destino y un nombre ")
	}

	return FileOrganizer{Destination: destination, Name: name, Origin: origin}, nil
}

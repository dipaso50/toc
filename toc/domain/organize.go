package domain

type Organize interface {
	ExistsDirectory(dirname string) (bool, error)
	CreateDirectory(dirname string) error
	MoveFile(file FileOrganizer) error
}

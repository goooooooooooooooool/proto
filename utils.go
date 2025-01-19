package proto

// Создает роль из строки
func NewRole(role string) Role {
	switch role {
	case "READER":
		return Role_READER
	case "AUTHOR":
		return Role_AUTHOR
	}
	return Role_READER
}

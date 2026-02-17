package database

import "github.com/Pyth0nHater/link-shorter/internal/models"

func InitMemory() *models.MemoryMap {
	Storage:=make(map[string]string)
	ReverseStorage:=make(map[string]string)
	return &models.MemoryMap{Storage: Storage, ReverseStorage:ReverseStorage}
}
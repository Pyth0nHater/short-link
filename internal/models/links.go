package models

import "sync"

type MainLink struct {
	MainLink string `json:"main_link"`
}

type ShortLink struct {
	ShortLink string `json:"short_link"`
}

type CreateLink struct {
	MainLink  string `json:"main_link"`
	ShortLink string `json:"short_link"`
}

type GetLinkFilter struct {
	MainLink  *string
	ShortLink *string
}

type GetLink struct {
	MainLink  string
	ShortLink string
}

type MemoryMap struct{
	Storage map[string]string
	ReverseStorage map[string]string
	Mutex sync.RWMutex
}
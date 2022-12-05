package models

type HavariaCache map[string]HavariaCacheEntry

func (tt HavariaCache) HasEntry(c string) bool {
	if val, ok := tt[c]; !ok {
		_ = val
		return false
	}

	return true
}

type HavariaCacheEntry struct {
	Time int `json:"time"`
}

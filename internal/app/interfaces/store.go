package interfaces

type URLStore interface {
	SaveURL(uuid, url string)
	GetURL(uuid string) (string, bool)
	IsStoreEmpty() bool
}

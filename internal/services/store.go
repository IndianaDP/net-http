package services

var store = map[string]string{}

func SaveURL(uuid, url string) {
	store[uuid] = url
}

func GetURL(uuid string) (string, bool) {
	url, ok := store[uuid]
	return url, ok
}

func IsStoreEmpty() bool {
	return len(store) == 0
}

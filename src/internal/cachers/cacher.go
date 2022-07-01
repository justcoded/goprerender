package cachers

type Сacher interface {
	Put(key string, data []byte) error
	Get(key string) ([]byte, error)
	Remove(key string) (bool, error)
	Len() int
	Purge()
}

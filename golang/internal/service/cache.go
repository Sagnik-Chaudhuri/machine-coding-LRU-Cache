package service

type CacheService interface {
	GetValue(key string) (string, error)
	PutValue(key string, value string)
}

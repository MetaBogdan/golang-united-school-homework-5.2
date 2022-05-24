package cache

import "time"

type saveVal struct {
	value  string
	record time.Time
}

type Cache struct {
	massivCash map[string]saveVal
}

func NewCache() Cache {

	return Cache{massivCash: make(map[string]saveVal)}
}

func (in *Cache) Get(key string) (string, bool) {
	if value, ok := in.massivCash[key]; ok {
		if value.record.IsZero() || time.Now().Before(value.record) {

			return value.value, true
		}
		delete(in.massivCash, key)
	}
	return "", false
}

func (in *Cache) Put(key, value string) {

	in.massivCash[key] = saveVal{
		value:  value,
		record: time.Time{},
	}
}

func (in *Cache) Keys() []string {
	listKeys := make([]string, 0)
	for k := range in.massivCash {
		listKeys = append(listKeys, k)

	}
	return listKeys

}

func (in *Cache) PutTill(key, value string, deadline time.Time) {
	in.massivCash[key] = saveVal{
		value:  value,
		record: deadline,
	}

}

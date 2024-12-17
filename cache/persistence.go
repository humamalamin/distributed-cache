package cache

import (
	"encoding/json"
	"os"
)

func (c *Cache) SaveToFile(filename string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	data := make(map[string]string)
	for k, v := range c.items {
		data[k] = v.Value.(*entry).value
	}

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func (c *Cache) LoadFromFile(filename string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make(map[string]string)
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&data); err != nil {
		return err
	}

	for k, v := range data {
		c.Set(k, v)
	}

	return nil
}

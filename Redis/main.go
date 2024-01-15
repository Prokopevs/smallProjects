package main
import (
    "fmt"
    
)

func main() {
    // type KVDatabase interface {
    //     Get(string) (string, error)
    //     GetKeys() ([]string, error)
    // }
    
    dbData := NewRedisDatabase()

    _, err := dbData.Create("hello", "there")
    if err != nil {
        fmt.Println(err)
    }
	_, err = dbData.Create("hello1", "there1")
    if err != nil {
        fmt.Println(err)
    }

    value, err := dbData.Get("hello")
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(value)

    keys, err := dbData.GetKeys()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(keys)
  
}

type RedisDatabase struct {
    data map[string]string 
}
    
func NewRedisDatabase() *RedisDatabase {
    return &RedisDatabase {
        data: make(map[string]string),
    }
}

func (d *RedisDatabase) Create(key, value string) (string, error) {
        d.data[key] = value
        return "", nil
}

func (d *RedisDatabase) Get(key string) (string, error) {
	return d.data[key], nil
}
    
func (d *RedisDatabase) GetKeys() ([]string, error) {
	slice := make([]string, 0, 4)

	for key := range d.data {
		slice = append(slice, key)
	}

	return slice, nil
}




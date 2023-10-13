package router

import "errors"

type OrderedMap struct {
  keys []string
  values  map[string]interface{}
  indices map[string]int
}

func NewOrderedMap() *OrderedMap{
  return &OrderedMap{
    keys: make([]string, 0),
    values: make(map[string]interface{}),
    indices: make(map[string]int),
  }
}

func (om *OrderedMap) Add(key string, value interface{}){
  if _, exists := om.values[key]; !exists {
    om.keys = append(om.keys, key)
    om.indices[key] = len(om.keys) - 1
  }
  om.values[key] = value
}

func (om *OrderedMap) GetByKey(key string) (interface{}, error) {
    value, exists := om.values[key]
    if exists{
      return value, nil
    }
    return nil, errors.New("Key not in map")
}

func (om *OrderedMap) GetByIndex(index int) (string, interface{}, error) {
    if index < 0 || index >= len(om.keys) {
        return "", nil, errors.New("Out of bounds index")
    }
    key := om.keys[index]
    value, _ := om.values[key]
    return key, value, nil
}

func (om *OrderedMap) GetKey(index int) (string, error) {
    if index < 0 || index >= len(om.keys) {
        return "", errors.New("index out of bounds")
    }
    
    key := om.keys[index]
    return key, nil
}

func (om *OrderedMap) GetIndex(key string) (int, error) {
    index, exists := om.indices[key]
    if !exists {
        return -1, errors.New("Key not found")
    }
    return index, nil
}

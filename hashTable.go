package HashTable

var array []*linkList

type linkList struct {
	Key string
	Value interface{}
	NextNode *linkList
}

func (m *model)Add(key string, value interface{}) {
	index := HashFunction(key)
	if index > len(array) {
		list := linkList{}
		list.Key = key
		list.Value = value
		lens := len(array)
		for i := 0; i <= (index - lens); i++ {
			array = append(array, new(linkList))
		}
		array[index] = &list
		return
	}
	list := array[index]
	for {
		if list.Key == "" && list.Value == nil {
			list.Key = key
			list.Value = value
			return
		}
		list = list.NextNode
	}
}

func (m *model)Get(key string) interface{} {
	index := HashFunction(key)
	list := array[index]
	for {
		if list.Key == "" && list.Value == nil {
			return nil
		}
		if list.Key == key {
			return list.Value
		}
		list = list.NextNode
	}
}

type model struct {

}

type HashTable interface {
	Add(key string, value interface{})
	Get(key string) interface{}
}


func NewHashTable()  HashTable{
	var ht HashTable
	ht = new(model)
	return ht
}
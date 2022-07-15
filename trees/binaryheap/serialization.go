package binaryheap

import "github.com/zhangdapeng520/zdpgo_type/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Heap[int])(nil)
	var _ containers.JSONDeserializer = (*Heap[int])(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap[T]) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap[T]) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}

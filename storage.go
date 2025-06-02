package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *storage[T] {
	return &storage[T]{FileName: fileName}
}
func (s *storage[T]) Save(data T) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}
	err = os.WriteFile(s.FileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	return nil
}
func (s *storage[T]) Load(data *T) error {
	file, err := os.ReadFile(s.FileName)
	if err != nil {

		return fmt.Errorf("error reading file: %v", err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %v", err)
	}
	return nil
}

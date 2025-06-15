package main

import (
	"fmt"
	"slices"
)

type Pair[K any, V any] struct {
	first  K
	second V
}

func CPair[K any, V any](first K, second V) Pair[K, V] {
	return Pair[K, V]{first: first, second: second}
}

func FindPair[K any, V any](arr []Pair[K, V], name any) (Pair[K, V], error) {
	for _, item := range arr {
		if Fst(item, nil) == name {
			return item, nil
		}
	}
	return Pair[K, V]{}, fmt.Errorf("findPair: pair with fst %s not found", name)
}

func DeletePair[K any, V any](arr []Pair[K, V], name any) error {
	idx := -1
	for i, item := range arr {
		if Fst(item, nil) == name {
			idx = i
		}
	}
	if idx == -1 {
		return fmt.Errorf("Could not find a pair with name \"%s\" to delete", name)
	}
	arr = slices.Delete(arr, idx, idx+1)
	return nil
}

func Fst[K any, V any](pair Pair[K, V], err error) any {
	if err != nil {
		return nil
	}
	return pair.first
}

func Snd[K any, V any](pair Pair[K, V], err error) any {
	if err != nil {
		return nil
	}
	return pair.second
}

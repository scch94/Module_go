package slices

import "fmt"

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			log(fmt.Sprintf("%s: se encontro el valor %v en la lista", packageName, item))
			return true
		}
	}
	return false
}

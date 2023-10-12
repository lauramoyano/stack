package pop

import "fmt"

func Pop(pila *[]int) (int, error) {
	if len(*pila) == 0 {
		return 0, fmt.Errorf("La pila está vacía")
	}
	item := (*pila)[len(*pila)-1]
	*pila = (*pila)[:len(*pila)-1]
	return item, nil
}

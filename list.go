package main

type List interface {
	Add(value int)
	Remove(index int)
}

func (list *ArrayList) Add(value int) {
	newLength := len(list.elements) + 1
	newArray := make([]int, newLength)

	for i := 0; i < newLength; i++ {
		if i == newLength-1 {
			newArray[i] = value
			break
		}

		newArray[i] = list.elements[i]
	}

	list.elements = newArray
}

func (list ArrayList) Length() int {
	return len(list.elements)
}

type ArrayList struct {
	elements []int
}

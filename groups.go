package schoolutil

import (
	"math/rand"
)

var r = rand.New(rand.NewSource(0123))

type Class struct {
	Students []Student
}

func (c Class) RandomStudent() Student {
	length := int64(len(c.Students))
	index := r.Int63n(length)

	return c.Students[index]
}

func (c Class) ShuffleStudents() []Student {
	shuffled := c.Students
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
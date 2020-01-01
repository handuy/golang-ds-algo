package main

import (
	"log"
	"syreclabs.com/go/faker"
)

type Set struct {
	// data chứa các key không trùng nhau
	data map[int]string
}

func(set *Set) New() {
	set.data = make(map[int]string)
}

func(set *Set) ContainElement(item int) bool {
	var exist bool
	_, exist = set.data[item]
	return exist
}

// Set chỉ chứa những element độc nhất
func(set *Set) AddElement(item int) {
	// Kiểm tra xem trường data trong Set đã có key = item chưa
	// Có rồi thì thôi không add
	if !set.ContainElement(item) {
		set.data[item] = faker.Name().Name()
	}
}

func(set *Set) DeleteElement(item int) {
	delete(set.data, item)
}

// Lấy ra những phần tử chung giữa 2 Set
func(set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()

	var value int
	for value = range set.data {
		if anotherSet.ContainElement(value) {
			intersectSet.AddElement(value)
		}
	}

	return intersectSet
}

func(set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()

	for value := range set.data {
		unionSet.AddElement(value)
	}

	for value := range anotherSet.data {
		unionSet.AddElement(value)
	}

	return unionSet
}

func main() {
	var set = &Set{}
	set.New()
	set.AddElement(0)
	set.AddElement(1)
	set.AddElement(3)

	var newSet = &Set{}
	newSet.New()
	newSet.AddElement(0)
	newSet.AddElement(1)
	newSet.AddElement(2)

	var intersectResult1 = set.Intersect(newSet)
	log.Println("intersect set", *intersectResult1)

	var intersectResult2 = newSet.Intersect(set)
	log.Println("intersect set", *intersectResult2)

	var unionResult1 = set.Union(newSet)
	log.Println("union set", *unionResult1)

	var unionResult2 = newSet.Union(set)
	log.Println("union set", *unionResult2)
}
package people

import "errors"

var people []Person

func GetPeople() []Person {
	return people
}

func GetPerson(id string) (error, Person) {
	for _, person := range people {
		if person.ID == id {
			return nil,person
		}
	}
	return errors.New("invalid person id"), Person{}
}

func CreatePerson(person Person) []Person {
	people = append(people, person)
	return people
}

func ModifyPerson(p Person) (error, []Person) {
	id := p.ID
	for index, person := range people {
		if person.ID == id {
			people = append(people[:index], people[index+1:]...)
			people = append(people, p)
			return nil, people
		}
	}
	return errors.New("error modifying"), people
}

func DeletePerson(id string) (error, []Person) {
	for index, person := range people {
		if person.ID == id {
			people = append(people[:index], people[index+1:]...)
			return nil, people
		}
	}
	return errors.New("error deleting"), people
}
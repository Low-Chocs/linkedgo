package main

import "fmt"

type Node struct {
	name      string
	last_name string
	age       int
	address   string
	career    string
	course    string
	state     bool
	next      *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(name string, last_name string, age int, address string, career string, course string, state bool) {

	n := Node{}
	n.name = name
	n.last_name = last_name
	n.age = age
	n.address = address
	n.career = career
	n.course = course
	n.state = state

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

func (l *LinkedList) Edit(name_to_edit string, new_state bool, new_age int, new_address string) {
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}

	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.name == name_to_edit {
			ptr.state = new_state
			ptr.age = new_age
			ptr.address = new_address
			fmt.Println("Succesfully edited")
			return
		}
		ptr = ptr.next
	}
}

func main() {
	student := LinkedList{}
	student.Print()
	student.Insert("Alberto", "Moreira", 20, "La Florida", "Sistemas", "EDD", false)
	student.Insert("Guillermo", "Nox", 27, "San Juan", "Sistemas", "Compi", true)
	student.Insert("Lionel", "Ronaldo", 30, "Primero de Julio", "Sistemas", "ORGA", false)
	student.Insert("Abel", "Villatoro", 24, "La Florida", "Sistemas", "EDD", true)
	student.Print()
	student.Edit("Lionel", true, 70, "Portugal")
	student.Print()
}

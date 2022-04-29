package main

func main() {
	ll := LinkedList{}
	ll.Print()
	ll.Insert("A")
	ll.Insert("B")
	ll.Insert("C")
	ll.Insert("D")
	ll.Insert("E")
	ll.Print()
	ll.Reverse().Print()
}

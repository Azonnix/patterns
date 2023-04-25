package main

import "fmt"

func main() {
	user1 := &user{
		name: "alex",
		age:  10,
	}

	user2 := &user{
		name: "mic",
		age:  12,
	}

	user3 := &user{
		name: "al",
		age:  13,
	}

	userCollection := collection{
		users: []*user{user1, user2, user3},
	}

	userIterator := userCollection.createIterator()

	for userIterator.hasNext() {
		us := userIterator.getNext()

		fmt.Println(us)
	}
}

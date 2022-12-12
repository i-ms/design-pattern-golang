package iterator_collection

import "fmt"

func App() {
	user1 := &User{
		name: "Mayank Singh",
		age:  26,
	}
	user2 := &User{
		name: "Vivaan Singh",
		age:  25,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %s and have age %d\n", user.name, user.age)
	}
}

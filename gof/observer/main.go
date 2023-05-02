package main

func main() {
	item := car{name: "bmw"}
	item.register(&client{email: "first@f.com"})
	item.register(&client{email: "second@f.com"})
	item.allNotify()
}

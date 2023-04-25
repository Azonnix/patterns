package main

func main() {
	res := &reseption{}
	doc := &doctor{}
	med := &medical{}

	doc.setNext(med)
	res.setNext(doc)

	p := patient{name: "ggg"}

	res.execute(&p)
}

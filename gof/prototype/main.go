package main

func main() {
	f := file{"hello"}

	f.print("data1")

	fCopy := f.copy()

	fCopy.print("data2")
}

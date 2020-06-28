package main

func main() {
	result := store("sampleKey", "sampleValue")
	if result < 0 {
		panic("failed store")
	}
}

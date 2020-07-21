package main

func main() {
	server := NewServer(":8080")
	server.HandleRoute("/", HandleRoot)

	server.Listen()
}

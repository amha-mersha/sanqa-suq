package main

func main() {
	cfg, err := LoadConfig(".env")
	if err != nil {
		panic(err)
	}

}

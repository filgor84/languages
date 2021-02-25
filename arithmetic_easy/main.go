package main

func main() {
	fileNames := [4]string{"data/small.txt", "data/1MB.txt", "data/10MB.txt", "data/20MB.txt"}

	for _, file := range fileNames {
		ParseFile(file, 1)
		ParseFile(file, 2)
		ParseFile(file, 4)

	}

}

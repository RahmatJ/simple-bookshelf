package author

func GenerateAuthor(name string) (result Author) {
	if name == "" {
		name = "Aoyama Gosho"
	}
	result = Author{}
	result.Id = "1"
	result.Name = name
	return
}

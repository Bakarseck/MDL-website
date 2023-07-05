package models

type Data struct {
	Category []string
	Books []Book
}

type Book struct {
	BookImg string
	BookName string
}

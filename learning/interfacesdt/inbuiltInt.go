package interfacesdt

import "fmt"


type Book struct{
	ID     int
	ISBN   string
	Title  string
	Author string
}

func NewBook(id int, isbn string, title string, author string) *Book {
	return &Book{
		ID:     id,
		ISBN:   isbn,
		Title:  title,
		Author: author,
	}
}

// calling inbuilt stringer interface to print all values from struct
func (b *Book) String() string{
	return fmt.Sprintf("Book details: %d - %s - %s - %s ", b.ID, b.ISBN, b.Title, b.Author)
}

func Bookshelf(){
	var bookList []*Book
	var book *Book

	book = NewBook(1, "231fhks", "Adventures of Sherlock Holmes", "Arthur Conan Doyle")
	bookList = append(bookList, book)

	book = NewBook(2, "231fhks", "Adventures of Sherlock Holmes1994", "Arthur Conan Doyle")
	bookList = append(bookList, book)

	// fmt.Println(bookList)

	for _, b := range bookList{
		fmt.Println(b)
	}


}
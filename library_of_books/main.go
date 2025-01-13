package main

import "fmt"

//part-1
type Book struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func CreateBook(title, author string, pages, copiesavailable int) *Book {
	return &Book{
		Title:           title,
		Author:          author,
		Pages:           pages,
		CopiesAvailable: copiesavailable,
	}
}

//part-2
func (b Book) Display() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nCopies Available: %d\n", b.Title, b.Author, b.Pages, b.CopiesAvailable)
}

//part-3
func (b *Book) Borrow() string {
	if b.CopiesAvailable > 0 {
		b.CopiesAvailable--
		return "Borrow Successful"
	}
	return "No copies available to borrow"
}

func (b *Book) ReturnBook() {
	b.CopiesAvailable++
}

//part-4
func SwapTitles(b1, b2 *Book) {
	b1.Title, b2.Title = b2.Title, b1.Title
}

func main() {
	book1 := CreateBook("Ramayana", "Valmiki", 14000, 5)
	book2 := CreateBook("Mahabharatha", "Vedavyas", 9000, 6)

	fmt.Println("Before Borrowing:")
	book1.Display()
	book2.Display()

	fmt.Println("\nBorrowing book1:")
	fmt.Println(book1.Borrow())
	book1.Display()

	fmt.Println("\nReturning book1:")
	book1.ReturnBook()
	book1.Display()

	fmt.Println("\nReturning book2:")
	book2.ReturnBook()
	book2.Display()

	fmt.Println("\nBorrowing book2:")
	fmt.Println(book2.Borrow())
	book2.Display()

	fmt.Println("\nSwapping titles of book1 and book2:")
	SwapTitles(book1, book2)
	book1.Display()
	book2.Display()
}

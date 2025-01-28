package library

import (
	"fmt"
	"strings"
)

type Book struct {
	Title    string
	Author   string
	Quantity int
}

type Library struct {
	Books    []*Book
	Borrowed map[string]int
}

func NewLibrary() *Library {
	return &Library{
		Books:    make([]*Book, 0),
		Borrowed: make(map[string]int),
	}
}

func (l *Library) AddBook(title, author string, quantity int) {
	book := &Book{Title: title, Author: author, Quantity: quantity}
	l.Books = append(l.Books, book)
	fmt.Println("Book added to the library.")
}

func (l *Library) DisplayBooks() {
	if len(l.Books) == 0 {
		fmt.Println("The library is empty.")
	} else {
		fmt.Println("Library Catalog:")
		for _, book := range l.Books {
			fmt.Printf("Title: %s, Author: %s, Available: %d\n", book.Title, book.Author, book.Quantity-l.Borrowed[book.Title])
		}
	}
}

func (l *Library) BorrowBook(title string) {
	for _, book := range l.Books {
		if strings.EqualFold(book.Title, title) {
			if book.Quantity > l.Borrowed[book.Title] {
				l.Borrowed[book.Title]++
				fmt.Println("Book borrowed successfully.")
			} else {
				fmt.Println("All copies of this book are currently borrowed.")
			}
			return
		}
	}
	fmt.Println("Book not found in the library.")
}

func (l *Library) ReturnBook(title string) {
	for _, book := range l.Books {
		if strings.EqualFold(book.Title, title) {
			if l.Borrowed[book.Title] > 0 {
				l.Borrowed[book.Title]--
				fmt.Println("Book returned successfully.")
			} else {
				fmt.Println("You haven't borrowed this book.")
			}
			return
		}
	}
	fmt.Println("Book not found in the library.")
}

func (l *Library) SearchBook(query string) {
	found := false
	for _, book := range l.Books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(query)) {
			fmt.Printf("Title: %s, Author: %s, Available: %d\n", book.Title, book.Author, book.Quantity-l.Borrowed[book.Title])
			found = true
		}
	}
	if !found {
		fmt.Println("No matching books found.")
	}
}

func (l *Library) DisplayBorrowedBooks() {
	if len(l.Borrowed) == 0 {
		fmt.Println("You haven't borrowed any books.")
	} else {
		fmt.Println("Borrowed Books:")
		for title, count := range l.Borrowed {
			fmt.Printf("Title: %s, Quantity: %d\n", title, count)
		}
	}
}

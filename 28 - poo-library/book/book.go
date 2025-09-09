package book

import "fmt"

// Interface
// En esta estructura se cumple esta interface

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title string
	autor string
	pages int
}

// Constructor
func NewBook(title string, autor string, pages int) *Book {
	return &Book{
		title: title,
		autor: autor,
		pages: pages,
	}
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

// b es un receptor
func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAutor: %s\nPages: %d\n", b.title, b.autor, b.pages)
}

// Composicion
type TextBook struct {
	Book
	editorial string
	level     string
}

// Constructor de la nueva estructura
func NewTextBook(title, autor string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, autor, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAutor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n",
		b.title, b.autor, b.pages, b.editorial, b.level)
}

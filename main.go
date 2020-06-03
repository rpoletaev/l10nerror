package main

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,ru
import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	_ "golang.org/x/text/message/catalog"
)

func main() {
	p := message.NewPrinter(language.Russian)
	p.Printf("Hello world!")
	p.Println()
	p.Printf("Hello", "world!")
	p.Println()
	person := "Alex"
	place := "Utah"
	parameter := "birdth date"
	p.Printf("Wrong parameter: %s", parameter)
	p.Println()
	p.Printf("%s is visiting %s!",
		person,
		place)
	p.Println()
	// Double arguments.
	m := 1.2345
	p.Printf("%.2[1]f miles traveled (%[1]f)", m)
	p.Println()
}

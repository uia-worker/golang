package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte // means "a byte slice"; expected type for io libraries
}

/*
 A method named save that takes as its receiver p, a pointer to Page
 It takes no parameters, and returns a value of type error
 It will save the Page's Body to a text file
 Title is used as the file name
 It returns an error value because that is the return type of WriteFile
 WriteFile is a standard library function that writes a byte slice to a file
 If anything go wrong while writing the file, an error value will be returned
 If all goes well, Page.save() will return nil (the zero-value for pointers, interfaces and some other values
 0600, passed as the third parameter to WriteFile, indicated that the file should be created with read-write
permissions for current user only (this is from UNIX)
*/
func (p *Page) save() error {

	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
 This function (or method) constructe the file name from the title parameter,
reads the file's contents into a new variable body, and returns a pointer
to a Page literal constructed with the proper title and body values
 Functions can return multiple values.
 io.ReadFile returns []byte and error
 In this version error is not handled yet; the "blank identifier" represented
by the underscore (_) symbol is used to throw away the error return value
(in essence, assigning the value to nothing, to a "hole")
*/
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

/*
  error handling might be done, in case ReadFile fails (file does not exist, fex)
	* add an if-structure that checks if ReadFile return an err different than <nil>
	* add an extra return parameter "error", to lead error to this variable instance
	* add an extra return value <nil> in the return statement
*/
func loadPageMod(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2 := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

package main

import(
	"fmt"
	"os"
	"log"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error{
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err:= os.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main(){
	p1 := &Page{Title : "TestPage", Body: []byte("This is a simple page")}
	p1.save()
	p2, _:= loadPage("TestPage")
	fmt.Println(string(p2.Body))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
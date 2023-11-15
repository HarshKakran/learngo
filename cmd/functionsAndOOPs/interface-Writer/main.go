package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	fname string
	lname string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.fname + " " + p.lname))
	return err
}

func main() {
	p := person{
		fname: "Harsh",
		lname: "Kakran",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	// s := []byte("Hello Gophers!👋🏻")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// } else {
	// 	fmt.Println("Wrote Successfully")
	// }

	// b := bytes.NewBufferString("Hello ")
	// fmt.Println(b.String())
	// b.WriteString("Gophers!")
	// fmt.Println(b.String())

	// b.Reset()
	// b.WriteString("Hey There!")
	// fmt.Println(b.String())

	// b.Write([]byte(" Happy Happy"))
	// fmt.Println(b.String())

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}

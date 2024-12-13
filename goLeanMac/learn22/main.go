package main

import "fmt"

//i := "ssss"
//var allType interface{}
//allType = i
//str := allType.(string)
//fmt.Println(str)

//os write
//file, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
//if err != nil {
//	fmt.Println("error is", err)
//	return
//} else {
//	fmt.Println(file)
//}
//fmt.Println("-------")
//var read io.Reader
//read = file
//var write io.Writer
//
//write = read.(io.Writer)
//write.Write([]byte("hello"))

type Reader interface {
	ReadBook() // ReadBook 是一个方法，不是类型
}

type Writer interface {
	WriteBook() // ReadBook 是一个方法，不是类型
}

type Book struct {
	// Book fields can be added here if needed
}

// ReadBook 方法实现了 Reader 接口
func (b *Book) ReadBook() {
	fmt.Println("reading a book")
}

func (this *Book) WriteBook() {
	fmt.Println("writing a book")
}
func main() {
	b := &Book{}
	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}

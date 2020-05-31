package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	//var wc WriterCloser = NewBufferedWriterCloser()
	//wc.Write([]byte("hello listeners , this is a test"))
	//wc.Close()
	var myObj interface{} = NewBufferedWriterCloser()
	var wc interface{}
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("hello listeners , this is a test"))
		wc.Close()
	}

	bwc := wc.(*BufferedWriterCloser)
	//bwc := wc.(io.Reader)
	//panic: interface conversion: *main.BufferedWriterCloser is not io.Reader: missing method Read

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	fmt.Println(bwc)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, err
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

package main

import (
	"errors"
	"fmt"
	"io"
)

type CircleByteBuffer struct {
	io.Reader
	io.Writer
	io.Closer
	Datas   []byte `json:"datas"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Size    int    `json:"size"`
	IsClose bool   `json:"is_close"`
}

func NewCircleByteBuffer(len int) *CircleByteBuffer {
	var e = new(CircleByteBuffer)
	e.Datas = make([]byte, len)
	e.Start = 0
	e.End = 0
	e.Size = len
	e.IsClose = false
	return e
}

func (e *CircleByteBuffer) GetLen() int {
	return (e.End + e.Size - e.Start) % e.Size
}

func (e *CircleByteBuffer) IsFull() bool {
	return (e.End+1)%e.Size == e.Start
}

func (e *CircleByteBuffer) IsEmpty() bool {
	return e.End == e.Start
}

func (e *CircleByteBuffer) GetFree() int {
	return e.Size - e.GetLen()
}

func (e *CircleByteBuffer) AddByte(b byte) error {
	if e.IsFull() {
		return errors.New("buffer is full")
	}
	e.Datas[e.End] = b
	e.End = (e.End + 1) % e.Size
	return nil
}

func (e *CircleByteBuffer) PopByte() (byte, error) {
	if e.IsEmpty() {
		return 0, errors.New("buffer si empty")
	}
	val := e.Datas[e.Start]
	e.Datas[e.Start] = 0
	e.Start = (e.Start + 1) % e.Size
	return val, nil
}

func (e *CircleByteBuffer) Close() error {
	e.IsClose = true
	return nil
}

//add []byte to CircleByteBuffer
func (e *CircleByteBuffer) Write(bts []byte) (int, error) {
	if e.IsFull() {
		return -1, errors.New("buffer is full")
	}
	ret := 0
	for i := 0; i < len(bts); i++ {
		err := e.AddByte(bts[i])
		if err != nil {
			return ret, err
		}
		ret++
	}
	return ret, nil
}

//show []byte in CircleByteBuffer
func (e *CircleByteBuffer) Show() error {
	if e.IsEmpty() {
		return errors.New("buffer is empty")
	}
	tmp := e.Start
	for i := 0; i < e.GetLen(); i++ {
		fmt.Printf("datas[%d] == %v", tmp, e.Datas[tmp])
		tmp = (tmp + 1) % e.Size
	}
	return nil
}

func (e *CircleByteBuffer) Read() (int, error) {
	if e.IsEmpty() {
		return -1, errors.New("buffer is empty")
	}
	ret := 0
	for {
		aa, err := e.PopByte()
		if err != nil {
			break
		}
		fmt.Println("read ", aa)
		ret++
	}
	return ret, nil
}

//file read eg:https://blog.csdn.net/whatday/article/details/103938124

func FileT() {

	fmt.Println("todo")
	//use ringbuffer file -> f2
	//todo

}

package main

import "fmt"

type ByteSlice []byte

// Value approach: need to return the updated slice.
func (slice ByteSlice) AppendViaValue(data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

// Pointer approach v2.
func (p *ByteSlice) AppendViaPointer(data []byte) {
	slice := *p
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	*p = slice
}

// Pointer approach v2.
// This approach acutally implements the io.Writer interface!!!
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	*p = slice

	return len(data), nil
}

func TestAppendViaValue() {
	originalSlice := ByteSlice([]byte{'a', 'b', 'c'})
	returnedSlice := originalSlice.AppendViaValue([]byte{'d', 'e', 'f'})
	fmt.Printf("originalSlice: %q, returnedSlice: %q\n", originalSlice, returnedSlice)
}

func TestAppendViaPointer() {
	slice := ByteSlice([]byte{'a', 'b', 'c'})
	slice.AppendViaPointer([]byte{'d', 'e', 'f'})
	fmt.Printf("slice: %q\n", slice)
}

func TestWrite() {
	slice := ByteSlice([]byte{'a', 'b', 'c'})
	count, err := slice.Write([]byte{'d', 'e', 'f'})
	fmt.Println("count:", count, ", err:", err)
}

func main() {
	TestAppendViaValue()
	TestAppendViaPointer()
	TestWrite()

	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Printf("%s", b)
}

package effective

import "os"

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	f := new(os.File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f

	/*
	 * composite literals examples
	 */

	// Example 1
	// f := File{fd, name, nil, 0}
	// return &f

	// Example 2
	// return &File{fd, name, nil, 0}

	// Example 3
	// return &File{fd: fd, name: name}
}

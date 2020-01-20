//
// fd.FD describes data in /proc/<pid>/fd.
//
// Use fd.FD() to create a new fd.FD object
// from data in a path.
//
package fd

//
// Copyright Arkady Maisnikov (jandre@gmail.com)
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//
// Abstraction for /proc/<pid>/fd
//
type FD struct {
	ID   int
	Link string
}

func New(path string) ([]FD, error) {
	var fds []FD

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fds, err
	}
	for _, file := range files {
		filename := file.Name()
		fd, err := strconv.Atoi(filename)
		if err != nil {
			continue
		}
		// See also https://stackoverflow.com/questions/18062026/resolve-symlinks-in-go
		link, err := os.Readlink(filename)
		if err != nil {
			link = filename
		}
		fds = append(fds, FD{ID: fd, Link: link})
	}
	return fds, nil
}

func Sprintf(fds []FD) string {
	s := ""
	for _, fd := range fds {
		s = s + fmt.Sprintf("\n%v", fd)
	}
	return s
}

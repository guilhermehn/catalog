# Catalog
> Walk through the given directory and returns all files within it

## Install
~~~
go get github.com/guilhermehn/catalog
~~~

## Usage
~~~ go
files, err := Catalog("path/to/dir")

onlyImages, err := Catalog("path/to/dir", ".png", ".jpg", ".gif")
// or
exts := []string{".png", ".jpg", ".gif"}
onlyImages, err := Catalog("path/to/dir", exts...)
~~~

## Example

     ├── test_dir
     |   ├── file.md
     |   ├── text.txt
     |   └── original.png
     |   └── dir
     |       └── modified.png

~~~ go
package main

import (
  "fmt"
  "github.com/guilhermehn/catalog"
)

func main() {
  files, err := Catalog("test_dir")

  if err != nil {
    panic(err.Error())
  }

  for _, file := range files {
    fmt.Println(file.FileInfo.Name())
  }
}
~~~

## License
MIT

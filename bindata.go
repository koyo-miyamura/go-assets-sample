package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3b4e08d4c718d774efe29935a4a6842efab4f09b = "<!DOCTYPE html>\n<html lang=\"ja\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n  <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">\n  <title>Document</title>\n</head>\n<body>\n  <h1>Hello World!!</h1>\n</body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1556938275, 1556938275907016325),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1556940445, 1556940445699411194),
		Data:     []byte(_Assets3b4e08d4c718d774efe29935a4a6842efab4f09b),
	}}, "")

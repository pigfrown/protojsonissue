protojson issue
===============

The protojson go package appears to exhibit different behaviour when converting from protobuf messages to 

All tests run main.go as is, then edit main.go to uncomment the second print statement:

# Linux Output

## With go version go1.16.6 linux/amd64

		$ go run main.go
		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}
		$ go run main.go
		{"aString":"somestring","someBytes":"c29tZWJ5dGVz","aNumber":"23"}
		{"aString":"somestring","someBytes":"c29tZWJ5dGVz","aNumber":"23"}	

First run has spaces after commas, second run does not.

## With go version go1.16.7 linux/amd64

		$ go run main.go
		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}
		$ go run main.go
		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}
		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}

Both runs have spaces with newer go version

# Windows Output

Same process as above gives this output: There are never any spaces after commas

## With go version go1.16.2 windows/amd64

Never any spaces:

		PS E:\protojsonissue\cmd\jsonexample> go run main.go
		{"aString":"somestring","someBytes":"c29tZWJ5dGVz","aNumber":"23"}
		PS E:\protojsonissue\cmd\jsonexample> go run main.go
		{"aString":"somestring","someBytes":"c29tZWJ5dGVz","aNumber":"23"}
		{"aString":"somestring","someBytes":"c29tZWJ5dGVz","aNumber":"23"}

## With go version go1.16.7 windows/amd64

Always spaces

		PS E:\protojsonissue\cmd\jsonexample> go run main.go
		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}
		PS E:\protojsonissue\cmd\jsonexample> go run main.go
		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}
		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}
	

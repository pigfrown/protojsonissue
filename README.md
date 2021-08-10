Run main.go and output (on linux) is:

		{"aString":"somestring", "someBytes":"c29tZWJ5dGVz", "aNumber":"23"}

Edit main.go and uncomment the send fmt.Println(). Output (on linux) is now:

	{"aString":"somestring","someBytes":"c29tZWJ5dGVz","aNumber":"23"}
	{"aString":"somestring","someBytes":"c29tZWJ5dGVz","aNumber":"23"}	

First run has spaces after commas, second run does not.

Behaviour is different on windows (no spaces).

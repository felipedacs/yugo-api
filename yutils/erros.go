package yutils

// Check texta e trata o possível erro
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

package main

import (
	"os"

	"github.com/taylormonacelli/africanbackend"
)

func main() {
	code := africanbackend.Execute()
	os.Exit(code)
}

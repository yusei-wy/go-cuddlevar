package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	cuddlevar "github.com/yusei-wy/go-cuddlevar"
)

func main() {
	singlechecker.Main(cuddlevar.Analyzer)
}

package main

import (
	"github.com/gostaticanalysis/noctor"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(noctor.Analyzer) }

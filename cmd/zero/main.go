package main

import (
	"github.com/gostaticanalysis/zero"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(zero.Analyzer) }

package main

import (
	"flag"
	"io/ioutil"
)

var cFile = flag.String("c", "", "output file")

var cHead = `
#include <stdio.h>

int
main() {
  char array[1000];
  char *ptr = array;

`

var cTail = `
  return 0;
}
`

var indent = `  `

func addIndent(indentLevel int, s string) string {
    var indented string
    for i := 0; i < indentLevel; i++ {
        indented += indent
    }
    indented += s
    return indented
}

func makeCSource(bfSource string) string {
    cSource := cHead
    indentLevel := 1

    for _,rune := range bfSource {
		var addC string
		switch rune {
		case '>':
			addC = "++ptr;\n"
		case '<':
			addC = "--ptr;\n"
		case '+':
			addC = "++*ptr;\n"
		case '-':
			addC = "--*ptr;\n"
		case '.':
			addC = "putchar(*ptr);\n"
		case ',':
			addC = "*ptr = getchar();\n"
		case '[':
			addC = "while (*ptr) {\n"
			indentLevel += 1
		case ']':
			addC = "}\n"
			indentLevel -= 1
		default:
			continue
		}

        toAdd := addIndent(indentLevel, addC)
        cSource += toAdd
    }
    cSource += cTail
    return cSource
}

func main() {
	flag.Parse()
	if *cFile == "" {
		panic("No c destination")
	}
	if len(flag.Args()) != 1 {
		panic("No bf source")
	}
	bfFile := (flag.Args())[0]

	bfSource, err := ioutil.ReadFile(bfFile)
	if err != nil {
		panic(err)
	}
    bfString := string(bfSource)

    cSource := makeCSource(bfString)

    outBytes := []byte(cSource)
    err = ioutil.WriteFile(*cFile, outBytes, 0644)
	if err != nil {
		panic(err)
	}
}

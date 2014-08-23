package main

import (
	"flag"
	"io/ioutil"
)

var tokens = map[rune]string {
	'>' : "++ptr;\n",
	'<' : "--ptr;\n",
	'+' : "++*ptr;\n",
	'-' : "--*ptr;\n",
	'.' : "putchar(*ptr);\n",
	',' : "*ptr = getchar();\n",
	'[' : "while (*ptr) {\n",
	']' : "}\n",
}

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
        if c, ok := tokens[rune]; ok {
            if rune == '[' {
                indentLevel += 1
            } else if rune == ']' {
                indentLevel -= 1
            }
            toAdd := addIndent(indentLevel, c)
            cSource += toAdd
        }
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

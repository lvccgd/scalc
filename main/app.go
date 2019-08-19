package main

import (
	"../parser"
	"../set"
	"bufio"
	"fmt"
	"os"
)

// Usage returns usgage cases
func Usage(appName string) {
	fmt.Printf("Usage: %s expression\n", appName)
	fmt.Println("where:")
	fmt.Println(" - expression: [operator set ... setN]")
	fmt.Println(" - operator  : SUM | INT | DIF")
	fmt.Println(" - set       : file | expression")
	fmt.Println("Example: [ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]")
}

func main() {
	var args string
	exprLen := len(os.Args)
	for i := 1; i < exprLen; i++ {
		args += " " + os.Args[i]
	}

	if result, _ := parser.CheckBrackets(args); !result || exprLen <= 1 {
		Usage(os.Args[0])
		os.Exit(1)
	}

	lex := parser.NewLexer(args)
	parser := parser.NewParser(lex)
	expr := parser.Parse()
	set := eval(expr)
	result := set.ToList()
	for _, v := range result {
		fmt.Println(v)
	}
}

func eval(expr parser.IExpression) *set.Set {
	switch exp := expr.(type) {
	case parser.Expression:
		var sets []*set.Set
		for _, e := range exp.Sets {
			sets = append(sets, eval(e))
		}
		return calc(exp.Operator, sets)
	case parser.Set:
		dataSet, err := scanSetFromFile(exp.String())
		if err != nil {
			return set.CreateSet()
		}
		return set.CreateSet(dataSet...)
	}
	return set.CreateSet()
}

func calc(operator string, sets []*set.Set) *set.Set {
	var ret *set.Set
	switch operator {
	case parser.SUM:
		ret = set.Union(sets...)
	case parser.DIF:
		ret = set.Difference(sets...)
	case parser.INT:
		ret = set.Intersection(sets...)
	}
	return ret
}

func scanSetFromFile(fileName string) ([]set.T, error) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Errorf("Open file error: %v", err)
		return nil, err
	}
	defer f.Close()

	var dataSet []set.T
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		dataSet = append(dataSet, scan.Text())
	}

	if err := scan.Err(); err != nil {
		fmt.Errorf("Scan file error: %v", err)
		return nil, err
	}

	return dataSet, nil
}

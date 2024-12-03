package main

import "fmt"

func main() {
	rule := RuleBuilderWithAfterCondition.Build()
	fmt.Println(rule)
}

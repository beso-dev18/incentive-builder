package main

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func main() {
	// Create a new KnowledgeBase
	knowledgeBase := pkg.NewKnowledgeBaseInstance("Example", "0.0.1")

	// Load the rule from a string or file
	ruleString := `
    rule CheckAge "Rule to check if age is above 18" {
        when
            Person.Age > 18
        then
            Person.IsAdult = true;
            Retract("CheckAge");
    }
    `
	library := pkg.NewRuleBuilder(knowledgeBase)
	err := library.BuildRuleFromResource(pkg.NewStringResource(ruleString))
	if err != nil {
		fmt.Println("Failed to load rule:", err)
		return
	}

	// Define the fact
	person := &Person{Age: 20}

	// Create a data context and add your facts
	dataContext := pkg.NewDataContext()
	err = dataContext.Add("Person", person)
	if err != nil {
		fmt.Println("Failed to add fact:", err)
		return
	}

	// Execute the rules
	gruleEngine := engine.NewGruleEngine()
	err = gruleEngine.Execute(dataContext, knowledgeBase)
	if err != nil {
		fmt.Println("Failed to execute rules:", err)
		return
	}

	// Output the result
	fmt.Printf("Person: %+v\n", person)
}

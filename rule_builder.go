package main

import "fmt"

type KPI struct {
	Actual  float64 `json:"actual"`
	Current float64 `json:"current"`
	Target  float64 `json:"target"`
}

type Tier struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	MinThreshold float64 `json:"min_threshold"`
	MaxThreshold float64 `json:"max_threshold"`
	MinOperator  string  `json:"min_operator"`
	MaxOperator  string  `json:"max_operator"`
	ValueFormula string  `json:"value_formula"`
}

type ConditionBuilder struct {
	Tiers          []Tier `json:"tiers"`
	CompareFormula string `json:"compare_formula"`
	ValueHolder    string `json:"value_holder"`
}

type RuleBuilder struct {
	Conditions   []ConditionBuilder `json:"conditions"`
	KPIs         map[string]KPI     `json:"kpis"`
	TotalFormula string             `json:"total_formula"`
}

const ruleTemplate = `
rule %s "%s" {
	when
		%s %s %f && %s %s %f
	then
		%s = %s;
		Retract("%s");
}
`

func (rb *RuleBuilder) Build() string {
	var rule string
	for _, condition := range rb.Conditions {
		for _, tier := range condition.Tiers {
			rule += fmt.Sprintf(
				ruleTemplate,
				tier.Name, tier.Description,
				condition.CompareFormula, tier.MinOperator, tier.MinThreshold, condition.CompareFormula, tier.MaxOperator, tier.MaxThreshold,
				condition.ValueHolder, tier.ValueFormula,
				tier.Name,
			)
		}
	}

	return rule
}

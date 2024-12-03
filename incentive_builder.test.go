package main

var RuleBuilderWithAfterCondition = RuleBuilder{
	Conditions: []ConditionBuilder{
		{
			Tiers: []Tier{
				{
					Name:         "GMVXanhDuongTier0",
					Description:  "GMV Xanh Duong tier 0",
					MinThreshold: 0,
					MaxThreshold: 0.25,
					MinOperator:  ">",
					MaxOperator:  "<=",
					ValueFormula: `kpi[gmv_xanh_duong].actual * 0.01`,
				},
				{
					Name:         "GMVXanhDuongTier1",
					Description:  "GMV Xanh Duong tier 1",
					MinThreshold: 0.25,
					MaxThreshold: 0.5,
					MinOperator:  ">",
					MaxOperator:  "<=",
					ValueFormula: `kpi[gmv_xanh_duong].actual * 0.02`,
				},
				{
					Name:         "GMVXanhDuongTier2",
					Description:  "GMV Xanh Duong tier 2",
					MinThreshold: 0.5,
					MaxThreshold: 0.75,
					MinOperator:  ">",
					MaxOperator:  "<=",
					ValueFormula: `kpi[gmv_xanh_duong].actual * 0.03`,
				},
				{
					Name:         "GMVXanhDuongTier3",
					Description:  "GMV Xanh Duong tier 3",
					MinThreshold: 0.75,
					MaxThreshold: 1,
					MinOperator:  ">",
					MaxOperator:  "<=",
					ValueFormula: `kpi[gmv_xanh_duong].actual * 0.04`,
				},
			},
			CompareFormula: "kpi[gmv_xanh_duong].actual / kpi[gmv_xanh_duong].target",
			ValueHolder:    "gmv_xanh_duong",
		},
		{
			Tiers: []Tier{
				{
					Name:         "GTActiveTier0",
					Description:  "GT Active tier 0",
					MinThreshold: 0,
					MaxThreshold: 0.7,
					MinOperator:  ">",
					MaxOperator:  "<=",
					ValueFormula: `0.5`,
				},
				{
					Name:         "GTActiveTier1",
					Description:  "GT Active tier 1",
					MinThreshold: 0.7,
					MaxThreshold: 1,
					MinOperator:  ">",
					MaxOperator:  "<=",
					ValueFormula: `1`,
				},
			},
			CompareFormula: "kpi[gt_active].actual / kpi[gt_active].target",
			ValueHolder:    "gt_active",
		},
	},
	KPIs: map[string]KPI{
		"gmv_xanh_duong": {
			Actual:  10000000,
			Current: 13000000,
			Target:  30000000,
		},
		"gt_active": {
			Actual:  76,
			Current: 81,
			Target:  100,
		},
	},
	TotalFormula: `gmv_xanh_duong + gt_active`,
}

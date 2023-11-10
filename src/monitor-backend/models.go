package main

type Rule struct {
	Name      string     `yaml:"name"`
	RuleSteps []RuleStep `yaml:"ruleSteps"`
}

type RuleStep struct {
	Name     string `yaml:"name"`
	Resource string `yaml:"resource"`
}

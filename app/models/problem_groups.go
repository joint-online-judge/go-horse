package models

const TableNameProblemGroup = "problem_groups"

// ProblemGroup mapped from table <problem_groups>
type ProblemGroup struct {
	Base
}

// TableName ProblemGroup's table name
func (*ProblemGroup) TableName() string {
	return TableNameProblemGroup
}

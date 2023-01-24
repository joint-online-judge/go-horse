package models

const TableNameProblemProblemSetLink = "problem_problem_set_links"

// ProblemProblemSetLink mapped from table <problem_problem_set_links>
type ProblemProblemSetLink struct {
	ProblemID    string `gorm:"column:problem_id;primaryKey"     json:"problem_id"`
	ProblemSetID string `gorm:"column:problem_set_id;primaryKey" json:"problem_set_id"`
	Position     int32  `gorm:"column:position;not null"         json:"position"`
}

// TableName ProblemProblemSetLink's table name
func (*ProblemProblemSetLink) TableName() string {
	return TableNameProblemProblemSetLink
}
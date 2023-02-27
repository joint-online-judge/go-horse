package model

import "github.com/google/uuid"

const TableNameProblemProblemSetLink = "problem_problem_set_links"

// ProblemProblemSetLink mapped from table <problem_problem_set_links>
type ProblemProblemSetLink struct {
	ProblemId    uuid.UUID  `gorm:"column:problem_id;primaryKey;type:uuid"         json:"problemId"`
	Problem      Problem    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	ProblemSetId uuid.UUID  `gorm:"column:problem_set_id;primaryKey;type:uuid"     json:"problemSetId"`
	ProblemSet   ProblemSet `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Position     int32      `gorm:"column:position;not null;index"                 json:"position"`
}

// TableName ProblemProblemSetLink's table name
func (*ProblemProblemSetLink) TableName() string {
	return TableNameProblemProblemSetLink
}

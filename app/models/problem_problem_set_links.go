package models

import "github.com/google/uuid"

const TableNameProblemProblemSetLink = "problem_problem_set_links"

// ProblemProblemSetLink mapped from table <problem_problem_set_links>
type ProblemProblemSetLink struct {
	ProblemID    uuid.UUID  `gorm:"column:problem_id;primaryKey;type:uuid"         json:"problem_id"`
	Problem      Problem    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	ProblemSetID uuid.UUID  `gorm:"column:problem_set_id;primaryKey;type:uuid"     json:"problem_set_id"`
	ProblemSet   ProblemSet `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Position     int32      `gorm:"column:position;not null"                       json:"position"`
}

// TableName ProblemProblemSetLink's table name
func (*ProblemProblemSetLink) TableName() string {
	return TableNameProblemProblemSetLink
}

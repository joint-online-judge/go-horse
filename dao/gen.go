// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                     = new(Query)
	Domain                *domain
	DomainInvitation      *domainInvitation
	DomainRole            *domainRole
	DomainUser            *domainUser
	Problem               *problem
	ProblemConfig         *problemConfig
	ProblemGroup          *problemGroup
	ProblemProblemSetLink *problemProblemSetLink
	ProblemSet            *problemSet
	Record                *record
	User                  *user
	UserLatestRecord      *userLatestRecord
	UserOauthAccount      *userOauthAccount
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Domain = &Q.Domain
	DomainInvitation = &Q.DomainInvitation
	DomainRole = &Q.DomainRole
	DomainUser = &Q.DomainUser
	Problem = &Q.Problem
	ProblemConfig = &Q.ProblemConfig
	ProblemGroup = &Q.ProblemGroup
	ProblemProblemSetLink = &Q.ProblemProblemSetLink
	ProblemSet = &Q.ProblemSet
	Record = &Q.Record
	User = &Q.User
	UserLatestRecord = &Q.UserLatestRecord
	UserOauthAccount = &Q.UserOauthAccount
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                    db,
		Domain:                newDomain(db, opts...),
		DomainInvitation:      newDomainInvitation(db, opts...),
		DomainRole:            newDomainRole(db, opts...),
		DomainUser:            newDomainUser(db, opts...),
		Problem:               newProblem(db, opts...),
		ProblemConfig:         newProblemConfig(db, opts...),
		ProblemGroup:          newProblemGroup(db, opts...),
		ProblemProblemSetLink: newProblemProblemSetLink(db, opts...),
		ProblemSet:            newProblemSet(db, opts...),
		Record:                newRecord(db, opts...),
		User:                  newUser(db, opts...),
		UserLatestRecord:      newUserLatestRecord(db, opts...),
		UserOauthAccount:      newUserOauthAccount(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Domain                domain
	DomainInvitation      domainInvitation
	DomainRole            domainRole
	DomainUser            domainUser
	Problem               problem
	ProblemConfig         problemConfig
	ProblemGroup          problemGroup
	ProblemProblemSetLink problemProblemSetLink
	ProblemSet            problemSet
	Record                record
	User                  user
	UserLatestRecord      userLatestRecord
	UserOauthAccount      userOauthAccount
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		Domain:                q.Domain.clone(db),
		DomainInvitation:      q.DomainInvitation.clone(db),
		DomainRole:            q.DomainRole.clone(db),
		DomainUser:            q.DomainUser.clone(db),
		Problem:               q.Problem.clone(db),
		ProblemConfig:         q.ProblemConfig.clone(db),
		ProblemGroup:          q.ProblemGroup.clone(db),
		ProblemProblemSetLink: q.ProblemProblemSetLink.clone(db),
		ProblemSet:            q.ProblemSet.clone(db),
		Record:                q.Record.clone(db),
		User:                  q.User.clone(db),
		UserLatestRecord:      q.UserLatestRecord.clone(db),
		UserOauthAccount:      q.UserOauthAccount.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		Domain:                q.Domain.replaceDB(db),
		DomainInvitation:      q.DomainInvitation.replaceDB(db),
		DomainRole:            q.DomainRole.replaceDB(db),
		DomainUser:            q.DomainUser.replaceDB(db),
		Problem:               q.Problem.replaceDB(db),
		ProblemConfig:         q.ProblemConfig.replaceDB(db),
		ProblemGroup:          q.ProblemGroup.replaceDB(db),
		ProblemProblemSetLink: q.ProblemProblemSetLink.replaceDB(db),
		ProblemSet:            q.ProblemSet.replaceDB(db),
		Record:                q.Record.replaceDB(db),
		User:                  q.User.replaceDB(db),
		UserLatestRecord:      q.UserLatestRecord.replaceDB(db),
		UserOauthAccount:      q.UserOauthAccount.replaceDB(db),
	}
}

type queryCtx struct {
	Domain                IDomainDo
	DomainInvitation      IDomainInvitationDo
	DomainRole            IDomainRoleDo
	DomainUser            IDomainUserDo
	Problem               IProblemDo
	ProblemConfig         IProblemConfigDo
	ProblemGroup          IProblemGroupDo
	ProblemProblemSetLink IProblemProblemSetLinkDo
	ProblemSet            IProblemSetDo
	Record                IRecordDo
	User                  IUserDo
	UserLatestRecord      IUserLatestRecordDo
	UserOauthAccount      IUserOauthAccountDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Domain:                q.Domain.WithContext(ctx),
		DomainInvitation:      q.DomainInvitation.WithContext(ctx),
		DomainRole:            q.DomainRole.WithContext(ctx),
		DomainUser:            q.DomainUser.WithContext(ctx),
		Problem:               q.Problem.WithContext(ctx),
		ProblemConfig:         q.ProblemConfig.WithContext(ctx),
		ProblemGroup:          q.ProblemGroup.WithContext(ctx),
		ProblemProblemSetLink: q.ProblemProblemSetLink.WithContext(ctx),
		ProblemSet:            q.ProblemSet.WithContext(ctx),
		Record:                q.Record.WithContext(ctx),
		User:                  q.User.WithContext(ctx),
		UserLatestRecord:      q.UserLatestRecord.WithContext(ctx),
		UserOauthAccount:      q.UserOauthAccount.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}

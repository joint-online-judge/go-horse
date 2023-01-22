// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/joint-online-judge/go-horse/model"
)

func newProblem(db *gorm.DB, opts ...gen.DOOption) problem {
	_problem := problem{}

	_problem.problemDo.UseDB(db, opts...)
	_problem.problemDo.UseModel(&model.Problem{})

	tableName := _problem.problemDo.TableName()
	_problem.ALL = field.NewAsterisk(tableName)
	_problem.CreatedAt = field.NewTime(tableName, "created_at")
	_problem.UpdatedAt = field.NewTime(tableName, "updated_at")
	_problem.DomainID = field.NewString(tableName, "domain_id")
	_problem.OwnerID = field.NewString(tableName, "owner_id")
	_problem.ProblemGroupID = field.NewString(tableName, "problem_group_id")
	_problem.ID = field.NewString(tableName, "id")
	_problem.URL = field.NewString(tableName, "url")
	_problem.Title = field.NewString(tableName, "title")
	_problem.Content = field.NewString(tableName, "content")
	_problem.Hidden = field.NewBool(tableName, "hidden")
	_problem.NumSubmit = field.NewInt32(tableName, "num_submit")
	_problem.NumAccept = field.NewInt32(tableName, "num_accept")
	_problem.Languages = field.NewString(tableName, "languages")

	_problem.fillFieldMap()

	return _problem
}

type problem struct {
	problemDo

	ALL            field.Asterisk
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DomainID       field.String
	OwnerID        field.String
	ProblemGroupID field.String
	ID             field.String
	URL            field.String
	Title          field.String
	Content        field.String
	Hidden         field.Bool
	NumSubmit      field.Int32
	NumAccept      field.Int32
	Languages      field.String

	fieldMap map[string]field.Expr
}

func (p problem) Table(newTableName string) *problem {
	p.problemDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p problem) As(alias string) *problem {
	p.problemDo.DO = *(p.problemDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *problem) updateTableName(table string) *problem {
	p.ALL = field.NewAsterisk(table)
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DomainID = field.NewString(table, "domain_id")
	p.OwnerID = field.NewString(table, "owner_id")
	p.ProblemGroupID = field.NewString(table, "problem_group_id")
	p.ID = field.NewString(table, "id")
	p.URL = field.NewString(table, "url")
	p.Title = field.NewString(table, "title")
	p.Content = field.NewString(table, "content")
	p.Hidden = field.NewBool(table, "hidden")
	p.NumSubmit = field.NewInt32(table, "num_submit")
	p.NumAccept = field.NewInt32(table, "num_accept")
	p.Languages = field.NewString(table, "languages")

	p.fillFieldMap()

	return p
}

func (p *problem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *problem) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 13)
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["domain_id"] = p.DomainID
	p.fieldMap["owner_id"] = p.OwnerID
	p.fieldMap["problem_group_id"] = p.ProblemGroupID
	p.fieldMap["id"] = p.ID
	p.fieldMap["url"] = p.URL
	p.fieldMap["title"] = p.Title
	p.fieldMap["content"] = p.Content
	p.fieldMap["hidden"] = p.Hidden
	p.fieldMap["num_submit"] = p.NumSubmit
	p.fieldMap["num_accept"] = p.NumAccept
	p.fieldMap["languages"] = p.Languages
}

func (p problem) clone(db *gorm.DB) problem {
	p.problemDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p problem) replaceDB(db *gorm.DB) problem {
	p.problemDo.ReplaceDB(db)
	return p
}

type problemDo struct{ gen.DO }

type IProblemDo interface {
	gen.SubQuery
	Debug() IProblemDo
	WithContext(ctx context.Context) IProblemDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProblemDo
	WriteDB() IProblemDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProblemDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProblemDo
	Not(conds ...gen.Condition) IProblemDo
	Or(conds ...gen.Condition) IProblemDo
	Select(conds ...field.Expr) IProblemDo
	Where(conds ...gen.Condition) IProblemDo
	Order(conds ...field.Expr) IProblemDo
	Distinct(cols ...field.Expr) IProblemDo
	Omit(cols ...field.Expr) IProblemDo
	Join(table schema.Tabler, on ...field.Expr) IProblemDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProblemDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProblemDo
	Group(cols ...field.Expr) IProblemDo
	Having(conds ...gen.Condition) IProblemDo
	Limit(limit int) IProblemDo
	Offset(offset int) IProblemDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProblemDo
	Unscoped() IProblemDo
	Create(values ...*model.Problem) error
	CreateInBatches(values []*model.Problem, batchSize int) error
	Save(values ...*model.Problem) error
	First() (*model.Problem, error)
	Take() (*model.Problem, error)
	Last() (*model.Problem, error)
	Find() ([]*model.Problem, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Problem, err error)
	FindInBatches(result *[]*model.Problem, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Problem) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProblemDo
	Assign(attrs ...field.AssignExpr) IProblemDo
	Joins(fields ...field.RelationField) IProblemDo
	Preload(fields ...field.RelationField) IProblemDo
	FirstOrInit() (*model.Problem, error)
	FirstOrCreate() (*model.Problem, error)
	FindByPage(offset int, limit int) (result []*model.Problem, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProblemDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p problemDo) Debug() IProblemDo {
	return p.withDO(p.DO.Debug())
}

func (p problemDo) WithContext(ctx context.Context) IProblemDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p problemDo) ReadDB() IProblemDo {
	return p.Clauses(dbresolver.Read)
}

func (p problemDo) WriteDB() IProblemDo {
	return p.Clauses(dbresolver.Write)
}

func (p problemDo) Session(config *gorm.Session) IProblemDo {
	return p.withDO(p.DO.Session(config))
}

func (p problemDo) Clauses(conds ...clause.Expression) IProblemDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p problemDo) Returning(value interface{}, columns ...string) IProblemDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p problemDo) Not(conds ...gen.Condition) IProblemDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p problemDo) Or(conds ...gen.Condition) IProblemDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p problemDo) Select(conds ...field.Expr) IProblemDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p problemDo) Where(conds ...gen.Condition) IProblemDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p problemDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IProblemDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p problemDo) Order(conds ...field.Expr) IProblemDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p problemDo) Distinct(cols ...field.Expr) IProblemDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p problemDo) Omit(cols ...field.Expr) IProblemDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p problemDo) Join(table schema.Tabler, on ...field.Expr) IProblemDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p problemDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProblemDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p problemDo) RightJoin(table schema.Tabler, on ...field.Expr) IProblemDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p problemDo) Group(cols ...field.Expr) IProblemDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p problemDo) Having(conds ...gen.Condition) IProblemDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p problemDo) Limit(limit int) IProblemDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p problemDo) Offset(offset int) IProblemDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p problemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProblemDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p problemDo) Unscoped() IProblemDo {
	return p.withDO(p.DO.Unscoped())
}

func (p problemDo) Create(values ...*model.Problem) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p problemDo) CreateInBatches(values []*model.Problem, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p problemDo) Save(values ...*model.Problem) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p problemDo) First() (*model.Problem, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Problem), nil
	}
}

func (p problemDo) Take() (*model.Problem, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Problem), nil
	}
}

func (p problemDo) Last() (*model.Problem, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Problem), nil
	}
}

func (p problemDo) Find() ([]*model.Problem, error) {
	result, err := p.DO.Find()
	return result.([]*model.Problem), err
}

func (p problemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Problem, err error) {
	buf := make([]*model.Problem, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p problemDo) FindInBatches(result *[]*model.Problem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p problemDo) Attrs(attrs ...field.AssignExpr) IProblemDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p problemDo) Assign(attrs ...field.AssignExpr) IProblemDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p problemDo) Joins(fields ...field.RelationField) IProblemDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p problemDo) Preload(fields ...field.RelationField) IProblemDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p problemDo) FirstOrInit() (*model.Problem, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Problem), nil
	}
}

func (p problemDo) FirstOrCreate() (*model.Problem, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Problem), nil
	}
}

func (p problemDo) FindByPage(offset int, limit int) (result []*model.Problem, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p problemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p problemDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p problemDo) Delete(models ...*model.Problem) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *problemDo) withDO(do gen.Dao) *problemDo {
	p.DO = *do.(*gen.DO)
	return p
}
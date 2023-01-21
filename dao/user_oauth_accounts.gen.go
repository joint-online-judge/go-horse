// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

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

func newUserOauthAccount(db *gorm.DB, opts ...gen.DOOption) userOauthAccount {
	_userOauthAccount := userOauthAccount{}

	_userOauthAccount.userOauthAccountDo.UseDB(db, opts...)
	_userOauthAccount.userOauthAccountDo.UseModel(&model.UserOauthAccount{})

	tableName := _userOauthAccount.userOauthAccountDo.TableName()
	_userOauthAccount.ALL = field.NewAsterisk(tableName)
	_userOauthAccount.CreatedAt = field.NewTime(tableName, "created_at")
	_userOauthAccount.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userOauthAccount.UserID = field.NewString(tableName, "user_id")
	_userOauthAccount.ID = field.NewString(tableName, "id")
	_userOauthAccount.OauthName = field.NewString(tableName, "oauth_name")
	_userOauthAccount.AccessToken = field.NewString(tableName, "access_token")
	_userOauthAccount.RefreshToken = field.NewString(tableName, "refresh_token")
	_userOauthAccount.ExpiresAt = field.NewInt32(tableName, "expires_at")
	_userOauthAccount.AccountID = field.NewString(tableName, "account_id")
	_userOauthAccount.AccountName = field.NewString(tableName, "account_name")
	_userOauthAccount.AccountEmail = field.NewString(tableName, "account_email")

	_userOauthAccount.fillFieldMap()

	return _userOauthAccount
}

type userOauthAccount struct {
	userOauthAccountDo

	ALL          field.Asterisk
	CreatedAt    field.Time
	UpdatedAt    field.Time
	UserID       field.String
	ID           field.String
	OauthName    field.String
	AccessToken  field.String
	RefreshToken field.String
	ExpiresAt    field.Int32
	AccountID    field.String
	AccountName  field.String
	AccountEmail field.String

	fieldMap map[string]field.Expr
}

func (u userOauthAccount) Table(newTableName string) *userOauthAccount {
	u.userOauthAccountDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userOauthAccount) As(alias string) *userOauthAccount {
	u.userOauthAccountDo.DO = *(u.userOauthAccountDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userOauthAccount) updateTableName(table string) *userOauthAccount {
	u.ALL = field.NewAsterisk(table)
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.UserID = field.NewString(table, "user_id")
	u.ID = field.NewString(table, "id")
	u.OauthName = field.NewString(table, "oauth_name")
	u.AccessToken = field.NewString(table, "access_token")
	u.RefreshToken = field.NewString(table, "refresh_token")
	u.ExpiresAt = field.NewInt32(table, "expires_at")
	u.AccountID = field.NewString(table, "account_id")
	u.AccountName = field.NewString(table, "account_name")
	u.AccountEmail = field.NewString(table, "account_email")

	u.fillFieldMap()

	return u
}

func (u *userOauthAccount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userOauthAccount) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 11)
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["id"] = u.ID
	u.fieldMap["oauth_name"] = u.OauthName
	u.fieldMap["access_token"] = u.AccessToken
	u.fieldMap["refresh_token"] = u.RefreshToken
	u.fieldMap["expires_at"] = u.ExpiresAt
	u.fieldMap["account_id"] = u.AccountID
	u.fieldMap["account_name"] = u.AccountName
	u.fieldMap["account_email"] = u.AccountEmail
}

func (u userOauthAccount) clone(db *gorm.DB) userOauthAccount {
	u.userOauthAccountDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userOauthAccount) replaceDB(db *gorm.DB) userOauthAccount {
	u.userOauthAccountDo.ReplaceDB(db)
	return u
}

type userOauthAccountDo struct{ gen.DO }

type IUserOauthAccountDo interface {
	gen.SubQuery
	Debug() IUserOauthAccountDo
	WithContext(ctx context.Context) IUserOauthAccountDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserOauthAccountDo
	WriteDB() IUserOauthAccountDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserOauthAccountDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserOauthAccountDo
	Not(conds ...gen.Condition) IUserOauthAccountDo
	Or(conds ...gen.Condition) IUserOauthAccountDo
	Select(conds ...field.Expr) IUserOauthAccountDo
	Where(conds ...gen.Condition) IUserOauthAccountDo
	Order(conds ...field.Expr) IUserOauthAccountDo
	Distinct(cols ...field.Expr) IUserOauthAccountDo
	Omit(cols ...field.Expr) IUserOauthAccountDo
	Join(table schema.Tabler, on ...field.Expr) IUserOauthAccountDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserOauthAccountDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserOauthAccountDo
	Group(cols ...field.Expr) IUserOauthAccountDo
	Having(conds ...gen.Condition) IUserOauthAccountDo
	Limit(limit int) IUserOauthAccountDo
	Offset(offset int) IUserOauthAccountDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserOauthAccountDo
	Unscoped() IUserOauthAccountDo
	Create(values ...*model.UserOauthAccount) error
	CreateInBatches(values []*model.UserOauthAccount, batchSize int) error
	Save(values ...*model.UserOauthAccount) error
	First() (*model.UserOauthAccount, error)
	Take() (*model.UserOauthAccount, error)
	Last() (*model.UserOauthAccount, error)
	Find() ([]*model.UserOauthAccount, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserOauthAccount, err error)
	FindInBatches(result *[]*model.UserOauthAccount, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserOauthAccount) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserOauthAccountDo
	Assign(attrs ...field.AssignExpr) IUserOauthAccountDo
	Joins(fields ...field.RelationField) IUserOauthAccountDo
	Preload(fields ...field.RelationField) IUserOauthAccountDo
	FirstOrInit() (*model.UserOauthAccount, error)
	FirstOrCreate() (*model.UserOauthAccount, error)
	FindByPage(offset int, limit int) (result []*model.UserOauthAccount, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserOauthAccountDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userOauthAccountDo) Debug() IUserOauthAccountDo {
	return u.withDO(u.DO.Debug())
}

func (u userOauthAccountDo) WithContext(ctx context.Context) IUserOauthAccountDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userOauthAccountDo) ReadDB() IUserOauthAccountDo {
	return u.Clauses(dbresolver.Read)
}

func (u userOauthAccountDo) WriteDB() IUserOauthAccountDo {
	return u.Clauses(dbresolver.Write)
}

func (u userOauthAccountDo) Session(config *gorm.Session) IUserOauthAccountDo {
	return u.withDO(u.DO.Session(config))
}

func (u userOauthAccountDo) Clauses(conds ...clause.Expression) IUserOauthAccountDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userOauthAccountDo) Returning(value interface{}, columns ...string) IUserOauthAccountDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userOauthAccountDo) Not(conds ...gen.Condition) IUserOauthAccountDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userOauthAccountDo) Or(conds ...gen.Condition) IUserOauthAccountDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userOauthAccountDo) Select(conds ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userOauthAccountDo) Where(conds ...gen.Condition) IUserOauthAccountDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userOauthAccountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserOauthAccountDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userOauthAccountDo) Order(conds ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userOauthAccountDo) Distinct(cols ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userOauthAccountDo) Omit(cols ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userOauthAccountDo) Join(table schema.Tabler, on ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userOauthAccountDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userOauthAccountDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userOauthAccountDo) Group(cols ...field.Expr) IUserOauthAccountDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userOauthAccountDo) Having(conds ...gen.Condition) IUserOauthAccountDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userOauthAccountDo) Limit(limit int) IUserOauthAccountDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userOauthAccountDo) Offset(offset int) IUserOauthAccountDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userOauthAccountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserOauthAccountDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userOauthAccountDo) Unscoped() IUserOauthAccountDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userOauthAccountDo) Create(values ...*model.UserOauthAccount) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userOauthAccountDo) CreateInBatches(values []*model.UserOauthAccount, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userOauthAccountDo) Save(values ...*model.UserOauthAccount) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userOauthAccountDo) First() (*model.UserOauthAccount, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOauthAccount), nil
	}
}

func (u userOauthAccountDo) Take() (*model.UserOauthAccount, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOauthAccount), nil
	}
}

func (u userOauthAccountDo) Last() (*model.UserOauthAccount, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOauthAccount), nil
	}
}

func (u userOauthAccountDo) Find() ([]*model.UserOauthAccount, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserOauthAccount), err
}

func (u userOauthAccountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserOauthAccount, err error) {
	buf := make([]*model.UserOauthAccount, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userOauthAccountDo) FindInBatches(result *[]*model.UserOauthAccount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userOauthAccountDo) Attrs(attrs ...field.AssignExpr) IUserOauthAccountDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userOauthAccountDo) Assign(attrs ...field.AssignExpr) IUserOauthAccountDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userOauthAccountDo) Joins(fields ...field.RelationField) IUserOauthAccountDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userOauthAccountDo) Preload(fields ...field.RelationField) IUserOauthAccountDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userOauthAccountDo) FirstOrInit() (*model.UserOauthAccount, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOauthAccount), nil
	}
}

func (u userOauthAccountDo) FirstOrCreate() (*model.UserOauthAccount, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOauthAccount), nil
	}
}

func (u userOauthAccountDo) FindByPage(offset int, limit int) (result []*model.UserOauthAccount, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userOauthAccountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userOauthAccountDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userOauthAccountDo) Delete(models ...*model.UserOauthAccount) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userOauthAccountDo) withDO(do gen.Dao) *userOauthAccountDo {
	u.DO = *do.(*gen.DO)
	return u
}

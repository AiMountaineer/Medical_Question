package models

import (
	"MedicalQuestionnaireService/libs"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type IdArray struct {
	ID []uint `json:"id"`
}

type SumRes struct {
	Total int64 `json:"total"`
}

// Field 查询字段结构体
type Field struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

type Relate struct {
	Value string
	Func  interface{}
}

// Search 查询参数结构体
type Search struct {
	Fields    []*Field  `json:"fields"`
	Relations []*Relate `json:"relations"`
	OrderBy   string    `json:"orderBy"`
	Sort      string    `json:"sort"`
	Limit     int       `json:"limit"`  // 页面大小
	Offset    int       `json:"offset"` // 页数
}

// GetAll 批量查询
func GetAll(model interface{}, s *Search) *gorm.DB {
	db := libs.Db.Model(model)
	sort := "desc"
	orderBy := "created_at"
	if len(s.Sort) > 0 {
		sort = s.Sort
	}
	if len(s.OrderBy) > 0 {
		orderBy = s.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", orderBy, sort))

	db.Scopes(FoundByWhere(s.Fields), Relation(s.Relations))

	return db
}

// Found 查询条件
func Found(s *Search) *gorm.DB {
	return libs.Db.Scopes(Relation(s.Relations), FoundByWhere(s.Fields))
}

// IsNotFound 判断是否是查询不存在错误
func IsNotFound(err error) bool {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
		color.Yellow("検索結果が存在しません")
		return true
	}
	return false
}

// Update 更新
func Update(v, d interface{}, id uint) error {
	if err := libs.Db.Model(v).Where("id = ?", id).Updates(d).Error; err != nil {
		color.Red(fmt.Sprintf("Update %+v to %+v\n", v, d))
		return err
	}
	return nil
}

// Relation 加载关联关系
func Relation(relates []*Relate) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(relates) > 0 {
			for _, re := range relates {
				if len(re.Value) > 0 {
					if re.Func != nil {
						db = db.Preload(re.Value, re.Func)
					} else {
						db = db.Preload(re.Value)
					}
				}
				color.Yellow(fmt.Sprintf("Preoad %s", re))
			}
		}
		return db
	}
}

// FoundByWhere 查询条件
func FoundByWhere(fields []*Field) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) > 0 {
			for _, field := range fields {
				if field != nil {
					if field.Condition == "" {
						field.Condition = "="
					}
					if value, ok := field.Value.(int); ok {
						if value >= 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(uint); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]int); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else {
						color.Red(fmt.Sprintf("未知のデータ型：%+v", field.Value))
					}
				}
			}
		}
		return db
	}
}

// DropTables 删除数据表
//func DropTables() {
//	_ = libs.Db.Migrator().DropTable(
//		libs.Config.DB.Prefix+"users",
//		libs.Config.DB.Prefix+"m_role",
//		libs.Config.DB.Prefix+"m_permission",
//		libs.Config.DB.Prefix+"articles",
//		libs.Config.DB.Prefix+"configs",
//		libs.Config.DB.Prefix+"tags",
//		libs.Config.DB.Prefix+"types",
//		libs.Config.DB.Prefix+"article_tags",
//		"casbin_rule")
//}

// Migrate 迁移数据表
func Migrate() {
	err := libs.Db.AutoMigrate(
	//	&User{},
	//	&Role{},
	//&Permission{},
	//	&gormadapter.CasbinRule{},
	)

	if err != nil {
		color.Yellow(fmt.Sprintf("テーブルの初期化にエラーが発生しました ：%+v", err))
	}
}

// Paginate 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = -1
		case pageSize == 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if page < 0 {
			offset = -1
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

// GetPermissionsForUser 获取角色权限
// func GetPermissionsForUser(uid uint) [][]string {
// 	return libs.Enforcer.GetPermissionsForUser(strconv.FormatUint(uint64(uid), 10))
// }

func Transaction(fs ...func(*gorm.DB) error) error {
	t := func(tx *gorm.DB) error {
		for _, f := range fs {
			if err := f(tx); err != nil {
				return err
			}
		}
		return nil
	}
	return libs.Db.Transaction(t)
}

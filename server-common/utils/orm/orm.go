package orm

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"reflect"
)

// Paginate page, pageSize 传入参数
// s 分页默认配置
func Paginate(page interface{}, pageSize interface{}, s map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		p := cast.ToInt(page)
		size := cast.ToInt(pageSize)
		if page == 0 {
			if s["page"] != nil {
				page = s["page"].(int)
			} else {
				page = 1
			}
		}

		if pageSize == 0 {
			if s["pageSize"] != nil {
				pageSize = s["pageSize"].(int)
			} else {
				pageSize = 10
			}
		}
		offset := (p - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

// FilterString 进行快速条件过滤
func FilterString(key, value, operator string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		returnDB := db
		if value != "" {
			switch operator {
			case "like":
				returnDB = returnDB.Where(key+" Like ? ", "%"+value+"%")
			case "=", ">=", "<=", "<":
				returnDB = returnDB.Where(key+" "+operator+" "+"?", value)
			}
		}
		return returnDB
	}
}

// InOrNotInFilter where in 或者 where not in
func InOrNotInFilter(key string, value interface{}, operator string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		returnDB := db
		if reflect.ValueOf(value).Len() != 0 {
			whereMap := map[string]interface{}{
				key: value,
			}
			switch operator {
			case "in":
				returnDB = returnDB.Where(whereMap)
			case "not in":
				returnDB = returnDB.Not(whereMap)
			}
		}
		return returnDB
	}
}

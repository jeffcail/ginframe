package models

import (
	"github.com/jeffcail/ginframe/server-common/utils/orm"
	"github.com/jeffcail/ginframe/server-user/core"
	"time"
)

type User struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Username  string    `gorm:"column:username;NOT NULL;comment:'用户名'"`
	Nickname  string    `gorm:"column:nickname;NOT NULL;comment:'昵称'"`
	Phone     string    `gorm:"column:phone;NOT NULL;comment:'手机号'"`
	Email     string    `gorm:"column:email;NOT NULL;comment:'邮箱'"`
	Gender    int8      `gorm:"column:gender;default:1;NOT NULL;comment:'性别'"`
	RoleId    string    `gorm:"column:role_id;NOT NULL;comment:'角色id'"`
	Password  string    `gorm:"column:password;NOT NULL;comment:'密码'"`
	Enable    int8      `gorm:"column:enable;default:1;NOT NULL;comment:'是否启用 1: 启用 2: 禁用'"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL;comment:'更新时间'"`
}

func (a *User) TableName() string {
	return "user"
}

// FindUserById 根据 id 查找user账户信息
func (a *User) FindUserById(id int64) (user *User, err error) {
	if err = core.Db.Where("id = ?", id).First(user).Error; err != nil {
		return
	}
	return
}

// FindUserByUsername 根据 username 查找user账户信息
func (a *User) FindUserByUsername(username string) (user *User, err error) {
	user = new(User)
	if err = core.Db.Where("username = ?", username).First(user).Error; err != nil {
		return
	}
	return
}

// FindUserByNickname 根据 nickname 查找user账户信息
func (a *User) FindUserByNickname(nickname string) (user *User, err error) {
	if err = core.Db.Where("nickname", nickname).First(user).Error; err != nil {
		return
	}
	return
}

// List user列表
func (a *User) List(page, pageSize int64) (int64, []*User, error) {
	page = 1
	pageSize = 10

	// 项目配置的默认分页
	s := make(map[string]interface{})
	s["page"] = 1
	s["pageSize"] = 10

	var count int64
	us := make([]*User, 0)
	err := core.Db.Table("user").Scopes(orm.Paginate(page, pageSize, s)).Count(&count).Find(&us).Error
	if err != nil {
		return 0, nil, err
	}
	return count, us, nil
}

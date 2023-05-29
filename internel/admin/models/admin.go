package models

import (
	"github.com/jeffcail/ginframe/core/db"
	"github.com/jeffcail/ginframe/utils/orm"
	"time"
)

type Admin struct {
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

func (a *Admin) TableName() string {
	return "admin"
}

// FindAdminById 根据 id 查找admin账户信息
func (a *Admin) FindAdminById(id int64) (admin *Admin, err error) {
	if err = db.Db.Where("id = ?", id).First(admin).Error; err != nil {
		return
	}
	return
}

// FindAdminByUsername 根据 username 查找admin账户信息
func (a *Admin) FindAdminByUsername(username string) (admin *Admin, err error) {
	admin = new(Admin)
	if err = db.Db.Where("username = ?", username).First(admin).Error; err != nil {
		return
	}
	return
}

// FindAdminByNickname 根据 nickname 查找admin账户信息
func (a *Admin) FindAdminByNickname(nickname string) (admin *Admin, err error) {
	if err = db.Db.Where("nickname", nickname).First(admin).Error; err != nil {
		return
	}
	return
}

// List 管理员账号列表
func (a *Admin) List(page, pageSize int64) (int64, []*Admin, error) {
	page = 1
	pageSize = 10

	// 项目配置的默认分页
	s := make(map[string]interface{})
	s["page"] = 1
	s["pageSize"] = 10

	var count int64
	admins := make([]*Admin, 0)
	err := db.Db.Table("admin").Scopes(orm.Paginate(page, pageSize, s)).Count(&count).Find(&admins).Error
	if err != nil {
		return 0, nil, err
	}
	return count, admins, nil
}

package gocrud_gorm_test

import (
	"context"
	"fmt"
	"github.com/kordar/gocrud"
	gocrud_gorm "github.com/kordar/gocrud-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

type SysAdmin struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Phone      string    `gorm:"type:varchar(32);unique;not null" json:"phone"`
	Avatar     string    `gorm:"type:varchar(255)" json:"avatar"`
	Username   string    `gorm:"type:varchar(64);unique;not null" json:"username"`
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	Status     int       `gorm:"default:0" json:"status"`
	Type       int       `gorm:"default:0" json:"type"`
	Salt       string    `gorm:"type:varchar(32)" json:"salt"`
	Version    int       `gorm:"default:0" json:"version"`
	Deleted    int       `gorm:"default:0" json:"deleted"`
	CreateTime time.Time `gorm:"type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"type:datetime" json:"update_time"`
}

func (sysAdmin *SysAdmin) TableName() string {
	return "sys_admin"
}

func db() *gorm.DB {
	dsn := "root:yunpengai.306@tcp(43.139.223.7:3307)/goadmin?charset=utf8&parseTime=true"
	mysqlConfig := gorm.Config{}
	mysqlConfig.Logger = logger.Default.LogMode(logger.Info)
	ins, err := gorm.Open(mysql.Open(dsn), &mysqlConfig)
	if err != nil {
		panic(err)
	}
	return ins
}

func TestCreate(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewFormBody("gorm", context.Background())
	body.Object = map[string]interface{}{
		"phone":       "112233",
		"username":    "AAAAAAA",
		"password":    "123456",
		"create_time": time.Now(),
		"update_time": time.Now(),
	}
	formBody := gocrud_gorm.NewGormFormBody(body)
	model, _ := formBody.Create(&SysAdmin{}, d, nil)
	fmt.Printf("------------%v\n", model)
}

func TestFormBody_Query(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewFormBody("gorm", context.Background())
	body.Conditions = []gocrud.Condition{
		{"", "phone", "", "1122", "", "LIKE", false},
	}
	body.Object = map[string]string{"aaa": "bbb", "ccc": "ddd"}
	formBody := gocrud_gorm.NewGormFormBody(body)
	list := make([]SysAdmin, 0)
	formBody.GormQuery(d, nil).Find(&list)
	fmt.Printf("============%v\n", list)
	mm := map[string]string{}
	err := formBody.Unmarshal(&mm)
	fmt.Printf("----------%v---%v\n", err, mm)
}

func TestFormBody_QuerySafe(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewFormBody("gorm", context.Background())
	body.Conditions = []gocrud.Condition{
		{"", "phone", "", "13389452031", "", "EQ", false},
	}
	formBody := gocrud_gorm.NewGormFormBody(body)
	tx, err := formBody.GormQuerySafe(d, nil)
	list := make([]SysAdmin, 0)
	if err == nil {
		tx.Find(&list)
	}
	fmt.Printf("============%v=====%v\n", list, err)
}

func TestFormBody_Update(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewFormBody("gorm", context.Background())
	body.Conditions = []gocrud.Condition{
		{"", "phone", "", "133******", "", "EQ", false},
	}
	sysAdmin := SysAdmin{ID: 0, Username: "demo0001"}
	formBody := gocrud_gorm.NewGormFormBody(body)
	a, err := formBody.Update(&sysAdmin, d, nil)
	if err != nil {
		fmt.Println("---------", err)
	} else {
		fmt.Println("00000000000", a.(*SysAdmin).ID)
	}
}

func TestFormBody_Save(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewFormBody("gorm", context.Background())
	sysAdmin := SysAdmin{Username: "demo0077400"}
	formBody := gocrud_gorm.NewGormFormBody(body)
	a, err := formBody.Save(&sysAdmin, d, nil)
	if err != nil {
		fmt.Println("---------", err)
	} else {
		fmt.Println("00000000000", a.(*SysAdmin).ID)
	}

}

func TestFormBody_Editor(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewEditorBody("gorm", context.Background())
	body.Conditions = []gocrud.Condition{
		{"phone", "", "", "werewfe", "", "EQ", true},
	}
	body.Editors = []gocrud.Editor{
		{"", "", "username", "AAA", ""},
	}
	editorBody := gocrud_gorm.NewGormEditorBody(body)
	err := editorBody.Updates(&SysAdmin{}, d, nil)
	if err != nil {
		fmt.Println("---------", err)
	} else {
		fmt.Println("00000000000")
	}

}

func TestFormBody_Delete(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewRemoveBody("gorm", context.Background())
	body.Conditions = []gocrud.Condition{
		{"phone", "", "", "werewfe", "", "EQ", true},
	}
	removeBody := gocrud_gorm.NewGormRemoveBody(body)
	err := removeBody.Delete(&SysAdmin{}, d, nil)
	if err != nil {
		fmt.Println("---------", err)
	} else {
		fmt.Println("00000000000")
	}

}

func TestFormBody_Page(t *testing.T) {
	gocrud_gorm.InitExec()
	d := db()
	body := gocrud.NewSearchBody("gorm", context.Background())
	body.Conditions = []gocrud.Condition{
		//{"phone", "", "", "werewfe", "", "EQ", true},
	}
	body.Page = 1
	searchBody := gocrud_gorm.NewGormSearchBody(body)
	tx, _ := searchBody.GormPaginate(d, nil)
	var list []SysAdmin
	tx.Find(&list)
	fmt.Printf("-----%+v", list)
}

type TestService struct {
	*gocrud.CommonResourceService
}

func (t TestService) ResourceName() string {
	return "test"
}

func (t TestService) Search(body gocrud.SearchBody) gocrud.SearchVO {
	gocrud_gorm.InitExec()
	d := db()
	sysAdmin := SysAdmin{}
	searchBody := gocrud_gorm.NewGormSearchBody(body)
	searchBody.GormQuery(d, nil).Where("phone = ?", "133*****031").First(&sysAdmin)
	fmt.Printf("==========%+v", sysAdmin)
	return gocrud.SearchVO{}
}

func TestTestService(t *testing.T) {
	tt := TestService{}
	body := gocrud.NewSearchBody("gorm", context.Background())
	tt.Search(body)
}

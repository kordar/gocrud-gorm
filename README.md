# gocrud-gorm

封装`gorm`类库，实现[`gocrud`](https://github.com/kordar/gocrud)接口。

## 安装

```go
go get github.com/kordar/gocrud-gorm v0.0.2
```

## 使用

- 初始化组件

```go
gocrud_gorm.InitExec()
```

- 创建调用相关实现类

```go
body := gocrud.NewFormBody("gorm", context.Background())
body.Object = map[string]interface{}{
    "phone":       "cccc234",
    "username":    "4444",
    "password":    "123456",
    "create_time": time.Now(),
    "update_time": time.Now(),
}
model, _ := body.Create(&SysAdmin{}, d, nil)
fmt.Printf("------------%v\n", model)
```

## `resource`实现类

```go
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
    var tx *gorm.DB = body.Query(d, nil)
    tx.Where("phone = ?", "1234").First(&sysAdmin)
    fmt.Printf("==========%+v", sysAdmin)
    return gocrud.SearchVO{}
}

func TestTestService(t *testing.T) {
    tt := TestService{}
    body := gocrud.NewSearchBody("gorm", context.Background())
    tt.Search(body)
}
```

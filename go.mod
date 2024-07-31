module github.com/kordar/gocrud-gorm

go 1.16

replace github.com/kordar/gocrud => ../gocrud

require (
	github.com/kordar/gocrud v0.0.0-00010101000000-000000000000
	github.com/spf13/cast v1.6.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
)

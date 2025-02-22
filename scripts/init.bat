@echo off

REM Initialize go module first
mkdir blog
cd blog
go mod init blog

REM Create application directories
mkdir config
mkdir router
mkdir model
mkdir service
mkdir util

REM Create pkg directories
mkdir pkg\database
mkdir pkg\logger
mkdir pkg\middleware

REM Create template directories
mkdir template\layout
mkdir template\post

REM Create static directories
mkdir static\css
mkdir static\js
mkdir static\img

REM Create data directories
mkdir data\db
mkdir data\upload
mkdir data\temp
mkdir data\log

REM Create config implementation
echo package config > config\config.go
echo. >> config\config.go
echo import ( >> config\config.go
echo     "encoding/json" >> config\config.go
echo     "os" >> config\config.go
echo ) >> config\config.go
echo. >> config\config.go
echo type Config struct { >> config\config.go
echo     AppName     string `json:"app_name"` >> config\config.go
echo     AppVersion  string `json:"app_version"` >> config\config.go
echo     ServerPort  string `json:"server_port"` >> config\config.go
echo     DbType     string `json:"db_type"` >> config\config.go
echo     DbPath     string `json:"db_path"` >> config\config.go
echo     TablePrefix string `json:"table_prefix"` >> config\config.go
echo } >> config\config.go
echo. >> config\config.go
echo var AppConfig Config >> config\config.go
echo. >> config\config.go
echo func Init() { >> config\config.go
echo     AppConfig = Config{ >> config\config.go
echo         AppName:    "Blog", >> config\config.go
echo         AppVersion: "1.0.0", >> config\config.go
echo         ServerPort: ":8080", >> config\config.go
echo         DbType:    "sqlite3", >> config\config.go
echo         DbPath:    "data/db/blog.db", >> config\config.go
echo     } >> config\config.go
echo } >> config\config.go

REM Create empty files
type nul > main.go
type nul > readme.txt

REM Create empty template files
type nul > template\layout\base.html
type nul > template\post\list.html
type nul > template\post\detail.html

REM Create empty static files
type nul > static\css\style.css
type nul > static\js\app.js

REM Create empty router file
type nul > router\router.go

REM Create empty database file
type nul > pkg\database\database.go

REM Install dependencies
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

echo Project structure created successfully!
pause 
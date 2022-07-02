# Med Backend
Search clinics in Poland

## Swagger LINK
https://app.swaggerhub.com/apis/Bronsun/Med/1.0.0 

#### PG Admin
- Check  PG Admin on [http://0.0.0.0:5050/browser/](http://0.0.0.0:5050/browser/)
- Login with Credential Email `admin@admin.com` Password `root`
- Connect Database Host as `postgres_db`, DB Username and Password as per `.env` set
- Note: if not configure `.env`, default Username `mamun` and password `123`

### Let's Build an API

1. [models](models) folder add a new file name `example_model.go`

```go
package models

import (
	"time"
)

type Example struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}
// TableName is Database Table Name of this model
func (e *Example) TableName() string {
	return "examples"
}
```
2. Add Model to [migration](pkg/database/migration.go)
```go
package database

import (
	"gin-boilerplate/models"
)
//Add list of model add for migrations
var migrationModels = []interface{}{&models.Example{}}
```
3. [controller](controllers) folder add a file `example_controller.go`
- Create API Endpoint 
- Use any syntax of GORM after `base.DB`, this is wrapper of `*gorm.DB`

```go
package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (base *Controller) CreateExample(ctx *gin.Context) {
	example := new(models.Example)

	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.DB.Create(&example).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &example)
}
```
4. [routers](routers) folder add a file `example.go`
```go
package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gin-boilerplate/controllers"
)


func TestRoutes(route *gin.Engine) {
	ctrl := controllers.Controller{DB: database.GetDB()}
	v1 := route.Group("/v1")
	v1.POST("/example/", ctrl.CreateExample)
}
```
5. Finally, register routes to [index.go](routers/index.go)
```go
package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	//Add All route
	TestRoutes(route)
}
```
- Congratulation, your new endpoint `0.0.0.0:8000/v1/example/`

### Deployment
#### Container Development Build
- Run `make build`

#### Container Production Build and Up
- Run `make production`


- [Server Config](pkg/config/server.go)
```go
func ServerConfig() string {
viper.SetDefault("server.host", "0.0.0.0")
viper.SetDefault("server.port", "8000")
appServer := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
return appServer
}
```
- [DB Config](pkg/config/db.go)
```go
func DbConfiguration() string {
	
dbname := viper.GetString("database.dbname")
username := viper.GetString("database.username")
password := viper.GetString("database.password")
host := viper.GetString("database.host")
port := viper.GetString("database.port")
sslMode := viper.GetString("database.ssl_mode")

dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
host, username, password, dbname, port, sslMode)
return dsn
}
```

### Useful Commands

- `make dev`: make dev for development work
- `make build`: make build container
- `make production`: docker production build and up
- `clean`: clean for all clear docker images

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)


package db

import (
	"github.com/go-xorm/xorm"
	"github.com/kanztu/goblog/pkg/config"
	"github.com/kanztu/goblog/pkg/model"
	"github.com/kanztu/goblog/pkg/server_context"
	"github.com/kanztu/goblog/pkg/utils"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSqliteDB(name string) *xorm.Engine {
	db, err := xorm.NewEngine("sqlite3", name)
	if err != nil {
		server_context.SrvCtx.Logger.Fatal(err)
	}

	tables := []interface{}{
		new(model.Blog),
		new(model.Tag),
		new(model.Content),
		new(model.PageCata),
	}
	for _, v := range tables {
		err = db.CreateTables(v)
		if err != nil {
			server_context.SrvCtx.Logger.Fatal(err)
		}
	}

	pages := []model.PageCata{
		{
			CataName: "Home",
			CataPath: "/",
		},
		{
			CataName: "About",
			CataPath: "/about",
		},
	}

	for _, v := range pages {
		_, err = db.Table(v.TableName()).Insert(&v)
		if err != nil {
			server_context.SrvCtx.Logger.Fatal(err)
		}
	}

	return db
}

func OpenSqliteDB(name string) *xorm.Engine {
	db, err := xorm.NewEngine("sqlite3", name)
	if err != nil {
		server_context.SrvCtx.Logger.Fatal(err)
	}
	return db
}

func InitDB() {
	dbFileName := config.CfgGlobal.DB.File
	exist, err := utils.Exists(dbFileName)
	if err != nil {
		server_context.SrvCtx.Logger.Fatal(err)
	}
	if exist {
		server_context.SrvCtx.DB = OpenSqliteDB(dbFileName)
	} else {
		server_context.SrvCtx.DB = CreateSqliteDB(dbFileName)
	}
	// server_context.SrvCtx.DB.ShowSQL(true)
}

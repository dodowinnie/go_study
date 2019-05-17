package models

import(
	"time"
	"github.com/astaxie/beego/orm"
	_"github.com/mattn/go-sqlite3"
	"github.com/Unknow/com"
	"os"
	"path"
)

const(
	_DB_NAME = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct{
	Id int64
	Title string
	Created time.Time `orm:"index"` // 建表时创建索引
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}

type Topic struct{
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"` // 建表时创建索引
	Upated time.Time `orm:"index"` 
	Views int64 `orm:"index"`
	Author string
	ReplyTime time.Time
	ReplyCount int64
	ReplyLastUserId int64

}

func RegisterDB(){
	// 判断数据库文件是否存在
	if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm) // 创建文件夹
		os.Create(_DB_NAME) // 创建数据库文件
	}

	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	// 注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// 注册数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)



}
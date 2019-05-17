package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Unknow/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"` // 建表时创建索引
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"` // 建表时创建索引
	Upated          time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
	CategoryId      int64 // 分类Id
}

func RegisterDB() {
	// 判断数据库文件是否存在
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm) // 创建文件夹
		os.Create(_DB_NAME)                          // 创建数据库文件
	}

	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	// 注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// 注册数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCatetory(name string) error {
	// 获取orm对象
	o := orm.NewOrm()

	cate := &Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		// 存在
		return err
	}

	// 插入操作
	_, err = o.Insert(cate)
	beego.Informational(cate)
	if err != nil {
		return err
	}
	return nil

}

func GetAllCategory() ([]*Category, error) {
	cates := make([]*Category, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func UpdateTopicCount(cid string, opt string) error {
	o := orm.NewOrm()
	cidInt, err := strconv.ParseInt(cid, 10, 64)
	cate := &Category{Id: cidInt}
	if err != nil {
		return err
	}
	err = o.Read(cate)
	if err != nil {
		return err
	}

	switch opt {
	case "1":
		cate.TopicCount += 1
	case "0":
		cate.TopicCount -= 1
		if cate.TopicCount < 0 {
			cate.TopicCount = 0
		}
	}
	_, err = o.Update(cate)
	if err != nil {
		return err
	}
	return nil

}

func DeleteCategory(id string) error {
	indInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: indInt}
	o := orm.NewOrm()
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}

	return nil
}

func AddTopic(title string, content string, categoryId string) error {
	idInt, err := strconv.ParseInt(categoryId, 10, 64)
	o := orm.NewOrm()
	topic := &Topic{
		Title:      title,
		Content:    content,
		Created:    time.Now(),
		Upated:     time.Now(),
		CategoryId: idInt,
	}
	_, err = o.Insert(topic)
	// 添加分类文章数
	err = UpdateTopicCount(categoryId, "1")
	if err != nil {
		return err
	}
	return nil
}

func UpdateTopic(title string, content string, id string) error {
	o := orm.NewOrm()
	idInt, err := strconv.ParseInt(id, 10, 64)
	topic := &Topic{Id: idInt}
	err = o.Read(topic)
	if err != nil {
		return err
	}
	topic.Content = content
	topic.Title = title
	topic.Upated = time.Now()
	_, err = o.Update(topic)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTopic(id string) error {
	o := orm.NewOrm()
	idInt, _ := strconv.ParseInt(id, 10, 64)
	topic := &Topic{Id: idInt}

	err := o.Read(topic)
	// 减分类文章数
	err = UpdateTopicCount(strconv.FormatInt(topic.CategoryId, 10), "0")
	_, err = o.Delete(topic)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	topics := make([]*Topic, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidInt, err := strconv.Atoi(tid)
	if err != nil {
		return nil, err
	}
	topic := &Topic{}
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidInt).One(topic)
	if err != nil {
		beego.Error(err)
	}
	topic.Views++
	topic.Upated = time.Now()
	_, err = o.Update(topic)
	return topic, err
}

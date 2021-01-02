package models

import "beego_blogweb/utils"

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

// 增加文章
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

// 插入文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title, tags, short, content, author, createtime) values(?, ?, ?, ?, ?, ?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

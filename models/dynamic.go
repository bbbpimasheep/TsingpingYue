package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dynamic struct {
	// type 0: 单纯转发一首古代的词。1：单纯转发一首现代的词（其实这个type也可以用来“公开发布自己的诗词”）。2：转发自己的收藏夹批注。
	Type int `bson:"type"`
	// 在type为0的时候，这个字段是对应的古代词的id。
	CiId primitive.ObjectID `bson:"ci_id"`
	// 在type为1的时候，这个字段是对应的现在作品的id。
	UserWorkId primitive.ObjectID `bson:"user_work_id"`
	// 在type为2的时候，这个字段是对应的收藏夹item的id。
	CollectionItemId primitive.ObjectID `bson:"collection_item_id"`
	// 这个动态下面的评论
	Comments primitive.ObjectID `bson:"comments"`
}
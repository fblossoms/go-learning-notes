package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users *mongo.Collection
var client *mongo.Client

func init() {
	var err error
	clientOpt := options.Client()
	uri := "mongodb://localhost:27017/admin"
	loggerOpt := options.Logger()
	loggerOpt.SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	clientOpt.ApplyURI(uri).SetLoggerOptions(loggerOpt)
	client, err = mongo.Connect(context.TODO(), clientOpt) // URI。懒连接
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(client)
	// 指定库、表
	db := client.Database("test")
	users = db.Collection("users") // 表table collection集合，相当于test.users。当插入数据时，才能刷新出这个表或库
}

// InsertOne 插入单条
func InsertOne() {
	feifei := User{Name: "Feifei", Age: 18}
	ior, err := users.InsertOne(context.TODO(), feifei)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ior.InsertedID)
}

// InsertMany 插入多条
func InsertMany() {
	minmin := User{Name: "Minmin", Age: 19}
	yanyan := User{Name: "Yanyan", Age: 20}
	imr, err := users.InsertMany(context.TODO(), []interface{}{minmin, yanyan})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(imr.InsertedIDs)
}

func InsertOne2() {
	yuyu := User{Name: "Yuyu", Age: 18}
	ior, err := users.InsertOne(context.TODO(), yuyu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ior.InsertedID)
}

func FindOne() {
	filter := bson.D{} // 查询条件
	sr := users.FindOne(context.TODO(), filter)
	fmt.Println(sr.Raw())
	fmt.Println(sr.Err()) // 查看是否有错误

	// 不指定类型，得出的数据不干净
	//var i any
	//err := sr.Decode(&i)
	//fmt.Println(err, "###")
	//fmt.Printf("%T, %[1]v", i)

	// 指定类型
	var u User
	err := sr.Decode(&u)
	fmt.Println(err, "###")
	fmt.Printf("%T, %[1]v", u)
}

// FindOne2 查询单条
func FindOne2() {
	filter := bson.D{}
	//filter = bson.D{bson.E{Key: "age", Value: 18}} // []E{E{Key:xx, Value:yy}}
	filter = bson.D{{"age", 19}}                // 简洁写法
	sr := users.FindOne(context.TODO(), filter) // 默认只查1条，LIMIT 1
	fmt.Println(sr.Raw())
	fmt.Println(sr.Err())
	var u User
	err := sr.Decode(&u)
	fmt.Println(err, "###")
	fmt.Printf("%T, %[1]v\n", u)
}

func (u User) String() string {
	return fmt.Sprintf("<User %v: %s, %d>", u.Id.Hex()[:24], u.Name, u.Age)
}

func FindMany() {
	filter := bson.D{}
	// cursor游标，指针（类似仪器的指针，不是指向地址的指针）
	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var us []*User
	for cursor.Next(context.TODO()) {
		//fmt.Println(cursor.Current)
		var u User
		err := cursor.Decode(&u)
		if err != nil {
			continue
		}
		us = append(us, &u)
	}
	fmt.Println(us)
}

func FindMany2() {
	filter := bson.D{}
	cursor, err := users.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var us []*User
	err = cursor.All(context.TODO(), &us)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us {
		fmt.Println(u)
	}
}

// FindByFilter 查询：复杂条件的处理
func FindByFilter(filter any) {
	FindAll(filter, nil)
}

func FindAll(filter any, opt *options.FindOptions) {
	cursor, err := users.Find(context.TODO(), filter, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var us []*User
	err = cursor.All(context.TODO(), &us)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us {
		fmt.Println(u)
	}
}

// UpdateOne 更新单条
func UpdateOne(filter bson.M, value bson.M) {
	result, err := users.UpdateOne(context.TODO(), filter, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.UpsertedID, result.UpsertedCount, result.ModifiedCount)
}

func UpdateMany(filter bson.M, value bson.M) {
	ur, err := users.UpdateMany(context.TODO(), filter, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ur.MatchedCount, ur.ModifiedCount, ur.UpsertedCount)
}

// Replace 替换
func Replace() {
	result, _ := users.ReplaceOne(context.TODO(), bson.M{"name": "Feifei"}, bson.M{"score": 88}) // id保留，所有字段清除，增加新的文档
	fmt.Println(result.MatchedCount, result.UpsertedCount, result.UpsertedID)
}

// DeleteOne 删除单条
func DeleteOne() {
	id, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	filter := bson.M{"_id": id}
	r, err := users.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.DeletedCount)
}

// DeleteMany 删除多条
func DeleteMany() {
	dr, err := users.DeleteMany(context.TODO(), bson.M{"score": bson.M{"$exists": true}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dr.DeletedCount)
}

type User struct {
	// 如果没有提供主键的话，会生成一个主键_id
	Id   primitive.ObjectID `bson:"_id,omitempty"` // 声明主键，12个字节，ID指定使用。声明了就必须写否则默认零值，会引起主键冲突
	Name string
	Age  int
}

func main() {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil { // 最后关闭
			log.Fatal(err)
		}
	}()
	//InsertOne()
	//InsertMany()
	//InsertOne2()

	//FindOne()
	//FindOne2()
	//FindMany()
	//FindMany2()

	// 有点繁琐的写法
	//filter := bson.D{}
	//filter = bson.D{{"age", 18}}                         // age == 18
	//filter = bson.D{{"name", bson.D{{"$eq", "Minmin"}}}} // $eq代表==
	//filter = bson.D{{"age", bson.D{{"$gt", 18}}}} 	   // age > 18
	// 符合习惯的写法
	//filter := bson.M{}
	//filter = bson.M{"name": "Feifei"}                                  // name == "Feifei"
	//filter = bson.M{"name": bson.M{"$eq": "Minmin"}}                   // name == "Minmin"
	//filter = bson.M{"age": bson.M{"$gt": 18}}                          // age > 18
	//filter = bson.M{"age": bson.M{"$in": []int{17, 18, 19}}}           // age IN [17, 18, 19]
	//filter = bson.M{"age": bson.M{"$nin": []int{17, 18, 19}}}          // age NOT IN [17, 18, 19]
	//filter = bson.M{"name": "Feifei", "age": 18}                       // name == "Feifei" AND age == 18
	//filter = bson.M{"$and": []bson.M{{"name": "Minmin"}, {"age": 19}}} // name == "Minmin" AND age == 19
	//filter = bson.M{"$and": []bson.M{bson.M{"name": "Minmin"}, bson.M{"age": 19}}}
	//filter = bson.M{"$and": []bson.M{{"name": "Yanyan"}, {"age": bson.M{"$gt": 18}}}} // name == "Yanyan" AND age > 18
	//filter = bson.M{"$or": []bson.M{{"name": "Yanyan"}, {"age": bson.M{"$gt": 18}}}}  // name == "Yanyan" OR age > 18
	//filter = bson.M{"age": bson.M{"$not": bson.M{"$eq": 18}}}                         // age NOT 18
	//
	//filter = bson.M{"gender": bson.M{"$exists": false}} // false：不存在gender字段
	//filter = bson.M{"age": bson.M{"$type": []int{18, 19, 20}}} // 数据类型为int
	//FindByFilter(filter)
	//opt := options.Find()
	//opt.SetSort(bson.M{"age": -1}).SetLimit(3).SetSkip(4) // 按 age字段降序（-1）排序（不支持多字段排序，1为以升序排序）。限制返回3个，跳过4个，常用于分页
	//opt.SetProjection(bson.M{"name": false}) // false遮蔽字段
	//FindAll(filter, opt)

	//filter := bson.M{}
	//id, _ := primitive.ObjectIDFromHex("000000000000000000000000") // 找到该对应实例
	//filter = bson.M{"_id": id}
	//value := bson.M{"$set": bson.M{"name": "Nannan", "gender": 0, "score": 60.6}} // 修改对应实例

	// 更新单条
	//filter = bson.M{"name": "Feifei"}
	//value := bson.M{"$inc": bson.M{"age": +2}} // 自增
	//UpdateOne(filter, value)

	// 更新多条
	//filter = bson.M{"age": 20}
	//value := bson.M{"$inc": bson.M{"age": +2}} // 自增
	//UpdateMany(filter, value)

	// 更新ID
	//id, _ := primitive.ObjectIDFromHex("000000000000000000000000") // 找到该对应实例
	//users.UpdateByID(context.TODO(), id, bson.M{"$set": bson.M{"name": "Yunyun"}})

	// 替换
	//Replace()

	// 删除单条
	//DeleteOne()

	// 删除多条
	DeleteMany()
	//fmt.Println(users)
}

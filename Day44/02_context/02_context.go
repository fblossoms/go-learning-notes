package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users *mongo.Collection
var client *mongo.Client

func init() {
	// 表table collection集合，相当于test.users。当插入数据时，才能刷新出这个表或库
}

func main() {
	// 字段
	//ctx := context.WithValue(context.Background(), "name", "Feifei")
	//fmt.Println(ctx.Value("name"))
	//fmt.Println(ctx.Value("age"))

	// 父子
	//parent := context.Background()
	//c1 := context.WithValue(parent, "k1", "v1")
	//c2 := context.WithValue(parent, "k2", "v2")
	//c3 := context.WithValue(c1, "k3", "v3")
	//c4 := context.WithValue(c3, "k4", "v4")
	//names := []string{"k1", "k2", "k3", "k4"}
	//for _, name := range names { // c4 -> c3 -> c1 -> parent(祖先) <- c2
	//	fmt.Println(name, c1.Value(name), c2.Value(name), c3.Value(name), c4.Value(name))
	//}

	// 应用：父子协程控制
	//parent := context.Background()
	//ctx, cancel := context.WithTimeout(parent, 10*time.Second)
	//defer cancel()
	//go func(c context.Context) {
	//	tick := time.NewTicker(time.Second)
	//	for {
	//		select {
	//		case <-tick.C: // 工作通道，有数据
	//			fmt.Println("每一秒做一次工作内容")
	//		case <-c.Done():
	//			fmt.Println("收到了父协程中ctx的结束信号，我子协程退出")
	//			return // 当前子协程结束
	//		}
	//	}
	//}(ctx)
	//fmt.Println("-------------- 1 子协程结束 ----------------")
	//
	//// <- ctx.Done()
	//select {
	//case <-ctx.Done():
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("主协程即将结束")
	//}
	//fmt.Println("-------------- 2 父协程结束 ----------------")

	// 应用：连接超时取消
	parent := context.Background()
	timeoutCtx, cancel := context.WithTimeout(parent, 2*time.Second)
	defer cancel()

	var err error
	clientOpt := options.Client()
	uri := "mongodb://localhost:27018/"

	clientOpt.ApplyURI(uri)
	client, err = mongo.Connect(context.TODO(), clientOpt) // URI。懒连接
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(timeoutCtx, nil)
	if err != nil {
		fmt.Println(timeoutCtx.Err().Error())
		fmt.Println(timeoutCtx.Err() == context.DeadlineExceeded) // 生产中常用，后续进行if或switch检查，若为true，返回提示给用户
		fmt.Printf("%T\n", err)
		log.Fatal(err)
	}
	fmt.Println(client)
	// 指定库、表
	db := client.Database("test")
	users = db.Collection("users")
}

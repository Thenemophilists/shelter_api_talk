package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	apiCors "shelter_api/cors"
	db "shelter_api/db" // 导入 db 包
	errPage "shelter_api/err_page"
	apiFavicon "shelter_api/favicon"
	"strconv"
)

// MyDocument 结构体定义数据库集合中文档的格式
type MyDocument struct {
	Id     int32  `bson:"id"`       // 根据数据库中的实际字段名设置
	Text   string `bson:"hitokoto"` // 根据数据库中的实际字段名设置
	From   string `bson:"from_who"` // 根据数据库中的实际字段名设置
	Config struct {
		UsingDB struct {
			Name    string
			Version string
		}
		Api struct {
			Version string
		}
	}
}

var debug bool

var hitokotoCollection *mongo.Collection

// main 函数是程序的入口点
func main() {
	viper := viper.New()
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")    // 读取配置文件路径
	err := viper.ReadInConfig() // 读取配置文件
	if err != nil {
		fmt.Println("[Error] Can't read config file")
		os.Exit(2)
	}
	debug = viper.GetBool("debug")               // 获取是否开启调试模式
	usingPort := viper.GetInt("port")             // 获取端口号
	faviconPath := viper.GetString("faviconPath") // 获取favicon路径
	database := viper.Sub("database")
	hitokotoDB := database.Sub("hitokotoDB")
	hitokotoDBUri := hitokotoDB.GetString("uri")                   // 获取数据库URI
	hitokotoDBName := hitokotoDB.GetString("dbName")               // 获取数据库名称
	hitokotoCollectionName := hitokotoDB.GetString("dbCollection") // 获取集合名称

	// 连接到MongoDB
	client := db.ConnectMongoDB(hitokotoDBUri)
	defer client.Disconnect(context.TODO()) // 确保在程序结束前关闭连接

	// 选择数据库和集合
	hitokotoDatabase := client.Database(hitokotoDBName)
	hitokotoCollection = hitokotoDatabase.Collection(hitokotoCollectionName)

	if debug {
		// 开启调试模式
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode) // 设置运行模式
	}
	api := gin.Default()              // 初始化Gin引擎
	api.Use(apiCors.CORSMiddleware()) // 使用中间件解决CORS跨域
	api.Use(apiFavicon.Set(faviconPath))
	api.Handle("GET", "/sentence", GetText) // 注册文案路由
	api.NoRoute(errPage.Return404Page)      // 注册404返回

	if err := api.Run(":" + strconv.Itoa(usingPort)); err != nil {
		fmt.Println("[Error] Can't launch Shelter API")
	}

}

// GetText 函数处理获取文本的请求
func GetText(c *gin.Context) {
	// 获取参数
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}
	var pipeline mongo.Pipeline
	if id == -1 {
		// 使用聚合随机返回一条数据         
		pipeline = mongo.Pipeline{
		{{"$sample", bson.D{{Key: "size", Value: 1}}}}, // 从集合中随机抽取一条文档
	}
	} else {
		pipeline = mongo.Pipeline{
			{{"$match", bson.M{"id": id}}}, // 从集合中抽取一条文档
		}
	}

	cursor, err := hitokotoCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	defer cursor.Close(context.TODO())

	// 获取文档
	var result MyDocument
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}

	result.Config.UsingDB.Name = "hitokotoDB"
	result.Config.UsingDB.Version = "0.0.1.000001"
	result.Config.API.Version = "0.0.1.000001"
	if debug {
		fmt.Println(result)
	}

	c.JSON(http.StatusOK, result)

}

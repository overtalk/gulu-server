package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gohouse/converter"
)

var (
	username string
	password string
	host     string
	db       string
	path     string
	port     int
	// tables name
	tables = []string{`player`, `nodes`, `openid_mapping`, `player_global_mail`, `player_invitation`, `pvp_result`, `single_mail`}
)

func main() {
	flag.StringVar(&host, "host", "localhost", "mysql host ( default `localhost` )")
	flag.StringVar(&username, "u", "root", "mysql username ( default `root` )")
	flag.StringVar(&password, "p", "root", "mysql password ( default `root` )")
	flag.StringVar(&db, "d", "sausage_shooter", "mysql database ( default `sausage_shooter` )")
	flag.StringVar(&path, "path", "/", "project path ( default `/` )")
	flag.IntVar(&port, "port", 3306, "mysql port ( default 3306 )")
	flag.Parse()

	// dns
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, db)

	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
	})

	for _, table := range tables {
		fmt.Printf("convert table[%s]...\n", table)
		filePath := fmt.Sprintf("%s/model/%s.go", path, table)
		if err := t2t.
			// 指定某个表,如果不指定,则默认全部表都迁移
			Table(table).
			// 表前缀
			//Prefix("prefix_").
			// 是否添加json tag
			//EnableJsonTag(true).
			// 生成struct的包名(默认为空的话, 则取名为: package model)
			PackageName("model").
			// tag字段的key值,默认是orm
			TagKey("json").
			// 是否添加结构体方法获取表名
			RealNameMethod("TableName").
			// 生成的结构体保存路径
			SavePath(filePath).
			// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
			Dsn(dns).
			// 执行
			Run(); err != nil {
			log.Printf("failed to convert table[%s], error = %v", table, err)
		}
	}
}

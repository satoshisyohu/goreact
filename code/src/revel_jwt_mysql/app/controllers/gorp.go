package controllers

import (
       	"database/sql"
       	"github.com/go-gorp/gorp"
       	_ "github.com/go-sql-driver/mysql"
       	r "github.com/revel/revel"
       	"github.com/revel/modules/db/app"
       	"revel_jwt_mysql/app/models"
       	"github.com/revel/revel"
    "time"
       	"io"
       	"os"
       	"log"
       	"fmt"
       	"encoding/csv"
       	"path/filepath"
)

var (
       	Dbm *gorp.DbMap
)

func failOnError(err error){
       	if err != nil {
       		log.Fatal("Error", err)
       	}
}

//DB の初期設定
func InitDB() {
       	db.Init()
       	// conf ファイルのに記載されているDB名を読み込み処理を行なう
       	// `revel revel_jwt_mysql run`と`revel revel_jwt_mysql run prod`で読み込むDBが異なる`prod`が本番用
       	uri := read_conf("db.uri")

       	// mysqlへの接続処理
       	db_access_uri := "revel_jwt_mysqlapp:revel_jwt_mysqlapp@tcp(" + uri +")/revel_jwt_mysqlapp"
       	db, err := sql.Open("mysql", db_access_uri)
       	if err != nil {
       		panic(err)
       	}
       	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

       	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
       		for col, size := range colSizes {
       			t.ColMap(col).MaxSize = size
       		}
       	}

       	// Tableに設定する値の長さを設定している
       	t := Dbm.AddTable(models.Revel_JWT_MySQL{}).SetKeys(true, "Revel_JWT_MySQL_Id")
       	setColumnSizes(t, map[string]int{
       		"Date": 10,
       		"Title": 40,
       		"Picture": 140,
       		"Picture_Link": 140,
       	})

       	Dbm.TraceOn("[gorp]", r.INFO)
       	Dbm.CreateTables()

       	// 時間情報のフォーマットを設定している
       	t_value := time.Now().Format("2006-01-02")

       	prev, err := filepath.Abs(".")
       	if err != nil {
       		defer fmt.Println("Error")
       	}
       	// Debug
       	revel.TRACE.Printf("check the Directory for the csv file")
       	revel.TRACE.Printf("%s", prev)

       	// 最初に登録するcsv形式のデータ
       	file, err := os.Open(prev + "/controllers/csv_data/revel_jwt_mysql.csv")
       	failOnError(err)
       	defer file.Close()

       	// 画像データのパスを作成するために取得
       	http_uri := read_conf("http.uri")
       	http_port := read_conf("http.port")
       	img_path := http_uri + ":" + http_port + "/public/img/"

       	f_count := 0
       	reader := csv.NewReader(file)
       	var revel_jwt_mysqls []*models.Revel_JWT_MySQL

       	// DBへの登録
       	for {
       		recoad, err := reader.Read()
       		if err == io.EOF {
       			break
       		} else {
       			failOnError(err)
       		}
       		img_data := img_path + recoad[2]
       		revel_jwt_mysqls = append(revel_jwt_mysqls, &models.Revel_JWT_MySQL{f_count, t_value, recoad[0], img_data,
       			                                                          recoad[1]})
       		f_count++
    }
       	for _, revel_jwt_mysql := range revel_jwt_mysqls {
       		if err := Dbm.Insert(revel_jwt_mysql); err != nil {
       			panic(err)
       		}
       	}
}

type GorpController struct {
       	*r.Controller
       	Txn *gorp.Transaction
}

func read_conf(conf_param string) (string){
       	var uri string = ""

       	p, found := revel.Config.String(conf_param)

       	if !found {
       		if uri == "" {
       			revel.ERROR.Fatal("Cound not find")
       		}
       	}

       	return p
}

func (c *GorpController) Begin() r.Result {
       	txn, err := Dbm.Begin()
       	if err != nil {
       		panic(err)
       	}
       	c.Txn = txn
       	return nil
}

func (c *GorpController) Commit() r.Result {
       	if c.Txn == nil {
       		return nil
       	}
       	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
       		panic(err)
       	}
       	c.Txn = nil
       	return nil
}

func (c *GorpController) Rollback() r.Result {
       	if c.Txn == nil {
       		return nil
       	}
       	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
       		panic(err)
       	}
       	c.Txn = nil
       	return nil
}

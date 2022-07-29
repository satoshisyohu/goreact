package controller

import (
	"fmt"
	"go_mvc/my-app/src/model"
	"go_mvc/my-app/src/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowAllBlog(c *gin.Context) {
	datas := model.GetAll()
	c.HTML(200, "alldata.html", gin.H{
		"datas": datas,
	})
}

func ShowOneData(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.JSON(200, gin.H{
		"id":     data.Id,
		"日付":     data.Brew_day,
		"ドリッパー":  data.Device,
		"挽き目":    data.Grind,
		"豆の種類":   data.Beans,
		"購入したお店": data.Shop,
		"評価":     data.Review,
		"その他":    data.Other,
	})
}

func CreateBlog(c *gin.Context) {
	id := c.PostForm(("id"))
	brew_day := c.PostForm("brew_day")
	device := c.PostForm("device")
	grind := c.PostForm("grind")
	beans := c.PostForm("beans")
	shop := c.PostForm("shop")
	review := c.PostForm("review")
	other := c.PostForm("other")

	data := model.Coffee_recipe{Id: id, Brew_day: brew_day, Device: device, Grind: grind, Beans: beans,
		Shop: shop, Review: review, Other: other}

	data.Create()
	c.Redirect(301, "/")
}

func Caliculatehome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		// htmlに渡す変数を定義
	})
}
func New(c *gin.Context) {
	content := c.PostForm("gram")
	gram, _ := strconv.Atoi(content)
	sum := service.GetGram(gram)

	c.HTML(200, "index.html", gin.H{
		"beans_quantity":      gram,
		"boil_water_quantity": sum,
	})
}

//ユーザー登録関連

func SignUpHome(c *gin.Context) {

	c.HTML(200, "signup.html", gin.H{})
}

func SignUp(c *gin.Context) {
	var form model.User

	err := c.Bind(&form)
	if err != nil {
		c.HTML(http.StatusBadRequest, "signup.htmp", gin.H{
			"err": err,
		})
		c.Abort()
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		err := model.CreateUser(username, password)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{
				"err": "すでにユーザ名が登録されています",
			})
			c.Redirect(302, "/")
		}
	}
}

func LoginHome(c *gin.Context) {

	c.HTML(200, "login.html", gin.H{})
}

func Login(c *gin.Context) {
	fmt.Println("can you see this word?")
	log.Println(c.PostForm("username"))
	log.Println(c.PostForm("password"))

	dbPassword := model.GetUser(c.PostForm("username")).Password
	log.Println(dbPassword)
	fmt.Println(dbPassword)

	formPassword := c.PostForm("password")

	err := service.CompareHashAndPassword(dbPassword, formPassword)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"err": "ユーザ名とパスワードが間違っえています",
		})
		c.Abort()
	} else {
		session := sessions.Default(c)
		session.Set("username", c.PostForm("usernmae"))
		session.Save()
		log.Println("Login Ok!")
		c.HTML(200, "loginok.html", gin.H{})
		// c.Redirect(302, "/")
	}

}

func Beans(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})

}

func Register(c *gin.Context) {
	dbcountry := c.PostForm("beans")
	dbquantity := c.PostForm("quantity")

	result := model.RegisterBeans(dbcountry, dbquantity)
	if result != nil {
		log.Println(result)
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"err": "豆登録時にエラーが発生しました。",
		})
		c.Redirect(302, "/")

	}

}

func GetAllData(c *gin.Context) {
	var datas = model.GetAllData()
	for i := 0; i < len(datas); i++ {
		var data = "Data" + strconv.Itoa(i+1)
		c.JSON(200,
			gin.H{
				data: datas[i],
			})
	}

}

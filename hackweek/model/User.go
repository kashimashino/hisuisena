package model

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kashimashino/hackweek/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"log"
	"net/http"
	"strconv"
	"time"

	"math/rand"
)


type Clothes struct {
	Owner    		string    	`json:"owner" gorm:"type:varchar(20);primary_key"`
	Avatar   		string  	`json:"avatar" gorm:"type:varchar(50)"`
	Tittle   		string     	`json:"tittle" gorm:"type:varchar(20)"`
	Likes    		int       	`json:"likes" gorm:"type:int"`
	ImageUrl 		string		`json:"image_url" gorm:"type:varchar(50)"`
	Cate     		int			`json:"cate" gorm:"type:int"`
}

type MakeActivities struct {
	Cate 	 		int  		`json:"cate" gorm:"primary_key"`
	Tittle 	 		string		`json:"tittle"`
	Info     		string      `json:"info"`
	ImageUrl 		string 		`json:"image_url"`
	Likes    		int  		`json:"likes"`
	Username 		string 		`json:"username"`
}


type User struct {
	gorm.Model
	Username 		string 		`gorm:"type:varchar(20);not null" json:"username"`
	Password 		string 		`gorm:"type:varchar(20);not null" json:"password"`
	Sex      		string		`json:"sex"`
	Info     		string  	`json:"info"`
	BirthDay 		string  	`json:"birth_day"`
	Location		string   	`json:"location"`
	Role     		int    		`gorm:"type:int" json:"role"` //0为管理员，1为一般用户
	Activities  	string 		`json:"action"`
}



//数据库操作

//搜索
func Search(form string) []Clothes {
	var clothes []Clothes
	form = "%"+form+"%"
	err := db.Where("tittle Like ?",form).Find(&clothes).Error
	if err != nil {
		return nil
	}
	if form == "%%" || form == "% %" {
		return nil
	}
	return clothes
}



//新增用户

//检查用户
func(users User) CheckUser() (code int) {
	db.Select("id").Where("username=?", users.Username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //用户名已存在  1001
	}
	return errmsg.SUCCSE
}

//创建用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

//查询用户列表   可分页
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

//编辑用户信息
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})

	maps["owner"]= data.Username
	maps["location"]=data.Location
	maps["info"] =data.Info
	maps["sex"] =data.Sex
	//maps["password"] = BcryptPw(data.Password)
	db.Model(&User{}).Where("id=?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//发布动态
func MakeActivity(data *MakeActivities) int {
	activity := MakeActivities{
		Username: data.Username,
		Cate:     data.Cate,
		Tittle:   data.Tittle,
		Info:     data.Info,
		ImageUrl: data.ImageUrl,
		Likes:    data.Likes,
	}
	err := db.Create(&activity).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}
//查看动态
func CheckActivities(data *MakeActivities)[]MakeActivities  {
	var activities []MakeActivities
	err = db.Where("username=?",data.Username).Find(&activities).Error
	if err != nil {
		return nil
	}
	return activities
}
//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//密码加密   scrypt
func ScryptPw(password string) string {

	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 21}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
//        bcrypt密码加密，觉得scrypt更好就写了scrypt
//func BcryptPw(password string) string  {
//	HashPw,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
//	if err != nil {
//	log.Fatal(err)
//}
//	fpw := base64.StdEncoding.EncodeToString(HashPw)
//	return fpw
//}

//登录验证
func CheckLogin(username string, password string) (code int,uid uint) {
	var user User

	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST,0
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG,0
	}
	//if bcrypt.CompareHashAndPassword(HashPw,[]byte(password)) {
	//	return errmsg.ERROR_PASSWORD_WRONG
	//}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGHT,0
	}
	return errmsg.SUCCSE,user.ID
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end - start) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}


//功能
func CheckClothes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var totalData []Clothes
	var data Clothes
 	a:= generateRandomNumber(0,23,8)
	for i := 0; i <= 7; i++ {
		db.Where("cate=?", id).Offset(a[i]).Limit(1).Find(&data)
		  totalData= append(totalData, data)
	}

		c.JSON(http.StatusOK, gin.H{
			"data":   totalData,
			"status": "OK",
			"msg":    errmsg.GetErrMsg(200),
		})

}


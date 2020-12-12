package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek/model"
	"github.com/kashimashino/hackweek/server"
	"github.com/kashimashino/hackweek/utils/errmsg"
	"net/http"
	"strconv"
)



var code int
//注册用户
func AddUser(c *gin.Context)  {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = data.CheckUser()
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"msg" : errmsg.GetErrMsg(code),
	})
}
//搜索
func Search(c *gin.Context)  {
	var form model.Clothes
	_ =c.ShouldBindJSON(&form)
	data := model.Search(form.Tittle)
	c.JSON(http.StatusOK,gin.H{
		"data" : data,
		"msg"  : errmsg.SUCCSE,
		"status" : "ok",
	})
}

//得到用户信息
func GetUser(c *gin.Context)  {
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenumber"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data := model.GetUsers(pageSize,pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"status" : "OK",
		"data" : data,
		"msg" : errmsg.GetErrMsg(code),
	})
}
// 编辑用户
func EditUser(c *gin.Context)  {
	var data model.User
	id ,_ :=strconv.Atoi(c.Param("id"))
	_=c.ShouldBindJSON(&data)
	model.EditUser(id,&data)

	c.JSON(http.StatusOK,gin.H{
		"status" : "Ok",
		"msg" : "message changes successful",
	})
}

//发布动态
func MakeActivity(c *gin.Context)  {
	var  data model.MakeActivities
	_=c.ShouldBindJSON(&data)
	code = model.MakeActivity(&data)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"msg" : errmsg.GetErrMsg(code),
	})
}

//查看动态
func CheckActivities(c *gin.Context)  {
	var data model.MakeActivities
	_=c.ShouldBindJSON(&data)
	showData := model.CheckActivities(&data)
	c.JSON(http.StatusOK,gin.H{
		"data" : showData,
		"status" : "OK",
		"msg" : errmsg.GetErrMsg(500),
	})
}

//删除用户
func DeleteUser(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"msg" : errmsg.GetErrMsg(code),
	})
}
//修改密码
func EditPassword(c *gin.Context)  {
	var data server.EditPwd
	id ,_ :=strconv.Atoi(c.Param("id"))
	_=c.ShouldBindJSON(&data)
	code = server.CheckPassword(&data)

	if code == errmsg.SUCCSE {
		server.EditPassword(id, &data)
	}
		c.JSON(http.StatusOK,gin.H{
			"status" : code,
			"msg" : errmsg.GetErrMsg(code),
		})
}

//
////处理OSS错误
//func handleError(err error) {
//	fmt.Println("Error",err)
//	os.Exit(-1)
//}
//
//
////上传头像
//func Upload(c *gin.Context)  {
//	// 从请求中读取文件
//	f, err := c.FormFile("f1")  // 从请求中获取携带的参数一样的
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error(),
//		})
//	}else {
//		// 将读取到的文件保存在本地（服务端本地）
//		dst := path.Join("./images/", f.Filename)
//	  _  =	c.SaveUploadedFile(f, dst)
//		c.JSON(http.StatusOK, gin.H{
//			"status": "OK",
//			"url":"http://hellosun.net.cn/hackweek/"+f.Filename,
//		})
//	}
//	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
//	accessKeyId := "LTAI4G9CDETryAM9M33stee8"
//	accessKeySecret := "QIlVgwI5n8H1SM0433IwljLoVxxKou"
//	bucketName := "hack-week"
//	objectName := "hackweek/"+f.Filename
//	localFileName := "./images/"+f.Filename
//
//	// 创建OSSClient实例。
//	client,err:= oss.New(endpoint,accessKeyId,accessKeySecret)
//	if err != nil {
//		handleError(err)
//	}
//	//获取储存空间
//	bucket,err := client.Bucket(bucketName)
//	if err  != nil {
//		handleError(err)
//	}
//	//上传文件
//	err = bucket.PutObjectFromFile(objectName,localFileName)
//	if err  != nil {
//		handleError(err)
//	}
//}
//


////上传图片
//func UploadPicture(c *gin.Context)  {
//	var todo model.MakeActivities
//	_=c.BindJSON(&todo)
//	err := model.GetDB().Create(&todo).Error
//	if err != nil {
//		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
//	}else{
//		c.JSON(http.StatusOK, gin.H{
//			"data": todo,
//			"code": 1,
//			"msg": "添加成功",
//		})
//	}
//}
//
//


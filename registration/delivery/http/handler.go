package http

import (
	"net/http"
	"registration/models"
	"registration/registration"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase registration.Usecase
}

func NewHandler(r *gin.Engine, a registration.Usecase) {
	handler := &Handler{a}

	r.POST("/class", handler.NewClass)
	r.GET("/class/all", handler.GetClassList)
	r.POST("/student/import", handler.ImportStudentList)
	r.PUT("/student/check-in", handler.CheckIn)
	r.GET("/class/status", handler.GetClassStatus)
	r.DELETE("/class", handler.DeleteClass)
}

func returnOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  data,
	})
}

func returnError(c *gin.Context, errMsg string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error":  true,
		"errMsg": errMsg,
	})
}

func (h *Handler) NewClass(c *gin.Context) {
	// 取post form ok
	var class models.Class
	err := c.BindJSON(&class)
	if err != nil {
		returnError(c, err.Error())
		return
	}
	if class.ClassName == "" && class.Date == "" {
		returnError(c, "不能為空")
		return
	}
	err = h.Usecase.NewClass(&class)
	if err != nil {
		returnError(c, err.Error())
	} else {
		returnOk(c, "ok")
	}

}

func (h *Handler) GetClassList(c *gin.Context) {
	classList, err := h.Usecase.GetClassList()
	if err != nil {
		returnError(c, err.Error())
	} else {
		returnOk(c, classList)
	}
}

type StudentList struct {
	StudentList []models.Student `json:"studentList"`
}

func (h *Handler) ImportStudentList(c *gin.Context) {
	sClassID := c.Query("classID")
	if sClassID == "" {
		returnError(c, "請回到首頁")
		return
	}
	classID, _ := strconv.Atoi(sClassID)

	var studentList []models.Student
	err := c.BindJSON(&studentList)
	if err != nil {
		returnError(c, err.Error())
		return
	}
	err = h.Usecase.ImportStudentList(classID, studentList)
	if err != nil {
		returnError(c, err.Error())
	} else {
		returnOk(c, "ok")
	}

	// return
}

type Student struct {
	EmployeeID string `json:"employeeID"`
}

func (h *Handler) CheckIn(c *gin.Context) {
	sClassID := c.Query("classID")
	if sClassID == "" {
		returnError(c, "請回到首頁")
		return
	}
	classID, _ := strconv.Atoi(sClassID)

	var student Student
	err := c.BindJSON(&student)
	if err != nil {
		returnError(c, err.Error())
		return
	}
	employeeID, _ := strconv.Atoi(student.EmployeeID)

	msg, err := h.Usecase.CheckIn(classID, employeeID)
	if err != nil {
		returnError(c, err.Error())
	} else {
		returnOk(c, msg)
	}

	// return
}

func (h *Handler) GetClassStatus(c *gin.Context) {
	// 取得get的 classID
	sClassID := c.Query("classID")
	if sClassID == "" {
		returnError(c, "請回到首頁")
		return
	}
	classID, _ := strconv.Atoi(sClassID)
	// 取redis studentList用classID取得
	// return
	studentList, err := h.Usecase.GetStudentList(classID)
	if err != nil {
		returnError(c, err.Error())
	} else {
		returnOk(c, studentList)
	}
}

func (h *Handler) DeleteClass(c *gin.Context) {
	// 取得get的 classID
	sClassID := c.Query("classID")
	if sClassID == "" {
		returnError(c, "請回到首頁")
		return
	}
	classID, _ := strconv.Atoi(sClassID)
	// return
	err := h.Usecase.DeleteClass(classID)
	if err != nil {
		returnError(c, err.Error())
	} else {
		returnOk(c, "ok")
	}
}

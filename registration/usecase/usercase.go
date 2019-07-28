package usecase

import (
	"encoding/json"
	"errors"
	"log"
	"registration/models"
	"registration/registration"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type redisUsecase struct {
	Conn *redis.Pool
}

func NewRedisUsecase(Conn *redis.Pool) registration.Usecase {
	return &redisUsecase{Conn}
}

func (r *redisUsecase) NewClass(class *models.Class) (err error) {
	// 取得流水號 usecase ok
	// 流水號塞到class struct裡 ok
	// format json ok
	// 塞入redis的classList 用zset就好 ok
	conn := r.Conn.Get()
	defer conn.Close()

	classID, err := redis.Int(conn.Do("INCR", "ClassID_Increment"))
	if err != nil {
		return
	}
	if classID == 0 {
		return errors.New("新增失敗")
	}

	class.ClassID = classID

	classJson, err := json.Marshal(class)
	if err != nil {
		return
	}

	_, err = conn.Do("ZADD", "classList", classID, string(classJson))

	return
}

func (r *redisUsecase) GetClassList() (classList []models.Class, err error) {
	// 從redis classList裡 取出
	// format json
	conn := r.Conn.Get()
	defer conn.Close()

	reply, err := redis.Values(conn.Do("ZRANGE", "classList", 0, -1))
	if err != nil {
		return classList, err
	}
	for _, value := range reply {
		valueByte, _ := value.([]byte)
		var class models.Class
		_ = json.Unmarshal(valueByte, &class)
		classList = append(classList, class)
	}

	return
}

func (r *redisUsecase) ImportStudentList(classID int, studentList []models.Student) (err error) {
	// 取得post form要有classID跟 studentlist
	// format json
	// 用hash 存 studentList，key值為classID，value為studentLIst json
	conn := r.Conn.Get()
	defer conn.Close()

	sClassID := strconv.Itoa(classID)

	redisKey := "studentList:" + string(sClassID)

	for _, student := range studentList {
		studentJson, err := json.Marshal(student)
		if err != nil {
			return err
		}
		_, err = conn.Do("HSET", redisKey, student.EmployeeID, string(studentJson))
	}

	return
}

func (r *redisUsecase) CheckIn(classID, employeeID int) (msg string, err error) {
	// 取得put form要有classID跟EmployeeID ok
	// 取redis studentList用classID取得 ok
	// format json ok
	// 更改classID的status為true ok
	// format json ok
	// 存回studentList ok

	conn := r.Conn.Get()
	defer conn.Close()

	sClassID := strconv.Itoa(classID)

	redisKey := "studentList:" + sClassID
	reply, err := redis.String(conn.Do("HGET", redisKey, employeeID))
	if err != nil {
		log.Println(err)
		return "沒有這個員工編號", nil
	}

	var student models.Student
	_ = json.Unmarshal([]byte(reply), &student)

	if student.Status {
		return "重複報到", nil
	}

	student.Status = true

	studentJson, err := json.Marshal(student)
	if err != nil {
		return
	}

	sEmployeeID := strconv.Itoa(employeeID)

	// 存回去
	_, err = conn.Do("HSET", redisKey, sEmployeeID, string(studentJson))

	return "報到成功", err
}

func (r *redisUsecase) GetStudentList(classID int) (studentList []models.Student, err error) {
	// 從redis classList裡 取出
	// format json
	conn := r.Conn.Get()
	defer conn.Close()

	sClassID := strconv.Itoa(classID)
	redisKey := "studentList:" + sClassID

	reply, err := redis.StringMap(conn.Do("HGETALL", redisKey))
	if err != nil {
		return
	}

	for _, value := range reply {
		var student models.Student
		_ = json.Unmarshal([]byte(value), &student)
		studentList = append(studentList, student)
	}

	return
}

func (r *redisUsecase) DeleteClass(classID int) (err error) {
	// 從redis classList裡 取出
	conn := r.Conn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("ZREMRANGEBYSCORE", "classList", classID, classID))
	if err != nil {
		return
	}
	if reply != 1 {
		return errors.New("classID錯誤")
	}

	sClassID := strconv.Itoa(classID)
	redisKey := "studentList:" + sClassID

	_, err = redis.Int(conn.Do("DEL", redisKey))
	if err != nil {
		return
	}

	return
}

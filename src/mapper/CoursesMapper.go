package mapper

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shijting/go-micro/src/boot"
	"github.com/shijting/go-micro/src/vars"
)

const (
	Field = "course_id, course_name, course_disp_name, course_price, course_price2, addtime"
)

func GetCourseList()*gorm.DB  {
	return boot.GetDB().Table(vars.TableCoursemain).Order("course_id desc").Limit(3)
}

const courseList = "select "+ Field +" from " + vars.TableCoursemain + " order by course_id desc limit ?"

func GetCourseListBySql(args ...interface{}) *gorm.DB {
	fmt.Println(courseList)
	return boot.GetDB().Raw(courseList, args...)
}
// 获取课程详情
func GetCourseDetail(courseId int) *gorm.DB {
	return boot.GetDB().Table(vars.TableCoursemain).Where("course_id = ?", courseId)
}
// 根据课程id获取课程统计
func GetCourseCountsById(courseId int) *gorm.DB {
	return boot.GetDB().Table(vars.TableCourseCounts).Where("course_id = ? and ", courseId)
}

package service

import (
	"context"
	"github.com/shijting/go-micro/src/courses"
	"github.com/shijting/go-micro/src/mapper"
)

type CourseServices struct {
}

func NewCourse() *CourseServices {
	return &CourseServices{}
}



func (h *CourseServices) ListForTop(ctx context.Context, request *courses.ListRequest, response *courses.ListResponse) error {
	result := make([]*courses.CourseModel, 0)
	err:= mapper.GetCourseListBySql(1).Find(&result).Error
	if err!=nil{
		return err
	}
	response.Result = result
	return nil
}
// FetchType == 1: 只取课程详情， 2 只取 课程统计， 3： 两者都取
func (this *CourseServices) GetDetail(ctx context.Context, req *courses.DetailRequest, resp *courses.DetailResponse) error  {
	if req.FetchType == 0 || req.FetchType == 1 || req.FetchType==3 {
		if err := mapper.GetCourseDetail(int(req.CourseId)).First(resp.Course).Error; err !=nil {
			return err
		}
		if req.FetchType== 2 || req.FetchType==3 {
			if err := mapper.GetCourseCountsById(int(req.CourseId)).Find(resp.Counts); err !=nil{
				return err.Error
			}
		}
	}

	return nil
}

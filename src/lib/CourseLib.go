package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-micro/framework/gin_"
	"github.com/shijting/go-micro/src/courses"
	"github.com/shijting/go-micro/src/service"
)

func init()  {
	courseService:=service.NewCourse()
	gin_.NewBuilder().WithService(courseService).
		WithMiddleware(CheckMiddleware()).
		WithEndpoint(Courselist_Endpoint(courseService)).
		WithRequest(Courselist_Request()).
		WithResponse(Course_Response()).Build("/test","GET")
	gin_.NewBuilder().WithService(courseService).
		WithMiddleware(CheckMiddleware()).
		WithEndpoint(CourseDetail_Endpoint(courseService)).
		WithRequest(CourseDetail_Request()).
		WithResponse(Course_Response()).Build("/detail/:course_id","GET")
}

func Test_Middleware()  gin_.Middleware {
	return func(next gin_.Endpoint) gin_.Endpoint {
		return func(context *gin.Context, request interface{}) (response interface{}, err error) {
			return next(context,request)
		}
	}
}

//获取列表相关
func Courselist_Endpoint(c *service.CourseServices)   gin_.Endpoint {
	return func(context *gin.Context, request interface{}) (response interface{}, err error) {
		rsp:=&courses.ListResponse{}
	    err=c.ListForTop(context,request.(*courses.ListRequest),rsp)
	    return rsp,err
	}
}
//这个函数的作用是怎么处理请求
func Courselist_Request() gin_.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&courses.ListRequest{}
		err:=context.BindQuery(bReq) //使用的是query 参数
		if err!=nil{
			return nil,err
		}
		return bReq,nil
	}
}

//获取详情响应
func CourseDetail_Endpoint(c *service.CourseServices)   gin_.Endpoint {
	return func(context *gin.Context, request interface{}) (response interface{}, err error) {
		rsp:=&courses.DetailResponse{Course: new(courses.CourseModel), Counts: make([]*courses.CourseCounts, 0)}
		err=c.GetDetail(context,request.(*courses.DetailRequest),rsp)
		return rsp,err
	}
}
//这个函数的作用是怎么处理请求
func CourseDetail_Request() gin_.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&courses.DetailRequest{}
		err:=context.BindUri(bReq) //使用的是uri 参数 GET /test/1
		if err!=nil{
			return nil,err
		}
		err = context.BindHeader(bReq)
		if err!=nil{
			return nil,err
		}
		return bReq,nil
	}
}
//这个函数作用是：怎么处理响应结果
func Course_Response()  gin_.DecodeResponseFunc  {
	return func(context *gin.Context, res interface{}) error {
		context.JSON(200,res)
		return nil
	}
}
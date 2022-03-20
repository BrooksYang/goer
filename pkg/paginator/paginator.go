package paginator

import (
	"math"

	"goapp/global"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Pagination 返回分页数据
type Pagination struct {
	Paging
	Data interface{} `json:"data"`
}

// Paging 分页数据
type Paging struct {
	Total       int64 `json:"total"`
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
	PerPage     int   `json:"per_page"`
}

// Paginator 分页操作类
type Paginator struct {
	PerPage  int
	Page     int
	Offset   int
	Total    int64
	LastPage int
	Sort     string
	Order    string

	query *gorm.DB     // db query 句柄
	ctx   *gin.Context // gin context，方便调用
}

// Paginate 分页
// c —— gin.context 用来获取分页的 URL 参数
// db —— GORM 查询句柄，用以查询数据集和获取数据总数
// data —— 模型数组，传址获取数据
// PerPage —— 每页条数
func Paginate(c *gin.Context, db *gorm.DB, data interface{}, perPage int) Pagination {

	// 初始化 Paginator 实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage)

	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	if err != nil {
		return Pagination{}
	}

	paging := Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		LastPage:    p.LastPage,
		Total:       p.Total,
	}

	return Pagination{
		Paging: paging,
		Data:   data,
	}
}

// 初始化分页必须用到的属性，基于这些属性查询数据库
func (p *Paginator) initProperties(perPage int) {

	p.PerPage = p.getPerPage(perPage)

	// 排序参数
	p.Order = p.ctx.DefaultQuery("order", "desc")
	p.Sort = p.ctx.DefaultQuery("sort", "id")

	p.Total = p.getTotalCount()
	p.LastPage = p.getLastPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p Paginator) getPerPage(perPage int) int {
	// 优先使用请求 per_page 参数
	queryPerPage := p.ctx.Query("per_page")
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	// 没有传参，使用默认
	if perPage <= 0 {
		perPage = global.Config.Paging.PerPage
	}

	return perPage
}

func (p Paginator) getCurrentPage() int {
	// 优先取用户请求的 page
	page := cast.ToInt(p.ctx.Query("page"))
	if page <= 0 {
		page = 1
	}

	// LastPage 等于 0 ，意味着数据不够分页
	if p.LastPage == 0 {
		return 0
	}

	// 请求页数大于总页数，返回总页数
	if page > p.LastPage {
		return p.LastPage
	}

	return page
}

func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p Paginator) getLastPage() int {
	if p.Total == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.Total) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}

	return int(nums)
}

package paginator

import (
	"math"

	"gorm.io/gorm"
)

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 10
)

type Paginator interface {
	Page() int
	PageSize() int
	Total() int
	TotalPages() int
	HasNextPage() bool
	HasPreviousPage() bool
	FirstPage() int
	LastPage() int
	NextPage() int
	PreviousPage() int
	GetItems() interface{}
}

type paginator[T any] struct {
	page       int
	pageSize   int
	total      int
	totalPages int
	items      *[]T
}

func (p *paginator[T]) Page() int             { return p.page }
func (p *paginator[T]) PageSize() int         { return p.pageSize }
func (p *paginator[T]) Total() int            { return p.total }
func (p *paginator[T]) TotalPages() int       { return p.totalPages }
func (p *paginator[T]) HasNextPage() bool     { return p.page < p.totalPages }
func (p *paginator[T]) HasPreviousPage() bool { return p.page > 1 }
func (p *paginator[T]) FirstPage() int        { return 1 }
func (p *paginator[T]) LastPage() int         { return p.totalPages }
func (p *paginator[T]) NextPage() int         { return min(p.page+1, p.totalPages) }
func (p *paginator[T]) PreviousPage() int     { return max(p.page-1, 1) }
func (p *paginator[T]) GetItems() interface{} { return p.items }

func NewPaginator[T any](db *gorm.DB, page, pageSize int, dest *[]T) (Paginator, error) {
	if page < 1 {
		page = DEFAULT_PAGE
	}
	if pageSize < 1 {
		pageSize = DEFAULT_PAGE_SIZE
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(dest).Error; err != nil {
		return nil, err
	}

	return &paginator[T]{
		page:       page,
		pageSize:   pageSize,
		total:      int(total),
		totalPages: totalPages,
		items:      dest,
	}, nil
}

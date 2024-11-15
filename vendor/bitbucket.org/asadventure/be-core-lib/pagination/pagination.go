package pagination

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Pagination Limit
const PaginationLimit = 48

const PageTimeLayout = time.RFC3339

// PageNumber Model
type PageMode uint

const (
	PageModeNumber PageMode = iota
	PageModeTime   PageMode = iota
)

// Pagination
type Pagination struct {
	// Start
	Start uint64
	// Limit
	Limit uint64
	// Page Number
	PageNumber uint64
	// Page Time
	PageTime *time.Time
	// PageMode
	PageMode PageMode
	// No Next Page
	NoNextPage bool
	// Total Pages
	TotalPages uint64
	// Records
	Records uint64
	// Path
	Path string
	// Sort
	Sort string
	// Other Params
	OtherParams string
	// Search
	Search Search
}

// TotalFunc total function
type TotalFunc func(...interface{}) (uint64, error)

// Prepare prepare pagination
func (p *Pagination) Prepare() {
	p.ValidatePage()
	p.CalculateStart()
	p.CalculateTotalPages()
}

// PrepareWithTotal prepare pagination with total
func (p *Pagination) PrepareWithTotal(totalFunc TotalFunc, args ...interface{}) (err error) {
	p.Records, err = totalFunc(args...)
	p.Prepare()
	return
}

// AddSearch add search
func (p *Pagination) AddSearch(search Search) {
	p.Search = search
}

// AddParam add a parameter
func (p *Pagination) AddParam(key string, value string) {
	if p.OtherParams != "" {
		p.OtherParams += "&"
	}
	p.OtherParams += fmt.Sprintf("%s=%s", key, value)
}

// ValidatePage validate page
func (p *Pagination) ValidatePage() {
	if p.PageMode == PageModeNumber && p.PageNumber == 0 {
		p.PageNumber = 1
	}
}

// CalculateStart calculate start
func (p *Pagination) CalculateStart() {
	if p.PageMode == PageModeNumber {
		p.Start = (p.PageNumber - 1) * p.Limit
	}
}

// CalculateTotalPages calculate total pages
func (p *Pagination) CalculateTotalPages() {
	if p.Limit == 0 {
		p.TotalPages = 1
	} else {
		p.TotalPages = uint64(math.Ceil(float64(p.Records) / float64(p.Limit)))
	}
}

// IsEmpty check if is empty
func (p *Pagination) IsEmpty() bool {
	return *p == Pagination{}
}

// First first page
func (p *Pagination) First() *Pagination {
	if p.Records == 0 || p.PageMode == PageModeTime {
		return &Pagination{}
	}

	return &Pagination{
		PageMode:   p.PageMode,
		PageNumber: 1,
		Limit:      p.Limit,
		Records:    p.Records,
		TotalPages: p.TotalPages,
		Path:       p.Path,
		Search:     p.Search,
		Sort:       p.Sort,
	}
}

// Prev previous page
func (p *Pagination) Prev() *Pagination {
	if p.PageMode == PageModeTime ||
		p.PageNumber < 2 || p.Records == 0 {
		return &Pagination{}
	}

	return &Pagination{
		PageMode:   p.PageMode,
		PageNumber: p.PageNumber - 1,
		Limit:      p.Limit,
		Records:    p.Records,
		TotalPages: p.TotalPages,
		Path:       p.Path,
		Search:     p.Search,
		Sort:       p.Sort,
	}
}

// Next next page
func (p *Pagination) Next() *Pagination {
	if (p.PageMode == PageModeNumber && p.PageNumber >= p.TotalPages) ||
		(p.PageMode == PageModeTime && p.NoNextPage) ||
		p.Records == 0 {
		return &Pagination{}
	}

	next := &Pagination{
		PageMode:   p.PageMode,
		PageNumber: p.PageNumber + 1,
		Limit:      p.Limit,
		Records:    p.Records,
		PageTime:   p.PageTime,
		TotalPages: p.TotalPages,
		Path:       p.Path,
		Search:     p.Search,
		Sort:       p.Sort,
	}

	switch p.PageMode {
	case PageModeNumber:
		p.PageNumber = p.PageNumber + 1
	case PageModeTime:
		if p.PageTime != nil {
			t := p.PageTime.Add(-time.Second)
			next.PageTime = &t
		}
	}

	return next
}

// Last last page
func (p *Pagination) Last() *Pagination {
	if p.PageMode == PageModeTime || p.Records == 0 {
		return &Pagination{}
	}

	return &Pagination{
		PageMode:   p.PageMode,
		PageNumber: p.TotalPages,
		Limit:      p.Limit,
		PageTime:   p.PageTime,
		Records:    p.Records,
		TotalPages: p.TotalPages,
		Path:       p.Path,
		Search:     p.Search,
		Sort:       p.Sort,
	}
}

// ParamsToString params to string
func (p *Pagination) ParamsToString(concat string) string {
	if p.OtherParams == "" {
		return ""
	}
	return fmt.Sprintf("%s%s", concat, p.OtherParams)
}

// ToString to string
func (p *Pagination) ToString() *string {
	if p.IsEmpty() {
		return nil
	}

	var page string
	switch p.PageMode {
	case PageModeNumber:
		page = strconv.Itoa(int(p.PageNumber))
	case PageModeTime:
		if p.PageTime != nil {
			page = p.PageTime.Format(PageTimeLayout)
		}
	}

	result := fmt.Sprintf(
		"%s?page=%s&limit=%d"+p.ParamsToString("&")+p.Search.ToString("&"),
		p.Path,
		page,
		p.Limit,
	)

	if p.Sort != "" {
		result = fmt.Sprintf("%s&sort=%s", result, p.Sort)
	}

	return &result
}

// SetPaginationPath set pagination path
func (p *Pagination) SetPaginationPath(path string) {
	p.Path = ""
	params := strings.Split(path, "/")
	for index, value := range params {
		if index > 1 {
			p.Path = fmt.Sprintf("%s/%s", p.Path, value)
		}
	}
}

// GetStart get start
func (p *Pagination) GetStart() uint64 {
	return p.Start
}

// GetLimit get limit
func (p *Pagination) GetLimit() uint64 {
	return p.Limit
}

// SetRecords set records
func (p *Pagination) SetRecords(records uint64) {
	p.Records = records
}

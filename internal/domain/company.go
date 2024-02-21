package domain

import (
	"time"

	"github.com/guregu/null"
	"github.com/uptrace/bun"
)

var (
	_ = time.Second
	_ = null.Bool{}
)

type Company struct {
	bun.BaseModel `bun:"table:company,alias:c" json:"-" swaggerignore:"true"`
	ID            string      `bun:"id,pk" gorm:"primary_key;column:id;type:VARCHAR;size:45;" json:"id"`
	Code          null.String `gorm:"column:code;type:VARCHAR;size:20" json:"code" swaggertype:"string"`
	Name          string      `gorm:"column:name;type:VARCHAR;" json:"name"`
	Status        int32       `gorm:"column:status;type:INT4;" json:"status"`
	CreateBy      string      `gorm:"column:create_by;type:VARCHAR;size:45;" json:"createBy"`
	CreateDate    time.Time   `gorm:"column:create_date;type:TIMESTAMPTZ;" json:"createDate" swaggertype:"string" format:"date-time"`
	UpdateBy      string      `gorm:"column:update_by;type:VARCHAR;size:45;" json:"updateBy"`
	UpdateDate    time.Time   `gorm:"column:update_date;type:TIMESTAMPTZ;" json:"updateDate" swaggertype:"string" format:"date-time"`
}

func (c *Company) TableName() string {
	return "company"
}

package models

// AdminAreaCreateDTO 添加行政区划表
type AdminAreaCreateDTO struct {
	// 上级行政区划ID
	ParentID uint64 `form:"parent_id" json:"parent_id" xml:"parent_id" yaml:"parent_id" bson:"parent_id" binding:"omitempty"`
	// 行政区划名称
	Name string `form:"name" json:"name" xml:"name" yaml:"name" bson:"name" binding:"required,min=2,max=255"`
	// 备注
	Memo string `form:"memo" json:"memo" xml:"memo" yaml:"memo" bson:"memo" binding:"omitempty"`
}

package models

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"eidng8.cc/microservices/admin-area/ent"
)

// aaCommonVO holds common field to AdminArea structures
type aaCommonVO struct {
	// ID of the entity.
	ID uint64 `json:"id,omitempty"`
	// 行政区划名称
	Name string `json:"name,omitempty"`
	// 备注
	Memo *sql.NullString `json:"memo,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (m *aaCommonVO) basicFromAdminArea(adminArea *ent.AdminArea) {
	m.ID = adminArea.ID
	m.Name = adminArea.Name
	m.Memo = adminArea.Memo
	m.CreatedAt = &adminArea.CreatedAt
	m.UpdatedAt = adminArea.UpdatedAt
}

// AdminAreaVO 行政区划表
type AdminAreaVO struct {
	aaCommonVO
	// ParentID holds the value of the "parent_id" field.
	Parent *AdminAreaVO `json:"parent,omitempty"`
	// Children 下级行政区划列表.
	Children []AdminAreaVO `json:"children,omitempty"`
}

func (m *AdminAreaVO) FromAdminArea(adminArea *ent.AdminArea) {
	m.basicFromAdminArea(adminArea)
	if parent, err := adminArea.Edges.ParentOrErr(); err == nil {
		m.Parent.FromAdminArea(parent)
	} else {
		m.Parent = &AdminAreaVO{aaCommonVO: aaCommonVO{ID: *adminArea.ParentID}}
	}
	if children, err := adminArea.Edges.ChildrenOrErr(); err == nil {
		m.Children = make([]AdminAreaVO, len(children))
		for _, child := range children {
			var childVO AdminAreaVO
			childVO.FromAdminArea(child)
			m.Children = append(m.Children, childVO)
		}
	}
}

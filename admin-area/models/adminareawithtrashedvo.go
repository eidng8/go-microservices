package models

import (
	"time"

	"eidng8.cc/microservices/admin-area/ent"
)

// AdminAreaWithTrashedVO 行政区划表（包含已删除记录）
type AdminAreaWithTrashedVO struct {
	aaCommonVO
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	Parent *AdminAreaWithTrashedVO `json:"parent,omitempty"`
	// Children 下级行政区划列表.
	Children []AdminAreaWithTrashedVO `json:"children,omitempty"`
}

func (m *AdminAreaWithTrashedVO) FromAdminArea(adminArea *ent.AdminArea) {
	m.basicFromAdminArea(adminArea)
	m.DeletedAt = adminArea.DeletedAt
	if parent, err := adminArea.Edges.ParentOrErr(); err == nil {
		m.Parent.FromAdminArea(parent)
	} else {
		m.Parent = &AdminAreaWithTrashedVO{
			aaCommonVO: aaCommonVO{ID: *adminArea.ParentID},
		}
	}
	if children, err := adminArea.Edges.ChildrenOrErr(); err == nil {
		m.Children = make([]AdminAreaWithTrashedVO, len(children))
		for _, child := range children {
			var childVO AdminAreaWithTrashedVO
			childVO.FromAdminArea(child)
			m.Children = append(m.Children, childVO)
		}
	}
}

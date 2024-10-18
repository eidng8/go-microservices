package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"

	"eidng8.cc/microservices/admin-area/models"
	"eidng8.cc/microservices/common"
)

func list(c *gin.Context, env *Env) {
	page := common.GetPaginationParams(c)
	areas, err := env.db.AdminArea.
		Query().
		Offset(page.Page * page.PerPage).
		Limit(page.PerPage).
		All(context.Background())
	if err != nil {
		log.Printf("Could not query admin areas: %v", err)
		common.ErrorJSON(c, err)
		return
	}
	vo := make([]models.AdminAreaVO, len(areas))
	for i, area := range areas {
		vo[i].FromAdminArea(area)
	}
	common.RespondJSON(c, vo)
}

func create(c *gin.Context, env *Env) {
	var data models.AdminAreaCreateDTO
	if err := c.ShouldBind(&data); err != nil {
		common.Error422JSON(c, err)
		return
	}
	area := env.db.AdminArea.Create()
	area.SetName(data.Name)
	if data.ParentID > 0 {
		area.SetParentID(data.ParentID)
	}
	if data.Memo != "" {
		area.SetMemo(&sql.NullString{String: data.Memo, Valid: true})
	}
	saved, err := area.Save(context.Background())
	if err != nil {
		common.ErrorJSON(c, err)
		return
	}
	var vo models.AdminAreaVO
	vo.FromAdminArea(saved)
	common.RespondJSON(c, vo)
}

func update(c *gin.Context, env *Env) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

func detail(c *gin.Context, env *Env) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

func remove(c *gin.Context, env *Env) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

func restore(c *gin.Context, env *Env) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

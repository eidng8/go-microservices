package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"eidng8.cc/microservices/common"
)

type AdminArea struct {
	Id        int64
	Name      string
	Memo      sql.NullString
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
	DeletedAt sql.NullString
	lft       int
	rgt       int
	ParentId  sql.NullInt64
}

func rowsToArray(rows *sql.Rows) ([]AdminArea, error) {
	var adminAreas []AdminArea
	for rows.Next() {
		adminArea := AdminArea{}
		err := rows.Scan(
			&adminArea.Id, &adminArea.Name, &adminArea.Memo,
			&adminArea.CreatedAt, &adminArea.UpdatedAt, &adminArea.DeletedAt,
			&adminArea.lft, &adminArea.rgt, &adminArea.ParentId,
		)
		if err != nil {
			return nil, err
		}
		adminAreas = append(adminAreas, adminArea)
	}
	return adminAreas, nil
}

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
	common.RespondJSON(c, areas)
}

func create(c *gin.Context) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

func update(c *gin.Context) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

func detail(c *gin.Context) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

func remove(c *gin.Context) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

func restore(c *gin.Context) {
	common.ErrorWithCodeJSON(c, http.StatusNotImplemented, nil)
}

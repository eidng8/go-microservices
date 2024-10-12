package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"eidng8.cc/microservices/common"
	"eidng8.cc/microservices/rdbms/mysql"
)

type AdminArea struct {
	Id        int64
	Name      string
	Memo      sql.NullString
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
	DeletedAt sql.NullString
	_lft      int
	_rgt      int
	ParentId  sql.NullInt64
}

func rowsToArray(rows *sql.Rows) ([]AdminArea, error) {
	var adminAreas []AdminArea
	for rows.Next() {
		adminArea := AdminArea{}
		err := rows.Scan(
			&adminArea.Id, &adminArea.Name, &adminArea.Memo,
			&adminArea.CreatedAt, &adminArea.UpdatedAt, &adminArea.DeletedAt,
			&adminArea._lft, &adminArea._rgt, &adminArea.ParentId,
		)
		if err != nil {
			return nil, err
		}
		adminAreas = append(adminAreas, adminArea)
	}
	return adminAreas, nil
}

func list(c *gin.Context) {
	db, err := mysql.Connect(host, user, pass, dbname)
	if err != nil {
		common.ErrorJSON(c, err)
		return
	}

	rows, err := db.Query("SELECT * FROM `admin_areas` LIMIT ? OFFSET ?", 10, 0)
	if err != nil {
		common.ErrorJSON(c, err)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			common.ErrorJSON(c, err)
		}
	}(rows)

	res, err := rowsToArray(rows)
	if err != nil {
		common.ErrorJSON(c, err)
		return
	}

	common.RespondJSON(c, res)
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

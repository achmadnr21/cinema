package middleware

import (
	"database/sql"
	"net/http"

	"github.com/achmadnr21/cinema/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DbMiddleware struct {
	db *sql.DB
}

var Dbmiddleware *DbMiddleware

func InitDbMiddleware(db *sql.DB) {
	if Dbmiddleware == nil {
		Dbmiddleware = &DbMiddleware{db: db}
	}
}
func RequirePermissionCinema(module, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil cinema_id dari URL param
		cinemaIDStr := c.Param("cinema_id")
		cinemaID, err := uuid.Parse(cinemaIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseError("Invalid cinema ID format"))
			return
		}

		// Ambil user_id dari context (pastikan di-set oleh auth middleware)
		userID := c.GetString("user_id")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseError("User not authenticated"))
			return
		}

		// Cek ke database
		query := `
			SELECT 1
			FROM employee_details ed
			JOIN role_engine.roles rer ON ed.role_id = rer.id
			JOIN role_engine.role_permissions rp ON rer.id = rp.role_id 
			JOIN role_engine.permissions p ON rp.permission_id = p.id
			WHERE ed.cinema_id = $1 AND ed.user_id = $2 AND p.module = $3 AND p.action = $4
			LIMIT 1
		`

		var dummy int
		err = Dbmiddleware.db.QueryRow(query, cinemaID, userID, module, action).Scan(&dummy)
		if err != nil {
			c.AbortWithStatusJSON(403, utils.ResponseError("Permission denied"))
			return
		}

		// Jika permission ada, lanjut
		c.Next()
	}
}

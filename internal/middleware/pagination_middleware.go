package middleware

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/utils/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type driverConstructor func(pq dto.PaginationQuery) pagination.Pagination

var paginationStrategies = map[dto.PaginationType]driverConstructor{
	dto.CursorBased: func(pq dto.PaginationQuery) pagination.Pagination {
		return pagination.NewCursorDriver(pq.LastSeenID, pq.PerPage, pq.SortField, pq.Asc)
	},
	dto.PageBased: func(pq dto.PaginationQuery) pagination.Pagination {
		return pagination.NewPageDriver(pq.Page, pq.PerPage, pq.SortField, pq.Asc)
	},
}

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pq dto.PaginationQuery
		pq.SetDefaults()
		if err := c.ShouldBindQuery(&pq); err != nil {
			response.New(c).
				Errors(err).
				MessageID("pagination.middleware.failed").
				Status(http.StatusBadRequest).
				Dispatch()
			return
		}
		constructor, ok := paginationStrategies[pq.Type]
		if !ok {
			constructor = paginationStrategies[dto.PageBased] // fallback
		}
		driver := constructor(pq)
		ctx := context.WithValue(c.Request.Context(), constant.PaginatorCtxKey, driver)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

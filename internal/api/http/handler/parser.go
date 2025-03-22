package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func parseQueryParams(r *http.Request) (limit int, offset int, orderBy *string, query *string, err error) {
	queryParams := r.URL.Query()

	limitStr := queryParams.Get("limit")
	if limitStr == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return 0, 0, nil, nil, fmt.Errorf("invalid limit value: %v", err)
		}
	}

	offsetStr := queryParams.Get("offset")
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			return 0, 0, nil, nil, fmt.Errorf("invalid offset value: %v", err)
		}
	}

	orderByStr := queryParams.Get("order_by")
	if orderByStr != "" {
		orderBy = &orderByStr
	}

	queryStr := queryParams.Get("query")
	if queryStr != "" {
		query = &queryStr
	}

	return limit, offset, orderBy, query, nil
}

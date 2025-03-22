package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func parseQueryParams(r *http.Request) (limit int, offset int, orderBy *string, query *string, err error) {
	// Extract the query parameters
	queryParams := r.URL.Query()

	// Extract limit and offset as integers, with default values
	limitStr := queryParams.Get("limit")
	if limitStr == "" {
		limit = 10 // Default limit
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return 0, 0, nil, nil, fmt.Errorf("invalid limit value: %v", err)
		}
	}

	offsetStr := queryParams.Get("offset")
	if offsetStr == "" {
		offset = 0 // Default offset
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			return 0, 0, nil, nil, fmt.Errorf("invalid offset value: %v", err)
		}
	}

	// Extract orderBy and query as pointers, they can be nil if not provided
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

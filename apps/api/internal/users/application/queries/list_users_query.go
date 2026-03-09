package queries

type ListUsersQuery struct {
	Page      int
	Limit     int
	SortBy    string
	SortOrder string
	Filters   map[string]string
}

func NewListUsersQuery(
	page int,
	limit int,
	sortBy string,
	sortOrder string,
	filters map[string]string,
) ListUsersQuery {

	return ListUsersQuery{
		Page:      page,
		Limit:     limit,
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Filters:   filters,
	}
}

package table

type PaginationLink struct {
	PageNumber int
	Offset     int
	Active     bool
}

type PaginationProps struct {
	Offset int
	Limit  int
	Total  int

	Target  string
	BaseURL string
}

func CreatePaginationLinks(props PaginationProps) []PaginationLink {
	var links []PaginationLink
	pageNumber := props.Offset / props.Limit

	start := 0
	// Calculate last page number, ensuring we don't show empty pages
	pageCount := (props.Total + props.Limit - 1) / props.Limit
	end := pageCount - 1 // Zero-based indexing

	if end > 5 {
		start = pageNumber - 2
		end = pageNumber + 2

		for start < 0 {
			end++
			start++
		}
		for end > pageCount-1 {
			end--
			start--
		}

	}

	for i := start; i <= end; i++ {
		links = append(links, PaginationLink{
			PageNumber: i,
			Offset:     i * props.Limit,
			Active:     i == pageNumber,
		})
	}
	return links
}

package table

import (
	"testing"
)

func TestCreatePaginationLinks(t *testing.T) {
	tests := []struct {
		name  string
		props PaginationProps
		want  []PaginationLink
	}{
		{
			name: "single page",
			props: PaginationProps{
				Offset: 0,
				Limit:  10,
				Total:  5,
			},
			want: []PaginationLink{
				{PageNumber: 0, Offset: 0, Active: true},
			},
		},
		{
			name: "3 pages",
			props: PaginationProps{
				Offset: 0,
				Limit:  10,
				Total:  29,
			},
			want: []PaginationLink{
				{PageNumber: 0, Offset: 0, Active: true},
				{PageNumber: 1, Offset: 10},
				{PageNumber: 2, Offset: 20},
			},
		},
		{
			name: "3 pages starting from page 2",
			props: PaginationProps{
				Offset: 20,
				Limit:  10,
				Total:  29,
			},
			want: []PaginationLink{
				{PageNumber: 0, Offset: 0},
				{PageNumber: 1, Offset: 10},
				{PageNumber: 2, Offset: 20, Active: true},
			},
		},
		{
			name: "7 pages starting from page 1",
			props: PaginationProps{
				Offset: 10,
				Limit:  10,
				Total:  69,
			},
			want: []PaginationLink{
				{PageNumber: 0, Offset: 0},
				{PageNumber: 1, Offset: 10, Active: true},
				{PageNumber: 2, Offset: 20},
				{PageNumber: 3, Offset: 30},
				{PageNumber: 4, Offset: 40},
			},
		},
		{
			name: "7 pages starting from page 3",
			props: PaginationProps{
				Offset: 30,
				Limit:  10,
				Total:  69,
			},
			want: []PaginationLink{
				{PageNumber: 1, Offset: 10},
				{PageNumber: 2, Offset: 20},
				{PageNumber: 3, Offset: 30, Active: true},
				{PageNumber: 4, Offset: 40},
				{PageNumber: 5, Offset: 50},
			},
		},
		{
			name: "7 pages starting from page 6",
			props: PaginationProps{
				Offset: 60,
				Limit:  10,
				Total:  69,
			},
			want: []PaginationLink{
				{PageNumber: 2, Offset: 20},
				{PageNumber: 3, Offset: 30},
				{PageNumber: 4, Offset: 40},
				{PageNumber: 5, Offset: 50},
				{PageNumber: 6, Offset: 60, Active: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreatePaginationLinks(tt.props)
			if len(got) != len(tt.want) {
				t.Fatalf("CreatePaginationLinks() = %v, want %v", got, tt.want)
			}
			for i, want := range tt.want {
				got := got[i]
				if got.PageNumber != want.PageNumber {
					t.Errorf("CreatePaginationLinks() = %v, want PageNumber=%d", got, want.PageNumber)
				}
				if got.Offset != want.Offset {
					t.Errorf("CreatePaginationLinks() = %v, want Offset=%d", got, want.Offset)
				}
				if got.Active != want.Active {
					t.Errorf("CreatePaginationLinks() = %v, want Active=%t", got, want.Active)
				}
			}
		})
	}

}

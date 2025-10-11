package entity

type FilterType string

const (
	StringRange FilterType = "STRING_RANGE"
	NumberRange FilterType = "NUMBER_RANGE"
	DateRange   FilterType = "DATE_RANGE"
)

func (filterType FilterType) String() string {
	return string(filterType)
}

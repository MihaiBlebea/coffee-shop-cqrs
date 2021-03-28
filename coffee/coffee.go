package coffee

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Coffee struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Size        Size           `json:"size"`
	Price       uint           `json:"price"`
	Ingredients map[string]int `json:"ingredients"`
}

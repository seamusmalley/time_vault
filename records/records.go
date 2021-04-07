package records

type record struct {
	id string `json:"ID"`
	hash string `json:"hash"`
	description string `json:"description"`
}

type recordList []record


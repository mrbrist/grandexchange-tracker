package utils

type ItemLookupData struct {
	Examine  string `json:"examine"`
	ID       int    `json:"id"`
	Members  bool   `json:"members"`
	Lowalch  int    `json:"lowalch,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Value    int    `json:"value,omitempty"`
	Highalch int    `json:"highalch,omitempty"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
}

func LookupItem(store *Items, id int) ItemLookupData {
	for _, i := range *store {
		if i.ID == id {
			return ItemLookupData(i)
		}
	}
	return ItemLookupData{}
}

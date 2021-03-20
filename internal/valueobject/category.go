package valueobject

// カテゴリ
const (
	CategoryFood                    string = "食費"
	CategoryDailyNecessities        string = "日用品"
	CategoryTraveling               string = "交通費"
	CategoryHobby                   string = "趣味"
	CategoryFurnitureAndAppliances  string = "家具・家電"
	CategoryEntertainment           string = "交際費"
	CategoryLiberalArtsAndEducation string = "教養・教育"
	CategoryHealthAndMedicine       string = "健康・医療"
	CategoryFinance                 string = "金融"
	CategoryHome                    string = "住宅"
	CategoryWaterAndUtility         string = "水道・光熱費"
	CategoryWifi                    string = "通信費"
	CategoryTax                     string = "税金"
	CategoryCar                     string = "自動車"
	CategoryOther                   string = "その他"
)

// Category type
type Category struct {
	value string
}

// NewCategory method
func NewCategory(value string) (*Category, error) {
	if value == "" {
		return nil, newRequiredError("カテゴリ")
	}
	if containMutibyte(value) {
		return nil, newContainMutibyteError("カテゴリ")
	}
	if !containCategories(value) {
		return nil, newContainItemsError("カテゴリ")
	}

	return &Category{value: value}, nil
}

func containCategories(value string) bool {
	if value == CategoryFood {
		return true
	}
	if value == CategoryDailyNecessities {
		return true
	}
	if value == CategoryTraveling {
		return true
	}
	if value == CategoryHobby {
		return true
	}
	if value == CategoryFurnitureAndAppliances {
		return true
	}
	if value == CategoryEntertainment {
		return true
	}
	if value == CategoryLiberalArtsAndEducation {
		return true
	}
	if value == CategoryHealthAndMedicine {
		return true
	}
	if value == CategoryFinance {
		return true
	}
	if value == CategoryHome {
		return true
	}
	if value == CategoryWaterAndUtility {
		return true
	}
	if value == CategoryWifi {
		return true
	}
	if value == CategoryTax {
		return true
	}
	if value == CategoryCar {
		return true
	}
	if value == CategoryOther {
		return true
	}
	return false
}

// Get method
func (vo *Category) Get() string {
	return vo.value
}

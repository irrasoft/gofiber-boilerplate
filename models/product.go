package models

// Product struct
type Product struct {
	Model
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"not null" json:"description"`
	Amount      float64 `gorm:"not null" json:"amount"`
}

// GetAllProducts ...
func GetAllProducts(query map[string]string, order []string, offset int, limit int) (products []Product, total int64, err error) {
	qs := orm.Model(&products)

	// query k=v
	for k, v := range query {
		qs = qs.Where(k, v)
	}

	// order
	for _, v := range order {
		qs = qs.Order(v)
	}

	// get count
	qs.Count(&total)

	qs = qs.Limit(limit).Offset(offset)

	if qs.Find(&products); products != nil {
		return products, total, nil
	}

	return nil, 0, err
}

// GetProductByID ...
func GetProductByID(id int) (product Product, err error) {
	err = orm.First(&product, id).Error
	if err == nil {
		return product, nil
	}
	return product, err
}

// CreateProduct ...
func CreateProduct(product *Product) (id uint, err error) {
	if err = orm.Create(&product).Error; err != nil {
		return 0, err
	}
	return product.ID, nil
}

// UpdateProduct ...
func UpdateProduct(product *Product) (id uint, err error) {
	if err = orm.Updates(&product).Error; err != nil {
		return 0, err
	}
	return product.ID, nil
}

// DeleteProduct ...
// soft mean soft delete
func DeleteProduct(id int, soft bool) (err error) {
	qs := orm
	if soft != true {
		qs = qs.Unscoped() // Delete permanently
	}

	return qs.Delete(&Product{}, id).Error
}

// MultiDeleteProducts ...
func MultiDeleteProducts(list []int, soft bool) (err error) {
	qs := orm
	if soft != true {
		qs = qs.Unscoped() // Delete permanently
	}

	return qs.Delete(&Product{}, list).Error
}

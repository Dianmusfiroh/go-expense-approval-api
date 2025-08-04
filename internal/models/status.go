package models


type Status struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(50)"`
	Expenses  []Expense  `gorm:"foreignKey:StatusID"`
	Approvals []Approval `gorm:"foreignKey:StatusID"`
}

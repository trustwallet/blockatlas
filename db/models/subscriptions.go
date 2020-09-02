package models

type (
	NotificationSubscription struct {
		//DeletedAt *time.Time `sql:"index"`
		Address   Address `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint    `gorm:"unique_index"`
	}

	AssetSubscription struct {
		//DeletedAt *time.Time `sql:"index"`
		Address   Address `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint    `gorm:"unique_index"`
	}
)

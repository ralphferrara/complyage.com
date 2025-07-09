package models

type API struct {
	IDAPI           uint   `gorm:"column:id_api;primaryKey;autoIncrement"`
	FidSite         string `gorm:"column:fid_site;size:45;index"`
	APIPrivate      string `gorm:"column:api_private;size:128;index"`
	APIPublic       string `gorm:"column:api_public;size:128;index"`
	APIStatus       string `gorm:"column:api_status;size:4;index"`
	APIRequirements string `gorm:"column:api_requirements;type:text"`
	APITerritories  string `gorm:"column:api_territories;type:text"`
}

func (API) TableName() string {
	return "apis"
}

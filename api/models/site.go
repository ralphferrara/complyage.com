package models

type Site struct {
	IDSite     uint   `gorm:"column:id_site;primaryKey;autoIncrement"`
	SiteName   string `gorm:"column:site_name;size:128"`
	SiteURL    string `gorm:"column:site_url;size:255;index"`
	SiteStatus string `gorm:"column:site_status;size:4;index"`
}

// TableName overrides the table name used by GORM
func (Site) TableName() string {
	return "sites"
}

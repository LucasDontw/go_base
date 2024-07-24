package model

import "time"

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key"`  // 內容ID
	Title          string        `gorm:"column:title"`           // 內容標題
	Description    string        `gorm:"column:description"`     // 內容描述
	Author         string        `gorm:"column:author"`          // 作者
	VideoURL       string        `gorm:"column:video_url"`       // 視頻撥放URL
	Thumbnail      string        `gorm:"column:thumbnail"`       // 封面圖URL
	Catgory        string        `gorm:"column:catgory"`         // 內容分類
	Duration       time.Duration `gorm:"column:duration"`        // 內容時長
	Resolution     string        `gorm:"column:resolution"`      // 分辨率 ex:720p
	FileSize       int64         `gorm:"column:file_size"`       // 文件大小
	Format         string        `gorm:"column:format"`          // 文件格式 ex:mp4
	Quality        int           `gorm:"column:quality"`         // 影片品質 1-高清 2-標清
	ApprovalStatus int           `gorm:"column:approval_status"` // 審核狀態 1-審核中 2-審核通過 3-審核不通過
	Created_at     time.Time     `gorm:"column:created_at"`
	Updated_at     time.Time     `gorm:"column:updated_at"`
}

func (c ContentDetail) TableName() string { //指定此結構對應到哪個table
	return "cms_account.t_content_details"
}

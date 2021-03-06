package schema

type Article struct {
	Schema
	ChannelId         int              `gorm:"column:channel_id;type=int;not null" json:"channelId"`
	Title             string           `gorm:"column:title;type=varchar(255);not null;default \"\"" json:"title"`
	Author            string           `gorm:"column:author;type=varchar(255);not null;default \"\"" json:"author"`
	Anchor            int              `gorm:"column:anchor;type=int;not null" json:"anchor"`
	During            int              `gorm:"column:during;type=int;default 0" json:"during"`
	PlayNumber        int              `gorm:"column:play_number;type=int;not null;default 0" json:"playNumber"`
	CoverImg          string           `gorm:"column:cover_img;type=varchar(255);not null;default \"\"" json:"coverImg"`
	Audio             string           `gorm:"column:audio;type=varchar(255);not null;default \"\"" json:"audio"`
	TagIds            string           `gorm:"column:tag_ids;type=varchar(255);not null;default \"\"" json:"tagText"`
	ContentText       string           `gorm:"column:content_text;type=text;not null" json:"contentText"`
}


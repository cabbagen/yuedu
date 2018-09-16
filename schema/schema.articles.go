package schema

type Article struct {
	Schema
	ChannelId         int              `grom:"column:channel_id;type=int;not null" json:"channelId"`
	Title             string           `grom:"column:title;type=varchar(255);not null;default \"\"" json:"title"`
	Author            string           `grom:"column:author;type=varchar(255);not null;default \"\"" json:"author"`
	Anchor            int              `gorm:"column:anchor;type=int;not null" json:"anchor"`
	PlayNumber        int              `grom:"column:play_numnber;type=int;not null;default 0" json:"playNumber"`
	CoverImg          string           `gorm:"column:cover_img;type=varchar(255);not null;default \"\"" json:"coverImg"`
	Audio             string           `gorm:"column:audio;type=varchar(255);not null;default \"\"" json:"audio"`
	TagText           string           `gorm:"column:tag_text;type=varchar(255);not null;default \"\"" json:"tagText"`
	ContentText       string           `gorm:"column:content_text;type=text;not null" json:"contentText"`
}


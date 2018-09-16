package schema

type Relation struct {
	Schema
	UserId            int                `gorm:"column:user_id;type=int;not null" json:"userId"`
	RelationUserId    int                `gorm:"column:relation_user_id;type=int;not null" json:"relationUserId"`
	RelationType      int8               `gorm:"column:relation_type;type=tinyint" json:"relationType"`
}

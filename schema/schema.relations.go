package schema

// 粉丝、关注 关系表说明
// ===========================================================
// userId                 用户自身

// RelationUserId         和用户相关的另一个用户

// RelationType           两个用户之间的关系
//
//                           1 =》userId 关注 RelationUserId
//                           2 =》RelationUserId 关注 userId
//                           3 =》RelationUserId、userId 相互关注


type Relation struct {
	Schema
	UserId            int                `gorm:"column:user_id;type=int;not null" json:"userId"`
	RelationUserId    int                `gorm:"column:relation_user_id;type=int;not null" json:"relationUserId"`
	RelationType      int8               `gorm:"column:relation_type;type=tinyint" json:"relationType"`
}

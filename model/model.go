package model

type Page struct {
	PageIndex *int `json:"page_index"`
	PageSize  *int `json:"page_size"`
	Total     *int `json:"total"`
}

type VideoResult struct {
	VideoId *int64        `bson:"_id"`
	Content *VideoContent `bson:"content"`
}

type VideoContent struct {
	Video      *Video       `bson:"video"`
	Statistics *Statistics  `bson:"statistics"`
	Author     *Author      `bson:"author"`
	TextExtra  []*TextExtra `bson:"text_extra"`
	Region     *string      `bson:"region"`
	Desc       *string      `bson:"desc"`
	CreateTime *int32       `bson:"create_time"`
}
type Video struct {
	PlayAddr *Image `bson:"play_addr"`
	VideoUrl *Image `bson:"cover"`
	Duration *int32 `bson:"duration"`
}

type TextExtra struct {
	HashTagId   *string `bson:"hashtag_id"`
	End         *int32  `bson:"end"`
	Start       *int32  `bson:"start"`
	HashTagName *string `bson:"hashtag_name"`
	Type        *int    `bson:"type"`
	IsCommerce  *bool   `bson:"is_commerce"`
}
type Image struct {
	Width   *int      `bson:"width"`
	Height  *int      `bson:"height"`
	UrlList []*string `bson:"url_list"`
}
type Statistics struct {
	ForwardCount     *int64 `bson:"forward_count"`
	LoseCommentCount *int32 `bson:"lose_comment_count"`
	DiggCount        *int32 `bson:"digg_count"`
	PlayCount        *int32 `bson:"play_count"`
	CommentCount     *int64 `bson:"comment_count"`
	AwemeId          *int32 `bson:"aweme_id"`
	LoseCount        *int32 `bson:"lose_count"`
	ShareCount       *int32 `bson:"share_count"`
	DownloadCount    *int32 `bson:"download_count"`
}
type Author struct {
	UserId        *string  `bson:"uid"`
	NickName      *string  `bson:"nickname"`
	AvatarLarger  *Image   `bson:"avatar_larger"`
	AvatarMedium  *Image   `bson:"avatar_medium"`
	AvatarThumb   *Image   `bson:"avatar_thumb"`
	Avatar168x168 *Image   `bson:"avatar_168x168"`
	Avatar300x300 *Image   `bson:"avatar_300x300"`
	CoverUrl      []*Image `bson:"cover_url"`
	UniqueId      *string  `bson:"unique_id"`
	Signature     *string  `bson:"signature"`
}

type TicUser struct {
	UserId      *int64  `bson:"user_id"`
	UniqueId    *string `bson:"unique_id"`
	NickName    *int64  `bson:"nick_name"`
	Region      *string `bson:"region"`
	Avatar      *string `bson:"avatar"`
	ItemCount   *int32  `bson:"item_count"`
	FansCount   *int32  `bson:"fans_count"`
	FollowCount *int32  `bson:"following_count"`
	LikeCount   *int32  `bson:"like_count"`
	Language    *string `bson:"language"`
}

type HashTag struct {
	Tag        *string `bson:"_id"`
	VideoCount *int64  `bson:"count"`
	Views      *int64  `bson:"views"`
}

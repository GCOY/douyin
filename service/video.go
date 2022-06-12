package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/model"
)

func SelectAllVideoByFavorite(userID int64) []model.Video{
	videos := dao.SelectAllVideoByPublishTimeDesc()
	favorites := dao.SelectAllVideoByID(userID)
	followUsers := dao.GetFollow(userID)
	for i:=0 ;i< len(videos); i++ {
		for j:=0 ;j< len(favorites);j++{
			if videos[i].Id == favorites[j].VideoID && favorites[j].FavoriteStatus == true{
				videos[i].IsFavorite = true
			}
		}
		for k := 0;k< len(followUsers); k++ {
			if followUsers[k].Id == videos[i].Author.Id {
				videos[i].Author.IsFollow = true
			}
		}
	}
	return videos
}
func SelectAllVideo() []model.Video {
	return dao.SelectAllVideoByPublishTimeDesc()
}

func SaveVideo(userid int64, filename string,cover string, nowTime int64) {
	dao.SaveVideo(userid,filename,cover,nowTime)
}


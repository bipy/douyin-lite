package schedules

func init() {
	go UpdateCommentCount()
	go UpdateFavoriteCount()
	go UpdateFollowCount()
	go UpdateFollowerCount()
}

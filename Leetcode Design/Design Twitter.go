// userID tweetID followerID followed ID
type data struct {
	TweetID int
	UserID  int
}

type Twitter struct {
	m        map[int]map[int]bool // <followerID, userID , < followedId , bool>>
	database []data
}

func Constructor() Twitter {
	return Twitter{
		m:        make(map[int]map[int]bool),
		database: make([]data, 0),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	datanode := data{TweetID: tweetId, UserID: userId}
	this.database = append(this.database, datanode)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	feed := make([]int, 0)
	fmt.Println(this.database)
	for i := len(this.database) - 1; i >= 0; i-- {
		_, ok := this.m[userId][this.database[i].UserID]
		if ok || this.database[i].UserID == userId {
			feed = append(feed, this.database[i].TweetID)
		}
		if len(feed) == 10 {
			return feed
		}
	}
	return feed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.m[followerId]; ok == false {
		this.m[followerId] = map[int]bool{}
	}
	this.m[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.m[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
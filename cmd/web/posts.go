package main

type post struct {
	Title       string `xml:"title"`
	Slug        string `xml:"-"`
	Description string `xml:"description"`
	Date        string `xml:"-"`
	RSSLink     string `xml:"link"`
	RSSPubDate  string `xml:"pubDate"`
}

func newPostsData() map[string]post {
	posts := map[string]post{}
	posts["behind-the-scenes"] = post{
		Title:       "Behind the Scenes: Building This Website",
		Slug:        "behind-the-scenes",
		Description: `I'm peeling back the curtain to give you a peek into the technologies and decisions that brought this site to life.`,
		Date:        "November, 2024",
		RSSPubDate:  "2024-11-12T08:00:00Z",
	}
	return posts
}

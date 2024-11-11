package main

type post struct {
	Title       string
	Slug        string
	Description string
	Date        string
}

func newPostsData() map[string]post {
	posts := map[string]post{}
	posts["behind-the-scenes"] = post{
		Title: "Behind the Scenes: Building This Website",
		Slug:  "behind-the-scenes",
		Description: `I'm peeling back the
        curtain to give you a peek into the technologies and decisions that
        brought this site to life.`,
		Date: "November, 2024",
	}
	return posts
}

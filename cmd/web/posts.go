package main

type post struct {
	Title       string
	Slug        string
	Description string
	Date        string
}

func newPostsData() map[string]post {
	posts := map[string]post{}
	posts["post-design-system-something-test"] = post{
		Title:       "Crafting a design system for a multiplanetary future",
		Slug:        "post-design-system-something-test",
		Description: `Most companies try to stay ahead of the curve when it comes to visual design, but for Planetaria we needed to create a brand that would still inspire us 100 years from now when humanity has spread across our entire solar system.`,
		Date:        "September 5, 2022",
	}
	posts["tester-post"] = post{
		Title:       "TESTER POST",
		Slug:        "tester-post",
		Description: `Most companies try to stay ahead of the curve when it comes to visual design, but for Planetaria we needed to create a brand that would still inspire us 100 years from now when humanity has spread across our entire solar system.`,
		Date:        "Oktober 15, 2022",
	}

	return posts
}

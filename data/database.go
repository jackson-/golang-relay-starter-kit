package data

// Data model structs
type Post struct {
	Id   string `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
	Author string `json:"author"`
}

type PostList struct {
	Id   string `json:"id"`
	Posts []Post `json:"posts"`
}

// Mock data
var latestPost = &Post{
	Id: "1",
	Title: "Illest Article Alive",
	Text: "This some article about something",
	Author:"Devin Jackson",
}

var allPosts = &PostList{
		Id:"1",
		Posts:[]Post {
		{
			Id: "1",
			Title: "Illest Article Alive",
			Text: "This some article about some ill stuff",
			Author:"Devin Jackson",
		},
		{
			Id: "2",
			Title: "Dopest Article Alive",
			Text: "This some article about dope stuff",
			Author:"Devin Jackson",
		},
		{
			Id: "3",
			Title: "Newest Article Alive",
			Text: "This some article about new stuff",
			Author:"Devin Jackson",
		},
	},
}

// Data getters/setters
func GetLatestPost() *Post {
	return latestPost
}

func GetAllPosts() *PostList{
	return allPosts
}

package groups

import (
	"strconv"

	//"github.com/gnolang/gno/examples/gno.land/p/avl"
	"gno.land/p/avl"
)

type Group struct {
	Name     string
	Posts    *avl.Tree // postsCtr -> *Post
	PostsCtr int
}

func (group *Group) AddPost(title string, body string) {
	ctr := group.PostsCtr
	group.PostsCtr++
	key := strconv.Itoa(ctr)
	post := &Post{
		Title: title,
		Body:  body,
	}
	posts2, _ := group.Posts.Set(key, post)
	group.Posts = posts2
}

func (group *Group) String() string {
	str := "# [group] " + group.Name + "\n"
	if group.Posts.Size() > 0 {
		group.Posts.Traverse(true, func(n *avl.Tree) bool {
			str += "\n"
			str += n.Value().(*Post).String()
			return false
		})
	}
	return str
}

type Post struct {
	Title    string
	Body     string
	Comments *avl.Tree
}

func (post *Post) String() string {
	str := "## " + post.Title + "\n"
	str += ""
	str += post.Body
	if post.Comments.Size() > 0 {
		post.Comments.Traverse(true, func(n *avl.Tree) bool {
			str += "\n"
			str += n.Value().(*Comment).String()
			return false
		})
	}
	return str
}

type Comment struct {
	Creator string
	Body    string
}

func (cmm Comment) String() string {
	return cmm.Body + " - @" + cmm.Creator + "\n"
}
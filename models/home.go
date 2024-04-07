package models

import "go-blog/config"

type HomeResponse struct {
	config.Viewer
	Categories []Category
	Posts      []PostMore
	Total      int
	Page       int
	Pages      []int
	PageEnd    bool
}

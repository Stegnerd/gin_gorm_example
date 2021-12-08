package main

import api "github.com/stegnerd/gin_gorm_example/api"

func main() {
	r := api.GetRouter()
	api.AddPostRoutes(r)
	err := r.Run()
	if err != nil {
		return
	}
}

package models

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

func GetImages() []string {
	var images []string
	client, err := getclient()
	if err != nil {
		fmt.Println(err.Error())
		return []string{}
	} else {
		opts := docker.ListImagesOptions{}
		Images, err := client.ListImages(opts)
		if err != nil {
			fmt.Println(err.Error())
			return []string{}
		}else {
			for _, image := range Images {
				images = append(images,image.RepoTags[0])
			}
			return images
		}
	}
}

package models

import (
	"fmt"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

type Container struct {
	ContainerName   string
	ContainerID     string
	Image           string
	ContainerStatus string
}

func getclient() (*docker.Client, error) {
	//endpoint := "tcp://" + ip + ":55555"
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		return nil, err
	} else {
		return client, nil
	}
}
func CreateContainer(container_name, image_name, env string) error {
	client, err := getclient()
	if err != nil {
		return err
	}
	var envs []string
	if env != "" {
		envs = strings.Split(env, ",")
	}
	config := &docker.Config{Image: image_name, Env: envs, Tty: true}
	hostconfig := &docker.HostConfig{NetworkMode: "host"}
	cont, err := client.CreateContainer(docker.CreateContainerOptions{Name: container_name, Config: config, HostConfig: hostconfig})
	if err != nil {
		return err
	} else {
		//fmt.Println(cont.ID)
		//fmt.Println(cont.Image)
		err = client.StartContainer(cont.ID, hostconfig)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

func StartContainer(container_id string) error {
	client, err := getclient()
	if err != nil {
		return err
	}
	hostconfig := &docker.HostConfig{}
	err = client.StartContainer(container_id, hostconfig)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func StopContainer(container_id string) error {
	client, err := getclient()
	if err != nil {
		return err
	}
	err = client.StopContainer(container_id, 10)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func DelContainer(container_id string) error {
	client, err := getclient()
	if err != nil {
		fmt.Println(err)
	}
	opts := docker.RemoveContainerOptions{ID: container_id}
	err = client.RemoveContainer(opts)
	if err != nil {
		return err
	} else {
		//删除容器时，释放端口
		//DelPort()
		return nil
	}
}

func GetContainer() []Container {
	var containers []Container
	client, err := getclient()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		opts := docker.ListContainersOptions{All: true}
		conts, err := client.ListContainers(opts)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, cont := range conts {
				containers = append(containers, Container{cont.Names[0], cont.ID, cont.Image, cont.Status})
			}
		}
	}
	return containers
}
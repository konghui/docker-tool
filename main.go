package main

import (
	"github.com/konghui/docker-tool/pkg/pull"
	"github.com/konghui/docker-tool/pkg/push"
)

func main() {
	var registry string = ""
	var repository string = ""
	var tag string = ""
	var user string = ""
	var passwd string = ""
	pull.NewImageSource(registry, repository, tag, user, passwd, true)
	push.NewImageDestination(registry, repository, tag, user, passwd, true)
}

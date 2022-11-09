package services

import (
	"github.com/corentings/stouuf-bot/commands/cours"
	"github.com/corentings/stouuf-bot/commands/karma"
	"sync"
)

type IServiceContainer interface {
	InjectKarmaCommandHandler() karma.KarmaCommand
	InjectCoursCommandHandler() cours.CoursCommand
}

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func GetServiceContainer() IServiceContainer {
	containerOnce.Do(func() {
		k = &kernel{}
	})
	return k
}

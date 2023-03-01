package socialfit

import (
	"github.com/diegocomunity/goschool/examples/socialfit/apps/frontoffice"
)

type socialfit struct{}

func New() *socialfit {
	return &socialfit{}
}

func (*socialfit) Run() {
	frontoffice.Bootstrap()
}

package config

import (
	"os"

	"github.com/gookit/color"
	"github.com/kubecolor/kubecolor/testutil"
)

func init() {
	os.Clearenv()
	color.ForceColor()
	color.Enable = true

	testutil.DiffAllowUnexported(Color{})
}

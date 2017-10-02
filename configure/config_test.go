package configure

import (
	"testing"

	"github.com/daviddexter/wajibu/handlers/types"
)

func TestUpdater(t *testing.T) {
	var u types.ConfigUpdater
	u.Path = "Deployed"
	u.Value = "true"
	d := Updater(&u)
	f := Loader()
	if f.Deployed != true {
		t.Error("Expected true, got ", f.Deployed)
	}

}

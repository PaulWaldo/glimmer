package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_contactPhotos_makeUI(t *testing.T) {
	cp := contactPhotos{}
	ui := cp.makeUI()
	assert.Empty(t, ui.Objects[0])
}

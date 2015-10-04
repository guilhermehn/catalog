package catalog

import (
	"testing"
)

func TestCatalog(t *testing.T) {
	testDir := "./test_files"
	images, err := Catalog(testDir, ".png")

	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}

	if len(images) < 2 {
		t.Error("Not all images were cataloged")
	}

	someExt := []string{".png", ".txt"}
	someFiles, err := Catalog(testDir, someExt...)

	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}

	if len(someFiles) < 3 {
		t.Error("Some files were not cataloged")
	}

	allFiles, err := Catalog(testDir)

	if err != nil {
		t.Errorf("ERROR: %s", err.Error())
	}

	if len(allFiles) != 4 {
		t.Error("Some files were not cataloged")
	}
}

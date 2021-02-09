package usecase

import (
	"testing"

	"github.com/Tatsuemon/anony/testutils"
)

func TestMain(m *testing.M) {
	testDB := testutils.PrepareTestDB()
	testutils.SetTestDB(testDB)
	m.Run()
	defer testutils.CloseTestDB()
}

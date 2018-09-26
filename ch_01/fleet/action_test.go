package fleet

import "testing"

func TestDestroyAstronomicalObject(t *testing.T) {
	expected := "187J3X1已被光粒毁灭。"

	actual, err := DestroyAstronomicalObject("187J3X1", "光粒")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("unexpected output: %v (actual) != %v (expected)", actual, expected)
	}
}

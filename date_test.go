package date

import (
	"testing"
)

// 閏年判定のテスト
func TestIsLeapYear(t *testing.T) {
	t.Run("2000年", func(t *testing.T) {
		t.Log("2000年")
		expected := true
		actual := IsLeapYear(2000)
		if actual != expected {
			t.Error("\n理想： ", expected, "\n実際： ", actual)
		}
	  })
	t.Run("2001年", func(t *testing.T) {
		t.Log("2001年")
		expected := false
		actual := IsLeapYear(2001)
		if actual != expected {
			t.Error("\n理想： ", expected, "\n実際： ", actual)
		}
	  })
	  t.Run("2004年", func(t *testing.T) {
		t.Log("2004年")
		expected := true
		actual := IsLeapYear(2004)
		if actual != expected {
			t.Error("\n理想： ", expected, "\n実際： ", actual)
		}
	  })
	t.Run("2100年", func(t *testing.T) {
		t.Log("2100年")
		expected := false
		actual := IsLeapYear(2100)
		if actual != expected {
			t.Error("\n理想： ", expected, "\n実際： ", actual)
		}
	  })
}
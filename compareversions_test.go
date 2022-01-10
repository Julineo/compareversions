package main

import "testing"

func TestCompareVersions(t *testing.T) {
	tests := []struct{
		version1 string
		version2 string
		result int
	}{
		{"0.1", "1.1", -1},
		{"1.1", "0.1", 1},
		{"1.1", "1.1", 0},
		{"1.2", "1.2.9.9.9.9", -1},
		{"1.2.9.9.9.9", "1.3", -1},
		{"1.3.4", "1.3", 1},
		{"1.10", "1.3.4", 1},
	}

	for _, test := range tests {
		result, err := compareVersions(test.version1, test.version2)
		if result != test.result && err == nil {
			t.Errorf("compareVesions of %v and %v is incorrect, got: %d, want: %d.", test.version1, test.version2, result, test.result)
		}
		if err != nil {
			t.Errorf("compareVersions of %v and %v returned unexpected error", test.version1, test.version2)
		}
	}
}

func TestCompareVersionsError(t *testing.T) {
	result, err := compareVersions("0.1", "w.1")
	if result != 0 {
		t.Errorf("compareVersions of %v and %v must return %d as result", "0.1", "w.1", 0)
	}
	if err == nil {
		t.Errorf("compareVersions of %v and %v must return non empty error", "0.1", "w.1")
	}
}
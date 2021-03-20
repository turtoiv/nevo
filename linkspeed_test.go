package main

import "testing"

func TestSpeedOokla(t *testing.T) {
	err := StartSpeedTest(Ookla)
	if err != nil {
		t.Errorf("speed test failed. Expected nil got %s\n", err.Error())
	}
}

func TestSpeedNetflix(t *testing.T) {
	err := StartSpeedTest(Netflix)
	if err != nil {
		t.Errorf("speed test failed. Expected nil got %s\n", err.Error())
	}
}

func TestGetLinkSpeed(t *testing.T) {
	a := new(OoklaProvider)
	a.getLinkSpeed()
	if a.getDownloadSpeed() == 0 || a.getUploadSpeed() == 0 {
		t.Errorf("speed test failed for Ookla provider. Expected a value different from 0")
	}

	b := new(NetflixProvider)
	b.getLinkSpeed()
	if b.getDownloadSpeed() == 0 || b.getUploadSpeed() != 0 {
		t.Errorf("speed test failed for Netflix provider")
	}
}

func TestUnsupportedProvider(t *testing.T) {
	err := StartSpeedTest(3)
	if err.Error() != "Unknown provider" {
		t.Errorf("speed test failed. Expected Unknown provider, got %s\n", err.Error())
	}
}
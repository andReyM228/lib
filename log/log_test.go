package log

import "testing"

func TestLog(t *testing.T) {
	log := Init()
	log.Debug("qwerty")
	log.Info("qwerty")
	log.Error("qwerty")
	log.Debugf("qwerty: %s", "ytrewq")
	log.Infof("qwerty: %s", "ytrewq")
	log.Errorf("qwerty: %s", "ytrewq")
}

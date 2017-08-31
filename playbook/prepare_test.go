package playbook

import "testing"

func TestGetTemplatePath(t *testing.T) {
	tps, err := getTemplatePath("good-playbook")
	if err != nil {
		t.Error(err)
	}

	for _, tp := range tps {
		t.Log(tp.Src)
		t.Log(tp.Dest)
	}

	_, err = getTemplatePath("bad0-playbook")
	if err == nil {
		t.Error("bad-playbook should have error")
	} else {
		t.Log(err)
	}

	_, err = getTemplatePath("bad1-playbook")
	if err == nil {
		t.Error("bad-playbook should have error")
	} else {
		t.Log(err)
	}

	_, err = getTemplatePath("bad2-playbook")
	if err == nil {
		t.Error("bad-playbook should have error")
	} else {
		t.Log(err)
	}
}
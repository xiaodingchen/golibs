package db

import "testing"

func Test_NewManager(t *testing.T) {
	config := fconfig()
	name1 := "master"
	name2 := "slave"
	cfgs := make(map[string]*Config)
	cfgs[name1] = config
	cfgs[name2] = config
	mgr := NewManager(cfgs)
	client1, err := mgr.Load(name1)
	if err != nil {
		t.Fatal("manager err", err)
	}

	if client1 == nil {
		t.Fatal("manager client nil")
	}

	client2, err := mgr.Load(name2)
	if err != nil {
		t.Fatal("manager err", err)
	}

	if client2 == nil {
		t.Fatal("manager client nil")
	}

	if client1 == client2 {
		t.Fatal("manager client fatal")
	}

	client3, err := mgr.Load(name1)
	if err != nil {
		t.Fatal("manager err", err)
	}

	if client3 == nil {
		t.Fatal("manager client nil")
	}

	if client1 != client3 {
		t.Fatal("manager client fatal")
	}

	t.Log("manager pass")
}

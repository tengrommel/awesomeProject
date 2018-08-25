package testcase02

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monstor.Store() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store() 测试成功！")
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.Restore() 错误，希望为=%v 实际为=%v", "红孩儿", monster.Name)
	}
}

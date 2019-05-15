package monster


import(
	"testing"
)

func TestStore(t *testing.T){
	monster := &Monster{
		Name : "brandon",
		Age : 20,
		Skill : "eating",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，期望=%v,实际=%v", true, res)
	}
	t.Logf("monster.Store()测试成功")
}

func TestRestore(t *testing.T){
	var monster  = &Monster{}
	res := monster.Restore()
	if !res {
		t.Fatalf("monster.Restore()错误，期望=%v,实际=%v", true, res)
	}

	if monster.Name != "brandon" {
		t.Fatalf("monster.Restore()错误，期望=%v,实际=%v", "brandon", monster.Name)
	}
	t.Logf("monster.Restore()测试成功，monster=%v", *monster)
}

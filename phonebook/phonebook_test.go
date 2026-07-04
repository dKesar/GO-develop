package phonebook

import "testing"

func TestAdd(t *testing.T) {
	pb := New()
	cases := []struct {
		name string
		phone string
	}{
		{name: "John Doe", phone: "1234567890"},
		{name: "Jane Doe", phone: "0987654321"},
		{name: "Tim Belan", phone: "1234567890"},
		{name: "Jim Halpert", phone: "0987654321"},
		{name: "Pam Beesly", phone: "1234567890"},
		{name: "Michael Scott", phone: "0987654321"},
		{name: "Dwight Schrute", phone: "1234567890"},
		{name: "Angela Martin", phone: "0987654321"},
		{name: "Kelly Kapoor", phone: "1234567890"},
		{name: "Ryan Howard", phone: "0987654321"},
	}
	
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := pb.Add(c.name, c.phone)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestFind(t *testing.T) {
	pb := New()
	
	pb.Add("John Doe", "1234567890")

	phone, err := pb.Find("John Doe")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if phone != "1234567890" {
		t.Errorf("Expected 1234567890, got %s", phone)
	}
}

func TestDuplicate(t *testing.T) {
    pb := New()
    
    // первый Add не должен давать ошибку — проверяем это тоже
    if err := pb.Add("John Doe", "1234567890"); err != nil {
        t.Fatalf("первый Add не должен давать ошибку: %v", err)
    }
    
    // второй Add с тем же именем — должна быть ошибка
    err := pb.Add("John Doe", "1234567890")
    if err == nil {
        t.Error("ожидали ошибку при дублировании, получили nil")
    }
}

func TestDelete(t *testing.T) {
	pb := New()
	pb.Add("John Doe", "1234567890")
	err := pb.Delete("John Doe")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = pb.Find("John Doe")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}

func TestCount(t *testing.T) {
	pb := New()
	pb.Add("John Doe", "1234567890")
	pb.Add("Jane Doe", "0987654321")
	pb.Add("Tim Belan", "1234567890")

	pb.Delete("Tim Belan")

	count := pb.Count()
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestEmptyName(t *testing.T) {
	pb := New()
	err := pb.Add("", "1234567890")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}


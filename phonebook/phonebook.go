package phonebook

import "fmt"

type PhoneBook struct {
    entries map[string]string
}

func New() *PhoneBook{
	return &PhoneBook{
		entries: make(map[string]string),
	}
}
func (add *PhoneBook) Add(name, phone string) error{
	if name == "" || phone == "" {
		return fmt.Errorf("name and phone are required")
	}

	_, ok := add.entries[name]
	if ok{
		return fmt.Errorf("name already exists")
	}

	add.entries[name] = phone
	return nil
}

func (find *PhoneBook) Find(name string) (string, error){
	phone, ok := find.entries[name]
	if !ok {
		return "", fmt.Errorf("name not found")
	}
	return phone, nil
}

func (del *PhoneBook) Delete(name string) error{
	_, ok := del.entries[name]
	if !ok {
		return fmt.Errorf("name not found")
	}
	delete(del.entries, name)
	return nil
}

func (count *PhoneBook) Count() int{
	return len(count.entries)
}

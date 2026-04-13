package domain

type User struct {
	ID      int
	Version int

	FullName    string  //FullName обязателен
	PhoneNumber *string //Для ниловых значений

}

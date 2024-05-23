package structs

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Role      string `json:"role"` // "husband" or "wife"
	FamilyID  int    `json:"familyId"`
}

// Структура для семьи
type Family struct {
	ID        int `json:"id"`
	HusbandID int `json:"husbandId"`
	WifeID    int `json:"wifeId"`
	Kisses    int `json:"kisses"`
	Debt      int `json:"debt"`
}

// Структура для блюда
type Dish struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Recipe      string `json:"recipe"`
	CookingTime int    `json:"cookingTime"` // In minutes
	Complexity  string `json:"complexity"`  // "easy", "medium", "hard"
	Taste       string `json:"taste"`       // "delicious", "good", "ok"
	Kisses      int    `json:"kisses"`      // Cost of the dish in kisses
}

// Структура для заказа блюда
type Order struct {
	ID         int    `json:"id"`
	DishID     int    `json:"dishId"`
	FamilyID   int    `json:"familyId"`
	Status     string `json:"status"`     // "pending", "cooking", "done"
	KissesPaid int    `json:"kissesPaid"` // Number of kisses paid for the dish
}

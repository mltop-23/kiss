package structs

// Структура для заказа блюда
type Order struct {
	ID         int    `json:"id"`         // Уникальный идентификатор заказа
	DishID     int    `json:"dishId"`     // Идентификатор блюда
	FamilyID   int    `json:"familyId"`   // Идентификатор семьи, сделавшей заказ
	Status     string `json:"status"`     // Статус заказа ("pending", "cooking", "done")
	KissesPaid int    `json:"kissesPaid"` // Количество "поцелуев", уплаченных за блюдо
}

//____________________________________________________________________________________________________________________

type User struct {
	ID        int    `json:"id"`
	FamilyID  int    `json:"familyId"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Role      string `json:"role"` // "husband" or "wife"
}

// Структура для семьи
type Family struct {
	ID        int `json:"id"`        // Уникальный идентификатор семьи
	HusbandID int `json:"husbandId"` // Идентификатор мужа
	WifeID    int `json:"wifeId"`    // Идентификатор жены
	Kisses    int `json:"kisses"`    // Количество "поцелуев" у семьи
	Debt      int `json:"debt"`      // Долг семьи в "поцелуях"
}

// Структура для блюда
type Dish struct {
	ID          int    `json:"id"`          // Уникальный идентификатор блюда
	Name        string `json:"name"`        // Название блюда
	Recipe      string `json:"recipe"`      // Рецепт блюда
	CookingTime int    `json:"cookingTime"` // Время приготовления в минутах
	Complexity  string `json:"complexity"`  // Сложность приготовления ("easy", "medium", "hard")
	Taste       string `json:"taste"`       // Вкус блюда ("delicious", "good", "ok")
	Kisses      int    `json:"kisses"`      // Cost of the dish in kisses
}

// Структура для приема пищи
type Meal struct {
	ID   int    `json:"id"`   // Уникальный идентификатор типа приема пищи
	Name string `json:"name"` // Название приема пищи ("breakfast", "lunch", "dinner")
}

// Структура для ПЛАНОВ приема пищи женщины
type MealPlanWoman struct {
	ID          int    `json:"id"`          // Уникальный идентификатор плана
	FamilyID    int    `json:"familyId"`    // Идентификатор семьи
	WhosePlanID int    `json:"whosePlanId"` // Идентификатор человека, назначившего план (жены)
	MealID      int    `json:"mealId"`      // Идентификатор типа приема пищи
	Date        string `json:"date"`        // Дата приема пищи
}

// Структура для ПЛАНОВ приема пищи мужчины
type MealPlanMan struct {
	ID          int    `json:"id"`          // Уникальный идентификатор плана
	FamilyID    int    `json:"familyId"`    // Идентификатор семьи
	WhosePlanID int    `json:"whosePlanId"` // Идентификатор человека, назначившего план (мужа)
	MealID      int    `json:"mealId"`      // Идентификатор типа приема пищи
	Date        string `json:"date"`        // Дата приема пищи
}

// Структура для блюд в плане приема пищи
type MealDish struct {
	ID         int `json:"id"`         // Уникальный идентификатор
	MealPlanID int `json:"mealPlanId"` // Идентификатор плана приема пищи
	DishID     int `json:"dishId"`     // Идентификатор блюда
}

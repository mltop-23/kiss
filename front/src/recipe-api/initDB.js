const sqlite3 = require('sqlite3').verbose();

let db = new sqlite3.Database('./bd/database.db');

db.serialize(() => {
  // Удаляем таблицы перед созданием новых
  db.run("DROP TABLE IF EXISTS user");
  db.run("DROP TABLE IF EXISTS family");
  db.run("DROP TABLE IF EXISTS calendar");
  db.run("DROP TABLE IF EXISTS week_menu");
  db.run("DROP TABLE IF EXISTS recipes");

  // Создаем таблицу пользователей
  db.run(`CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    family_id INTEGER,
    name TEXT,
    lastname TEXT,
    number INTEGER,
    email TEXT,
    password TEXT,
    gender TEXT,
    FOREIGN KEY(family_id) REFERENCES family(id)
  )`);

  // Создаем таблицу семей
  db.run(`CREATE TABLE IF NOT EXISTS family (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    token INTEGER,
    men_id INTEGER,
    women_id INTEGER,
    FOREIGN KEY(men_id) REFERENCES user(id),
    FOREIGN KEY(women_id) REFERENCES user(id)
  )`);

  // Создаем таблицу календаря
  db.run(`CREATE TABLE IF NOT EXISTS calendar (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date TIMESTAMP,
    family_id INTEGER,
    family_men_id INTEGER,
    family_women_id INTEGER,
    recipes_id INTEGER,
    FOREIGN KEY(family_id) REFERENCES family(id),
    FOREIGN KEY(family_men_id) REFERENCES family(men_id),
    FOREIGN KEY(family_women_id) REFERENCES family(women_id),
    FOREIGN KEY(recipes_id) REFERENCES recipes(id)
  )`);

  // Создаем таблицу рецептов
  db.run(`CREATE TABLE IF NOT EXISTS recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    src TEXT,
    name TEXT,
    description TEXT,
    ingredients TEXT,
    title TEXT,
    calories INTEGER,
    carbohydrates INTEGER,
    fats INTEGER
  )`);

  // Создаем таблицу недельного меню
  db.run(`CREATE TABLE IF NOT EXISTS week_menu (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    day TEXT,
    meal_type TEXT,
    recipe_id INTEGER,
    FOREIGN KEY(recipe_id) REFERENCES recipes(id)
  )`);

  // Вставляем данные в таблицу рецептов
  let stmtRecipes = db.prepare("INSERT INTO recipes (src, name, description, ingredients, title, calories, carbohydrates, fats) VALUES (?, ?, ?, ?, ?, ?, ?, ?)");
  const recipes = [
    { 'src': '/borscht.jpg', 'name': 'Борщ', 'description': 'Традиционный украинский суп с буряком', 'ingredients': 'Вода, мясо, буряк, картофель, капуста', 'title': 'Полный рецепт борща...', 'calories': 150, 'carbohydrates': 20, 'fats': 5 },
    { 'src': '/pelmeni.jpeg', 'name': 'Пельмени', 'description': 'Тесто с мясной начинкой, отваренное в воде', 'ingredients': 'Тесто, мясо, лук', 'title': 'Полный рецепт пельменей...', 'calories': 200, 'carbohydrates': 30, 'fats': 10 },
    { 'src': '/olivier.png', 'name': 'Салат Оливье', 'description': 'Классический салат с колбасой и овощами', 'ingredients': 'Картофель, морковь, зеленый горошек, колбаса, майонез', 'title': 'Полный рецепт салата Оливье...', 'calories': 180, 'carbohydrates': 15, 'fats': 12 },
    { 'src': '/blini.jpg', 'name': 'Блины', 'description': 'Тонкие русские блины с начинкой', 'ingredients': 'Мука, яйца, молоко, масло', 'title': 'Полный рецепт блинов...', 'calories': 250, 'carbohydrates': 35, 'fats': 10 },
    { 'src': '/shchi.jpeg', 'name': 'Щи', 'description': 'Русский суп из капусты и мяса', 'ingredients': 'Вода, капуста, мясо, картофель, морковь', 'title': 'Полный рецепт щей...', 'calories': 120, 'carbohydrates': 10, 'fats': 4 },
    { 'src': '/vareniki.jpg', 'name': 'Вареники', 'description': 'Тесто с картофельной или творожной начинкой', 'ingredients': 'Тесто, картофель, творог, масло', 'title': 'Полный рецепт вареников...', 'calories': 220, 'carbohydrates': 40, 'fats': 8 },
    { 'src': '/kulebyaka.jpg', 'name': 'Кулебяка', 'description': 'Слоеный пирог с мясом и рыбой', 'ingredients': 'Тесто, мясо, рыба, рис, яйца', 'title': 'Полный рецепт кулебяки...', 'calories': 300, 'carbohydrates': 25, 'fats': 15 },
    { 'src': '/borscht.jpg', 'name': 'Борщ', 'description': 'Традиционный украинский суп с буряком', 'ingredients': 'Вода, мясо, буряк, картофель, капуста', 'title': 'Полный рецепт борща...', 'calories': 150, 'carbohydrates': 20, 'fats': 5 },
    { 'src': '/pelmeni.jpeg', 'name': 'Пельмени', 'description': 'Тесто с мясной начинкой, отваренное в воде', 'ingredients': 'Тесто, мясо, лук', 'title': 'Полный рецепт пельменей...', 'calories': 200, 'carbohydrates': 30, 'fats': 10 },
    { 'src': '/olivier.png', 'name': 'Салат Оливье', 'description': 'Классический салат с колбасой и овощами', 'ingredients': 'Картофель, морковь, зеленый горошек, колбаса, майонез', 'title': 'Полный рецепт салата Оливье...', 'calories': 180, 'carbohydrates': 15, 'fats': 12 },
    { 'src': '/blini.jpg', 'name': 'Блины', 'description': 'Тонкие русские блины с начинкой', 'ingredients': 'Мука, яйца, молоко, масло', 'title': 'Полный рецепт блинов...', 'calories': 250, 'carbohydrates': 35, 'fats': 10 },
    { 'src': '/shchi.jpeg', 'name': 'Щи', 'description': 'Русский суп из капусты и мяса', 'ingredients': 'Вода, капуста, мясо, картофель, морковь', 'title': 'Полный рецепт щей...', 'calories': 120, 'carbohydrates': 10, 'fats': 4 },
    { 'src': '/vareniki.jpg', 'name': 'Вареники', 'description': 'Тесто с картофельной или творожной начинкой', 'ingredients': 'Тесто, картофель, творог, масло', 'title': 'Полный рецепт вареников...', 'calories': 220, 'carbohydrates': 40, 'fats': 8 },
    { 'src': '/borscht.jpg', 'name': 'Борщ', 'description': 'Традиционный украинский суп с буряком', 'ingredients': 'Вода, мясо, буряк, картофель, капуста', 'title': 'Полный рецепт борща...', 'calories': 150, 'carbohydrates': 20, 'fats': 5 },
    { 'src': '/pelmeni.jpeg', 'name': 'Пельмени', 'description': 'Тесто с мясной начинкой, отваренное в воде', 'ingredients': 'Тесто, мясо, лук', 'title': 'Полный рецепт пельменей...', 'calories': 200, 'carbohydrates': 30, 'fats': 10 },
    { 'src': '/olivier.png', 'name': 'Салат Оливье', 'description': 'Классический салат с колбасой и овощами', 'ingredients': 'Картофель, морковь, зеленый горошек, колбаса, майонез', 'title': 'Полный рецепт салата Оливье...', 'calories': 180, 'carbohydrates': 15, 'fats': 12 },
    { 'src': '/blini.jpg', 'name': 'Блины', 'description': 'Тонкие русские блины с начинкой', 'ingredients': 'Мука, яйца, молоко, масло', 'title': 'Полный рецепт блинов...', 'calories': 250, 'carbohydrates': 35, 'fats': 10 },
    { 'src': '/shchi.jpeg', 'name': 'Щи', 'description': 'Русский суп из капусты и мяса', 'ingredients': 'Вода, капуста, мясо, картофель, морковь', 'title': 'Полный рецепт щей...', 'calories': 120, 'carbohydrates': 10, 'fats': 4 },
    { 'src': '/vareniki.jpg', 'name': 'Вареники', 'description': 'Тесто с картофельной или творожной начинкой', 'ingredients': 'Тесто, картофель, творог, масло', 'title': 'Полный рецепт вареников...', 'calories': 220, 'carbohydrates': 40, 'fats': 8 },
    { 'src': '/kulebyaka.jpg', 'name': 'Кулебяка', 'description': 'Слоеный пирог с мясом и рыбой', 'ingredients': 'Тесто, мясо, рыба, рис, яйца', 'title': 'Полный рецепт кулебяки...', 'calories': 300, 'carbohydrates': 25, 'fats': 15 },
    { 'src': '/borscht.jpg', 'name': 'Борщ', 'description': 'Традиционный украинский суп с буряком', 'ingredients': 'Вода, мясо, буряк, картофель, капуста', 'title': 'Полный рецепт борща...', 'calories': 150, 'carbohydrates': 20, 'fats': 5 },
    { 'src': '/pelmeni.jpeg', 'name': 'Пельмени', 'description': 'Тесто с мясной начинкой, отваренное в воде', 'ingredients': 'Тесто, мясо, лук', 'title': 'Полный рецепт пельменей...', 'calories': 200, 'carbohydrates': 30, 'fats': 10 },
    { 'src': '/olivier.png', 'name': 'Салат Оливье', 'description': 'Классический салат с колбасой и овощами', 'ingredients': 'Картофель, морковь, зеленый горошек, колбаса, майонез', 'title': 'Полный рецепт салата Оливье...', 'calories': 180, 'carbohydrates': 15, 'fats': 12 },
    { 'src': '/blini.jpg', 'name': 'Блины', 'description': 'Тонкие русские блины с начинкой', 'ingredients': 'Мука, яйца, молоко, масло', 'title': 'Полный рецепт блинов...', 'calories': 250, 'carbohydrates': 35, 'fats': 10 },
    { 'src': '/shchi.jpeg', 'name': 'Щи', 'description': 'Русский суп из капусты и мяса', 'ingredients': 'Вода, капуста, мясо, картофель, морковь', 'title': 'Полный рецепт щей...', 'calories': 120, 'carbohydrates': 10, 'fats': 4 },
    { 'src': '/vareniki.jpg', 'name': 'Вареники', 'description': 'Тесто с картофельной или творожной начинкой', 'ingredients': 'Тесто, картофель, творог, масло', 'title': 'Полный рецепт вареников...', 'calories': 220, 'carbohydrates': 40, 'fats': 8 }
  ];
  
  recipes.forEach((recipe) => {
    stmtRecipes.run(recipe.src, recipe.name, recipe.description, recipe.ingredients, recipe.title, recipe.calories, recipe.carbohydrates, recipe.fats);
  });

  stmtRecipes.finalize();

  // Вставляем данные в таблицу недельного меню
  const weekMenu = [
    { day: 'Понедельник', meal_type: 'breakfast', recipe_id: 1 },
    { day: 'Понедельник', meal_type: 'lunch', recipe_id: 2 },
    { day: 'Понедельник', meal_type: 'dinner', recipe_id: 3 },
    { day: 'Вторник', meal_type: 'breakfast', recipe_id: 4 },
    { day: 'Вторник', meal_type: 'lunch', recipe_id: 5 },
    { day: 'Вторник', meal_type: 'dinner', recipe_id: 6 },
    { day: 'Среда', meal_type: 'breakfast', recipe_id: 7 },
    { day: 'Среда', meal_type: 'lunch', recipe_id: 8 },
    { day: 'Среда', meal_type: 'dinner', recipe_id: 9 },
    { day: 'Четверг', meal_type: 'breakfast', recipe_id: 10 },
    { day: 'Четверг', meal_type: 'lunch', recipe_id: 11 },
    { day: 'Четверг', meal_type: 'dinner', recipe_id: 12 },
    { day: 'Пятница', meal_type: 'breakfast', recipe_id: 13 },
    { day: 'Пятница', meal_type: 'lunch', recipe_id: 14 },
    { day: 'Пятница', meal_type: 'dinner', recipe_id: 15 },
    { day: 'Суббота', meal_type: 'breakfast', recipe_id: 16 },
    { day: 'Суббота', meal_type: 'lunch', recipe_id: 17 },
    { day: 'Суббота', meal_type: 'dinner', recipe_id: 18 },
    { day: 'Воскресенье', meal_type: 'breakfast', recipe_id: 19 },
    { day: 'Воскресенье', meal_type: 'lunch', recipe_id: 20 },
    { day: 'Воскресенье', meal_type: 'dinner', recipe_id: 21 },
  ];
  

  let stmtWeekMenu = db.prepare("INSERT INTO week_menu (day, meal_type, recipe_id) VALUES (?, ?, ?)");
  weekMenu.forEach((menu) => {
    stmtWeekMenu.run(menu.day, menu.meal_type, menu.recipe_id);
  });

  stmtWeekMenu.finalize();

  // Вставляем данные в таблицы пользователей и семей
  db.run(`INSERT INTO family (token, men_id, women_id) VALUES (123456, 1, 2)`);
  db.run(`INSERT INTO user (family_id, name, lastname, number, email, password, gender) VALUES 
    (1, 'Иван', 'Иванов', 1234567890, 'ivanov@example.com', 'password123', 'мужчина'),
    (1, 'Мария', 'Иванова', 0987654321, 'ivanova@example.com', 'password123', 'женщина')`);
});

db.close();

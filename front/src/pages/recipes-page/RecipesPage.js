/*
Пример результата запроса к API:
{
  "recipes": [
    {
      "id": 1,
      "src": "/borscht.jpg",
      "name": "Борщ",
      "description": "Традиционный украинский суп с буряком"
    }
  ]
}

Пример кода с комментариями:
*/

// Импортируем необходимые библиотеки и стили
import React, { useState, useEffect } from 'react'; // Импортируем React и хуки useState и useEffect
import { Link } from 'react-router-dom'; // Импортируем компонент Link из react-router-dom для создания ссылок
import axios from 'axios'; // Импортируем axios для выполнения HTTP-запросов
import './RecipesPage.css'; // Импортируем стили для компонента

// Основной функциональный компонент страницы рецептов
const RecipesPage = () => {
  const [recipes, setRecipes] = useState([]); // Состояние для хранения списка рецептов

  // useEffect с пустым массивом зависимостей означает, что этот эффект выполнится один раз после монтирования компонента
  useEffect(() => {
    // Выполняем GET-запрос для получения списка рецептов
    axios.get('http://localhost:3001/api/recipes')
      .then(response => {
        // Устанавливаем полученные рецепты в состояние
        setRecipes(response.data.recipes); // Пример ответа от API: [{"id":1,"src":"/borscht.jpg","name":"Борщ","description":"Традиционный украинский суп с буряком"}]
      })
  }, []);

  return (
    // Основная обертка страницы рецептов
    <div className="recipes-page">
      {/* Сетка для отображения рецептов */}
      <div className="recipes-grid">
        {/* Мапируем список рецептов для создания карточек */}
        {recipes.map((recipe) => (
          // Компонент Link для создания ссылки на страницу рецепта
          <Link to={`/recipe/${recipe.id}`} key={recipe.id} className="recipe-card">
            {/* Изображение рецепта */}
            <img src={recipe.src} alt={recipe.name} className="recipe-image" /> {/* Пример: src="/borscht.jpg", alt="Борщ" */}
            {/* Заголовок рецепта */}
            <h2>{recipe.name}</h2>  {/* "Пример: "Борщ" */} 
            {/* Описание рецепта */}
            <p>{recipe.description}</p> {/* Пример: "Традиционный украинский суп с буряком" */} 
          </Link>
        ))}
      </div>
    </div>
  );
};

// Экспортируем компонент для использования в других частях приложения
export default RecipesPage;
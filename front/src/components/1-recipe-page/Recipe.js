// src/components/Recipe.js
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import './Recipe.css';

const Recipe = () => {
  const { id } = useParams();
  const [recipe, setRecipe] = useState(null);

  useEffect(() => {
    axios.get(`http://localhost:3001/api/recipes/${id}`)
      .then(response => {
        setRecipe(response.data);
      })
      .catch(error => {
        console.error('There was an error fetching the recipe!', error);
      });
  }, [id]);

  if (!recipe) {
    return <div>Loading...</div>;
  }

  return (
    <div className="recipe-container">
      <img src={recipe.src} alt={recipe.name} className="recipe-image" />
      <div className="recipe-details">
        <h1>{recipe.name}</h1>
        <p>{recipe.description}</p>
        <div className="recipe-section">
          <h2>ИНГРЕДИЕНТЫ</h2>
          <p>{recipe.ingredients}</p> {/* Комментарий: Текст для ингредиентов */}
        </div>
        <div className="recipe-section">
          <h2>ПОЛНЫЙ РЕЦЕПТ</h2>
          <p>{recipe.title}</p> {/* Комментарий: Полный рецепт */}
        </div>
        <div className="recipe-section">
          <h2>ПИЩЕВАЯ ЦЕННОСТЬ</h2>
          <div className="nutrition-table">
            <div>
              <span>Калории:</span>
              {recipe.calories}
            </div>
            <div>
              <span>Углеводы:</span>
              {recipe.carbohydrates}
            </div>
            <div>
              <span>Жиры:</span>
              {recipe.fats}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Recipe;

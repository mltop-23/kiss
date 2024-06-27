import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import './Gallery.css';

const Gallery = () => {
  const [recipes, setRecipes] = useState([]);

  useEffect(() => {
    const fetchRecipes = async () => {
      try {
        const response = await axios.get('http://localhost:3001/api/recipes');
        setRecipes(response.data.recipes);
      } catch (error) {
        console.error('Error fetching recipes:', error);
      }
    };

    fetchRecipes();
  }, []);

  return (
    <div className="gallery">
      {recipes.map((recipe) => (
        <Link to={`/recipe/${recipe.id}`} className="gallery-item" key={recipe.id}>
          <img src={recipe.src} alt={recipe.name} className="gallery-image" />
          <h3>{recipe.name}</h3>
          <p>{recipe.description}</p>
        </Link>
      ))}
    </div>
  );
};

export default Gallery;

// server.js
const express = require('express');
const sqlite3 = require('sqlite3').verbose();
const cors = require('cors');
const app = express();
const port = 3001;

app.use(cors());

let db = new sqlite3.Database('./bd/database.db');

app.get('/api/recipes', (req, res) => {
  db.all("SELECT id, src, name, description FROM recipes", [], (err, rows) => {
    if (err) {
      res.status(400).json({ "error": err.message });
      return;
    }
    res.setHeader('Content-Type', 'application/json');
    res.send(JSON.stringify({ "recipes": rows }, null, 2)); // Форматирование JSON с отступами в 2 пробела
  });
});

app.get('/api/week-menu', (req, res) => {
  db.all(`
    SELECT week_menu.day, week_menu.meal_type, recipes.id, recipes.src, recipes.title, recipes.description 
    FROM week_menu 
    JOIN recipes ON week_menu.recipe_id = recipes.id
  `, [], (err, rows) => {
    if (err) {
      res.status(400).json({ "error": err.message });
      return;
    }
    let weekMenu = rows.reduce((acc, row) => {
      if (!acc[row.day]) {
        acc[row.day] = { breakfast: [], lunch: [], dinner: [] };
      }
      acc[row.day][row.meal_type].push({
        id: row.id,
        src: row.src,
        title: row.title,
        description: row.description
      });
      return acc;
    }, {});
    res.setHeader('Content-Type', 'application/json');
    res.send(JSON.stringify({ "weekMenu": weekMenu }, null, 2)); // Форматирование JSON с отступами в 2 пробела
  });
});

app.get('/api/recipes/:id', (req, res) => {
  const { id } = req.params;
  db.get('SELECT * FROM recipes WHERE id = ?', [id], (err, row) => {
    if (err) {
      res.status(500).json({ error: err.message });
      return;
    }
    res.setHeader('Content-Type', 'application/json');
    res.send(JSON.stringify(row, null, 2)); // Форматирование JSON с отступами в 2 пробела
  });
});

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});

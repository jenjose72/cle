const express = require('express');
const app = express();

app.use(express.json());

let items = [{ id: 1, name: "Item1" }];

app.get('/items', (req, res) => {
  res.json(items);
});

app.post('/items', (req, res) => {
  items.push(req.body);
  res.json(req.body);
});

app.listen(3000, () => console.log("Server running"));
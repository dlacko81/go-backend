
const express = require('express');
const GO = require('../models/GO');
const router = express.Router();

router.post('/', async (req, res) => {
  try {
    const go = new GO(req.body);
    await go.save();
    res.status(201).json(go);
  } catch (err) {
    res.status(400).json({ error: err.message });
  }
});

router.get('/', async (req, res) => {
  const gos = await GO.find();
  res.json(gos);
});

router.get('/search', async (req, res) => {
  const { name } = req.query;
  const results = await GO.find({ name: new RegExp(name, 'i') });
  res.json(results);
});

router.get('/report', async (req, res) => {
  const report = await GO.aggregate([
    { $group: { _id: "$country", total: { $sum: "$volume" } } }
  ]);
  res.json(report);
});

module.exports = router;

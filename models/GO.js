const mongoose = require('mongoose');

const GOSchema = new mongoose.Schema({
  name: String,
  country: String,
  volume: Number,
  tech: String,
  date: Date
});

module.exports = mongoose.model('GO', GOSchema);

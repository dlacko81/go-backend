const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');
require('dotenv').config();

const goRoutes = require('./routes/goRoutes');

const app = express();
app.use(cors());
app.use(express.json());

app.use('/api/gos', goRoutes);

mongoose.connect(process.env.MONGO_URI)
  .then(() => {
    app.listen(process.env.PORT || 5000, () => {
      console.log('Backend running');
    });
  })
  .catch(err => console.log(err));

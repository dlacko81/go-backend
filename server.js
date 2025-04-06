const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');
require('dotenv').config();

const goRoutes = require('./routes/goRoutes');

const app = express();

// CORS setup to allow the frontend to access the backend
const corsOptions = {
  origin: 'https://go-frontend-chi.vercel.app',  // Replace with your Vercel frontend URL
  methods: ['GET', 'POST'],
  allowedHeaders: ['Content-Type'],
};

app.use(cors(corsOptions)); // Use the CORS middleware with the configured options

app.use(express.json());
app.use('/api/gos', goRoutes);

mongoose.connect(process.env.MONGO_URI)
  .then(() => {
    app.listen(process.env.PORT || 5000, () => {
      console.log('Backend running');
    });
  })
  .catch(err => console.log(err));

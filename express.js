const express = require('express');
const mongoose = require('mongoose');
const bodyParser = require('body-parser');
const app = express();
const port = 3000;

app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.static('public')); // Mengatur direktori statis

mongoose.connect('mongodb://localhost/regisGIS', {
  useNewUrlParser: true,
  useUnifiedTopology: true
});

const User = mongoose.model('User', {
  name: String,
  email: String,
  password: String
});

app.get('/', (req, res) => {
  // Tampilkan halaman pendaftaran
  res.sendFile(__dirname + '/index.html');
});

app.post('/register', (req, res) => {
    // Tangkap data dari formulir dan simpan ke MongoDB
    const { name, email, password, confirm } = req.body;
  
    if (password === confirm) {
      const user = new User({ name, email, password });
  
      user.save((err) => {
        if (err) {
          console.error(err);
          res.send('Gagal mendaftar');
        } else {
          res.redirect('/'); // Mengarahkan kembali ke halaman utama (index.html)
        }
      });
    } else {
      res.send('Password tidak cocok');
    }
  });
  
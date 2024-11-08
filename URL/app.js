// // Importing the mongoose library
// const mongoose = require('mongoose');

// // MongoDB connection URL (replace with your own connection string)
// const mongoURI = 'mongodb://localhost:27017/MERN-Test'; // Local MongoDB URI or use your MongoDB Atlas URI

// // Connect to MongoDB
// mongoose.connect(mongoURI, { useNewUrlParser: true, useUnifiedTopology: true })
//   .then(() => {
//     console.log('MongoDB connected successfully');
//   })
//   .catch((err) => {
//     console.error('MongoDB connection error:', err);
//   });

// // Optional: Create a simple schema and model
// const Schema = mongoose.Schema;
// const userSchema = new Schema({
//   name: String,
//   email: String
// });

// const User = mongoose.model('users', userSchema);

// // Create a sample user document (for testing)
// const newUser = new User({
//   name: 'John Doe',
//   email: 'john.doe@example.com'
// });

// newUser.save()
//   .then(() => {
//     console.log('User saved successfully');
//   })
//   .catch((err) => {
//     console.error('Error saving user:', err);
//   });

// Import Express module
const express = require('express');
const app = express();
const port = 3000;

// Middleware to parse JSON bodies
app.use(express.json());

// Sample data to work with
let data = [
    { id: 1, name: 'Item 1' },
    { id: 2, name: 'Item 2' }
];

// GET endpoint to fetch all items
app.get('/items', (req, res) => {
    res.json(data);
});

// POST endpoint to add a new item
app.post('/items', (req, res) => {
    const newItem = {
        id: data.length + 1,
        name: req.body.name
    };
    data.push(newItem);
    res.status(201).json(newItem);
});

// PUT endpoint to update an existing item by ID
app.put('/items/:id', (req, res) => {
    const itemId = parseInt(req.params.id, 10);
    const item = data.find(d => d.id === itemId);
    if (item) {
        item.name = req.body.name;
        res.json(item);
    } else {
        res.status(404).json({ message: 'Item not found' });
    }
});

// DELETE endpoint to remove an item by ID
app.delete('/items/:id', (req, res) => {
    const itemId = parseInt(req.params.id, 10);
    const itemIndex = data.findIndex(d => d.id === itemId);
    if (itemIndex !== -1) {
        data.splice(itemIndex, 1);
        res.status(204).send();
    } else {
        res.status(404).json({ message: 'Item not found' });
    }
});

// Start the server
app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});

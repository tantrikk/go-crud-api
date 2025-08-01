import React, { useState, useEffect } from 'react';
import './App.css';
import UserList from './components/UserList';
import UserForm from './components/UserForm';
import { userService } from './services/api';

function App() {
  const [users, setUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  // Fetch all users
  const fetchUsers = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await userService.getAllUsers();
      setUsers(data || []);
    } catch (err) {
      setError('Failed to fetch users');
      console.error('Error fetching users:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  // Create or update user
  const handleSaveUser = async (userData) => {
    try {
      setError(null);
      if (selectedUser) {
        // Update existing user
        await userService.updateUser(selectedUser.id, userData);
      } else {
        // Create new user
        await userService.createUser(userData);
      }
      setSelectedUser(null);
      fetchUsers();
    } catch (err) {
      setError('Failed to save user');
      console.error('Error saving user:', err);
    }
  };

  // Delete user
  const handleDeleteUser = async (id) => {
    if (window.confirm('Are you sure you want to delete this user?')) {
      try {
        setError(null);
        await userService.deleteUser(id);
        fetchUsers();
      } catch (err) {
        setError('Failed to delete user');
        console.error('Error deleting user:', err);
      }
    }
  };

  // Edit user
  const handleEditUser = (user) => {
    setSelectedUser(user);
  };

  // Cancel edit
  const handleCancelEdit = () => {
    setSelectedUser(null);
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>User Management System</h1>
      </header>
      
      <main className="App-main">
        {error && (
          <div className="error-message">
            {error}
          </div>
        )}
        
        <div className="container">
          <div className="form-section">
            <h2>{selectedUser ? 'Edit User' : 'Add New User'}</h2>
            <UserForm
              user={selectedUser}
              onSave={handleSaveUser}
              onCancel={handleCancelEdit}
            />
          </div>
          
          <div className="list-section">
            <h2>Users List</h2>
            {loading ? (
              <div className="loading">Loading...</div>
            ) : (
              <UserList
                users={users}
                onEdit={handleEditUser}
                onDelete={handleDeleteUser}
              />
            )}
          </div>
        </div>
      </main>
    </div>
  );
}

export default App;
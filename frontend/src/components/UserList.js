import React from 'react';
import './UserList.css';

function UserList({ users, onEdit, onDelete }) {
  if (users.length === 0) {
    return <div className="no-users">No users found. Add one above!</div>;
  }

  return (
    <div className="user-list">
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id}>
              <td>{user.name}</td>
              <td>{user.email}</td>
              <td className="actions">
                <button 
                  className="edit-btn" 
                  onClick={() => onEdit(user)}
                >
                  Edit
                </button>
                <button 
                  className="delete-btn" 
                  onClick={() => onDelete(user.id)}
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default UserList;
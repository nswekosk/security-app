import React, { useState } from 'react';

const Login: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleUsernameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // Fetch options
    const requestOptions: RequestInit = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    };
    // Send a request to the backend
    const response = await fetch('/api/login', requestOptions);
    if(response.ok) {
      const data = await response.json();
      // Save JWT to localStorage or state management
      localStorage.setItem('token', data.token);
      // Redirect to dashboard or file explorer
    } else {
      // Handle login error
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" value={username} onChange={handleUsernameChange} placeholder="Username" required />
      <input type="password" value={password} onChange={handlePasswordChange} placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
  );
};

export default Login;
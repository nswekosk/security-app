import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import jwt_decode from 'jwt-decode';
import AuthContext from './AuthContext';

type AuthProviderProps = {
    children: React.ReactNode;
};
  
const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [token, setToken] = useState<string | null>(() => sessionStorage.getItem('token'));
  const navigate = useNavigate();

  useEffect(() => {
    // If a token exists, check its expiration time
    if (token) {
      const decodedToken: any = jwt_decode(token);
      const currentTime = Date.now() / 1000;

      // If the token is expired, log out
      if (decodedToken.exp < currentTime) {
        logOut();
      }
    }
  }, [token]);

  const logIn = (token: string) => {
    sessionStorage.setItem('token', token);
    setToken(token);
  };

  const logOut = () => {
    sessionStorage.removeItem('token');
    setToken(null);
    navigate('/login');
  };

  return (
    <AuthContext.Provider value={{ token, logIn, logOut }}>
      {children}
    </AuthContext.Provider>
  );
};
export { AuthProvider };
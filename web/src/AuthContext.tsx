import React, { createContext, useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import jwt_decode from 'jwt-decode';

interface AuthContextProps {
  token: string | null;
  logIn: (token: string) => void;
  logOut: () => void;
}

const AuthContext = createContext<AuthContextProps | null>(null);



export default AuthContext;
'use client';

import React, { createContext, useState, useContext, useEffect } from 'react';

const AuthContext = createContext();

class User {
  constructor(name, email) {
    this.name = name;
    this.email = email;
  }
}

export function AuthProvider({ children }) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = localStorage.getItem('accessToken');
    if (token) {
      setIsLoggedIn(true);
      setUser(getUser(token));
    }
    setLoading(false);
  }, []);

  const getTokenPayload = (token) => {
    const base64Payload = token.split('.')[1];
    const payload = Buffer.from(base64Payload, 'base64');
    return JSON.parse(payload.toString());
  };

  const getUser = (token) => {
    const payload = getTokenPayload(token);
    return new User(payload.claims.name, payload.sub);
  };

  const login = (token, refreshToken) => {
    localStorage.setItem('accessToken', token);
    localStorage.setItem('refreshToken', refreshToken);
    setIsLoggedIn(true);
    setUser(getUser(token));
  };

  const logout = () => {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    setIsLoggedIn(false);
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, user, login, logout, loading}}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}
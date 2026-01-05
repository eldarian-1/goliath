import { createContext, useContext, useState, useEffect } from 'react';
import { fetchWithAuth } from '../helpers/fetch';

const AuthContext = createContext();

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchWithAuth('/api/v1/auth/me')
      .then(res => res.ok ? res.json() : null)
      .then(data => setUser(data))
      .finally(() => setLoading(false));
  }, []);

  const register = async (email, password) => {
    const res = await fetchWithAuth('/api/v1/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) throw new Error('Registration failed');

    const data = await fetchWithAuth('/api/v1/auth/me').then(r => r.json());
    setUser(data);
  };

  const login = async (email, password) => {
    const res = await fetchWithAuth('/api/v1/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) throw new Error('Login failed');

    const data = await fetchWithAuth('/api/v1/auth/me').then(r => r.json());
    setUser(data);
  };

  const logout = async () => {
    await fetchWithAuth('/api/v1/auth/logout', { method: 'POST' });
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ user, register, login, logout, loading }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext);

import { createContext, useContext, useState, useEffect } from 'react';
import { fetchWithAuth } from '../helpers/fetch';

const AuthContext = createContext();

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchWithAuth('/api/v1/auth/me')
      .then(res => res.ok ? res.json() : null)
      .then(data => setUser(data))
      .finally(() => setLoading(false));
  }, []);

  const register = async (email, password) => {
    try {
      const res = await fetchWithAuth('/api/v1/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });
      
      if (res.status !== 204) {
        const errorData = await res.json().catch(() => ({}));
        const errorMessage = errorData.message || 'Registration failed';
        setError(errorMessage);
        return;
      }

      const data = await fetchWithAuth('/api/v1/auth/me').then(r => r.json());
      setUser(data);
    } catch (err) {
      setError('Registration failed. Please try again.');
    }
  };

  const login = async (email, password) => {
    try {
      const res = await fetchWithAuth('/api/v1/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });
      
      if (res.status !== 204) {
        const errorData = await res.json().catch(() => ({}));
        const errorMessage = errorData.message || 'Login failed';
        setError(errorMessage);
        return;
      }

      const data = await fetchWithAuth('/api/v1/auth/me').then(r => r.json());
      setUser(data);
    } catch (err) {
      setError('Login failed. Please try again.');
    }
  };

  const logout = async () => {
    await fetchWithAuth('/api/v1/auth/logout', { method: 'POST' });
    setUser(null);
  };

  const clearError = () => setError(null);

  return (
    <AuthContext.Provider value={{ user, register, login, logout, loading, error, clearError }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext);

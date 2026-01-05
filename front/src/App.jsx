import { useAuth } from "./contexts/auth";
import Login from "./Login";

export default function App() {
  const { user, register, login, logout, loading } = useAuth();

  if (loading) {
    return <div>Loading...</div>
  }

  if (user === null) {
    return <Login register={register} login={login} />
  }

  return (
    <div>
      <h1>Welcome, {user.id}</h1>
      <button onClick={logout}>Logout</button>
    </div>
  )
}

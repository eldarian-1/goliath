import { useEffect, useState } from "react"

function User(user) {
  return <div>{user.name}</div>
}

export default function App() {
  const [users, setUsers] = useState(null)

  useEffect(() => {
    const fetchUsers = async () => {
      const res = await fetch('http://localhost:8080/api/v1/users');
      setUsers((await res.json()).users);
    };

    fetchUsers();
  }, []);

  return (
    <>
      { users === null && <div>Loading...</div> }
      { users !== null && users.map(user => <User key={user.id} {...user} />) }
    </>
  )
}

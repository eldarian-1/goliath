import { useState } from "react";

export default function Login({register, login}) {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const changeEmail = (event) => setEmail(event.target.value)
  const changePassword = (event) => setPassword(event.target.value)

  const handleRegister = () => register(email, password)
  const handleLogin = () => login(email, password)

  return (
    <div>
      <input type="text" placeholder="Email" onChange={changeEmail} value={email} />
      <input type="password" placeholder="Password" onChange={changePassword} value={password} />
      <button onClick={() => handleRegister(register)}>Register</button>
      <button onClick={() => handleLogin(login)}>Login</button>
    </div>
  )
}

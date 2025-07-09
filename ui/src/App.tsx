import { Routes, Route } from 'react-router-dom'
import Home from './pages/Home'
import SignupUser from './pages/SignupUser'
import SignupMerchant from './pages/SignupMerchant'
import Login from './pages/Login'

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/signup/user" element={<SignupUser />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup/merchant" element={<SignupMerchant />} />
    </Routes>
  )
}
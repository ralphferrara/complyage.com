import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'

type User = {
   user_id: number
   email: string
}

export default function Members() {
   const navigate = useNavigate()
   const [loading, setLoading] = useState(true)
   const [user, setUser] = useState<User | null>(null)

   useEffect(() => {
      const checkAuth = async () => {
         try {
            const res = await fetch('/auth/me', {
               method: 'GET',
               credentials: 'include', // important if using cookies
            })

            if (res.ok) {
               const data = await res.json()
               if (data.success) {
                  setUser(data.data)
               } else {
                  navigate('/login')
               }
            } else {
               navigate('/login')
            }
         } catch (err) {
            console.error(err)
            navigate('/login')
         } finally {
            setLoading(false)
         }
      }

      checkAuth()
   }, [navigate])

   if (loading) {
      return (
         <main className="min-h-screen flex items-center justify-center">
            <p>Checking your login status...</p>
         </main>
      )
   }

   if (!user) {
      return null // Redirect happens, so this shouldn't render
   }

   return (
      <main className="min-h-screen flex items-center justify-center">
         <div className="max-w-md w-full text-center">
            <h1 className="text-3xl font-bold mb-4">Welcome to Members Area</h1>
            <p className="text-base-content/70">
               Hello, <strong>{user.email}</strong>!
            </p>
            <p className="mt-4">Your user ID is: {user.user_id}</p>
         </div>
      </main>
   )
}

import React from 'react'
import { Link } from 'react-router-dom'

export default function SignupUser() {
   return (
      <main className="min-h-screen flex items-center justify-center bg-base-200 p-4">
         <div className="w-full max-w-md bg-base-100 shadow-lg rounded-lg p-8">
            <h1 className="text-3xl font-bold mb-4 text-center">User Signup</h1>
            <p className="text-sm text-center mb-8 text-base-content/70">
               Create your user account to access age-restricted content safely.
            </p>

            <form className="space-y-4">
               <div>
                  <label className="label">
                     <span className="label-text">Full Name</span>
                  </label>
                  <input
                     type="text"
                     placeholder="John Doe"
                     className="input input-bordered w-full"
                  />
               </div>

               <div>
                  <label className="label">
                     <span className="label-text">Email</span>
                  </label>
                  <input
                     type="email"
                     placeholder="you@example.com"
                     className="input input-bordered w-full"
                  />
               </div>

               <div>
                  <label className="label">
                     <span className="label-text">Password</span>
                  </label>
                  <input
                     type="password"
                     placeholder="********"
                     className="input input-bordered w-full"
                  />
               </div>

               <button type="submit" className="btn btn-primary w-full">
                  Sign Up
               </button>
            </form>

            <p className="mt-6 text-center text-sm">
               Already have an account?{' '}
               <Link to="/login" className="link link-primary">
                  Log In
               </Link>
            </p>
         </div>
      </main>
   )
}

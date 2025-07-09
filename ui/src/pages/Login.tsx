import React from 'react'
import { Link } from 'react-router-dom'

export default function Login() {
   return (
      <main className="min-h-screen flex items-center justify-center bg-base-200 p-4">
         <div className="w-full max-w-md bg-base-100 shadow-lg rounded-lg p-8">
            <h1 className="text-3xl font-bold mb-4 text-center">Log In</h1>
            <p className="text-sm text-center mb-8 text-base-content/70">
               Log in to your ComplyAge account.
            </p>

            <form className="space-y-4">
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

               <div className="form-control">
                  <label className="label cursor-pointer justify-start gap-2">
                     <input type="checkbox" className="checkbox" />
                     <span className="label-text">Remember me</span>
                  </label>
               </div>

               <button type="submit" className="btn btn-primary w-full">
                  Log In
               </button>
            </form>

            <div className="mt-6 text-center text-sm">
               <p>
                  Don&apos;t have an account?{' '}
                  <Link to="/signup/user" className="link link-primary">
                     Sign Up as User
                  </Link>{' '}
                  or{' '}
                  <Link to="/signup/merchant" className="link link-secondary">
                     Merchant
                  </Link>
               </p>
               <p className="mt-2">
                  <a href="#" className="link link-neutral">
                     Forgot your password?
                  </a>
               </p>
            </div>
         </div>
      </main>
   )
}

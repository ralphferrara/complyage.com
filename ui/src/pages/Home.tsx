import React from 'react'
import { Link } from 'react-router-dom'

export default function Home() {
   return (
      <main className="min-h-screen bg-base-100 text-base-content">
         {/* Navbar */}
         <div className="navbar bg-base-100 shadow-md">
            <div className="flex-1">
               <Link to="/" className="text-2xl font-bold text-primary">
                  ComplyAge
               </Link>
            </div>
            <div className="flex gap-2">
               <Link to="/login" className="btn btn-ghost">Login</Link>
               <Link to="/signup" className="btn btn-primary">Sign Up</Link>
            </div>
         </div>

         {/* Hero Section */}
         <section className="hero min-h-[60vh] bg-base-200 bg-cover" style={{
            backgroundImage: "url('https://picsum.photos/id/200/800/400')"
          }} >
            <div className="hero-content text-center">
               <div className="max-w-2xl">
                  <h1 className="text-5xl font-bold">
                     Effortless Age Verification for Any Platform
                  </h1>
                  <p className="py-6">
                     Protect your business and stay compliant with the simplest, fastest age verification API available.
                  </p>
                  <button className="btn btn-primary">Get Started</button>
               </div>
            </div>
         </section>

         {/* Carousel */}
         <section className="py-12">
            <div className="max-w-4xl mx-auto">
               <div className="carousel w-full rounded-box">
                  <div id="slide1" className="carousel-item relative w-full">
                     <img src="https://picsum.photos/id/237/800/400" className="w-full object-cover" alt="Security" />
                     <div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                        <a href="#slide3" className="btn btn-circle">❮</a>
                        <a href="#slide2" className="btn btn-circle">❯</a>
                     </div>
                  </div>
                  <div id="slide2" className="carousel-item relative w-full">
                     <img src="https://picsum.photos/id/238/800/400" className="w-full object-cover" alt="Verification" />
                     <div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                        <a href="#slide1" className="btn btn-circle">❮</a>
                        <a href="#slide3" className="btn btn-circle">❯</a>
                     </div>
                  </div>
                  <div id="slide3" className="carousel-item relative w-full">
                     <img src="https://picsum.photos/id/239/800/400" className="w-full object-cover" alt="Compliance" />
                     <div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                        <a href="#slide2" className="btn btn-circle">❮</a>
                        <a href="#slide1" className="btn btn-circle">❯</a>
                     </div>
                  </div>
               </div>
            </div>
         </section>

         {/* Features Section */}
         <section className="py-16 px-4 max-w-6xl mx-auto">
            <h2 className="text-4xl font-bold text-center mb-12">Why Choose ComplyAge?</h2>
            <div className="grid md:grid-cols-3 gap-8">
               <div className="card bg-base-100 shadow-lg">
                  <div className="card-body">
                     <h3 className="card-title">Easy Integration</h3>
                     <p>Plug & play API and SDKs for any stack. Get up and running in minutes.</p>
                  </div>
               </div>
               <div className="card bg-base-100 shadow-lg">
                  <div className="card-body">
                     <h3 className="card-title">Global Compliance</h3>
                     <p>Stay ahead of local and international age restriction laws automatically.</p>
                  </div>
               </div>
               <div className="card bg-base-100 shadow-lg">
                  <div className="card-body">
                     <h3 className="card-title">Rock-Solid Security</h3>
                     <p>All data encrypted, privacy-first. We never store sensitive user IDs longer than needed.</p>
                  </div>
               </div>
            </div>
         </section>

         {/* Call To Action */}
         <section className="py-12 bg-primary text-primary-content text-center">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
               Ready to protect your platform?
            </h2>
            <p className="mb-6">Start verifying ages in under 5 minutes.</p>
            <button className="btn btn-secondary">Sign Up Now</button>
         </section>
      </main>
   )
}

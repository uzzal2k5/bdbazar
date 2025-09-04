import React, { useState } from 'react';
import { Store, Users, DollarSign, Shield, TrendingUp, Star, CheckCircle, ArrowRight, Globe, Headphones, BarChart3 } from 'lucide-react';
import { AuthModal } from './AuthModal';
import { User } from '../types';
import { Header } from './Header'; // Adjust path if needed

interface SellerLoginPageProps {
  onLogin: (user: User) => void;
}

export const SellerLoginPage: React.FC<SellerLoginPageProps> = ({ onLogin }) => {
  const [showAuthModal, setShowAuthModal] = useState(false);

  const benefits = [
    {
      icon: DollarSign,
      title: 'Low Selling Fees',
      description: 'Only 3% commission on sales - one of the lowest in the industry',
      highlight: '3% Fee'
    },
    {
      icon: Users,
      title: 'Massive Customer Base',
      description: 'Reach over 10 million active buyers across Bangladesh and beyond',
      highlight: '10M+ Buyers'
    },
    {
      icon: Shield,
      title: 'Secure Payments',
      description: 'Get paid safely with our secure payment processing system',
      highlight: '100% Secure'
    },
    {
      icon: Globe,
      title: 'Global Reach',
      description: 'Sell locally and internationally with our shipping network',
      highlight: 'Worldwide'
    },
    {
      icon: BarChart3,
      title: 'Analytics & Insights',
      description: 'Track your performance with detailed sales analytics',
      highlight: 'Real-time Data'
    },
    {
      icon: Headphones,
      title: '24/7 Support',
      description: 'Get help whenever you need it with our dedicated seller support',
      highlight: 'Always Available'
    }
  ];

  const testimonials = [
    {
      name: 'Rashida Begum',
      business: 'Handmade Crafts',
      location: 'Dhaka',
      image: 'https://images.pexels.com/photos/1239291/pexels-photo-1239291.jpeg?auto=compress&cs=tinysrgb&w=150',
      quote: 'BDBazar transformed my small craft business. I went from selling locally to reaching customers nationwide. My monthly sales increased by 400%!',
      sales: '₹2.5L+ monthly sales',
      rating: 4.9,
      products: 150
    },
    {
      name: 'Mohammad Karim',
      business: 'Electronics Store',
      location: 'Chittagong',
      image: 'https://images.pexels.com/photos/1222271/pexels-photo-1222271.jpeg?auto=compress&cs=tinysrgb&w=150',
      quote: 'The platform is incredibly user-friendly. Setting up my store took just minutes, and the marketing tools helped me reach the right customers.',
      sales: '₹5L+ monthly sales',
      rating: 4.8,
      products: 300
    },
    {
      name: 'Fatima Khan',
      business: 'Fashion Boutique',
      location: 'Sylhet',
      image: 'https://images.pexels.com/photos/1181686/pexels-photo-1181686.jpeg?auto=compress&cs=tinysrgb&w=150',
      quote: 'BDBazar gave me the tools to grow my fashion business beyond my wildest dreams. The customer support is exceptional!',
      sales: '₹3.2L+ monthly sales',
      rating: 5.0,
      products: 200
    },
    {
      name: 'Abdul Rahman',
      business: 'Organic Foods',
      location: 'Rajshahi',
      image: 'https://images.pexels.com/photos/1043471/pexels-photo-1043471.jpeg?auto=compress&cs=tinysrgb&w=150',
      quote: 'Started with just 10 products, now I have over 500! BDBazar\'s logistics support made nationwide delivery possible for my organic food business.',
      sales: '₹1.8L+ monthly sales',
      rating: 4.7,
      products: 500
    }
  ];

  const steps = [
    {
      number: 1,
      title: 'Create Your Account',
      description: 'Sign up with your email and business details. It takes less than 2 minutes!',
      icon: Users,
      details: ['Provide basic business information', 'Verify your email address', 'Choose your seller plan']
    },
    {
      number: 2,
      title: 'Set Up Your Store',
      description: 'Customize your storefront with your brand, logo, and business information.',
      icon: Store,
      details: ['Upload your business logo', 'Write your store description', 'Set up payment methods']
    },
    {
      number: 3,
      title: 'List Your Products',
      description: 'Add your products with high-quality photos and detailed descriptions.',
      icon: BarChart3,
      details: ['Upload product photos', 'Write compelling descriptions', 'Set competitive prices']
    },
    {
      number: 4,
      title: 'Start Selling',
      description: 'Receive orders, manage inventory, and fulfill customer requests.',
      icon: TrendingUp,
      details: ['Process incoming orders', 'Manage your inventory', 'Communicate with customers']
    },
    {
      number: 5,
      title: 'Grow Your Business',
      description: 'Use our analytics and marketing tools to scale your business.',
      icon: DollarSign,
      details: ['Analyze sales performance', 'Run promotional campaigns', 'Expand your product range']
    }
  ];

  const stats = [
    { label: 'Active Buyers', value: '10M+', icon: Users },
    { label: 'Sellers', value: '50K+', icon: Store },
    { label: 'Total Sales', value: '₹500Cr+', icon: DollarSign },
    { label: 'Products Listed', value: '2M+', icon: BarChart3 }
  ];

  const handleGetStarted = () => {
    setShowAuthModal(true);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-green-50">


      {/* Hero Section */}
      <div className="relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-r from-blue-600/10 to-green-600/10"></div>
        <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div>
              <div className="flex items-center space-x-2 mb-6">
                <Store className="h-8 w-8 text-blue-600" />
                <span className="text-2xl font-bold text-gray-900">BD Bazar</span>
                <span className="bg-green-100 text-green-800 px-3 py-1 rounded-full text-sm font-medium">
                  For Sellers
                </span>
              </div>

              <h1 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
                Start Selling on
                <span className="text-blue-600"> Bangladesh's</span>
                <br />
                <span className="text-green-600">Leading Marketplace</span>
              </h1>

              <p className="text-xl text-gray-600 mb-8">
                Join thousands of successful sellers and grow your business with our powerful e-commerce platform.
                Low fees, high reach, maximum profits.
              </p>

              <div className="flex flex-col sm:flex-row gap-4 mb-8">
                <button
                  onClick={handleGetStarted}
                  className="bg-blue-600 text-white px-8 py-4 rounded-lg hover:bg-blue-700 transition-colors font-semibold text-lg flex items-center justify-center space-x-2"
                >
                  <span>Start Selling Today</span>
                  <ArrowRight className="h-5 w-5" />
                </button>
                <button className="border border-gray-300 text-gray-700 px-8 py-4 rounded-lg hover:bg-gray-50 transition-colors font-semibold">
                  Watch Demo
                </button>

              </div>

              <div className="bg-green-50 border border-green-200 rounded-lg p-4">
                <div className="flex items-center space-x-2 mb-2">
                  <CheckCircle className="h-5 w-5 text-green-600" />
                  <span className="font-semibold text-green-800">Special Launch Offer</span>
                </div>
                <ul className="text-sm text-green-700 space-y-1">
                  <li>• Free store setup (worth ₹5,000)</li>
                  <li>• First month commission-free</li>
                  <li>• Dedicated onboarding support</li>
                </ul>
              </div>
            </div>

            <div className="relative">
              <div className="bg-white rounded-2xl shadow-2xl p-8 border border-gray-100">
                <div className="text-center mb-6">
                  <h3 className="text-2xl font-bold text-gray-900 mb-2">Join BD Bazar Today</h3>
                  <p className="text-gray-600">Already have an account? <button className="text-blue-600 hover:text-blue-700 font-medium">Sign in</button></p>
                </div>

                <form className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">Business Name</label>
                    <input
                      type="text"
                      placeholder="Enter your business name"
                      className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">Email Address</label>
                    <input
                      type="email"
                      placeholder="Enter your email"
                      className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">Phone Number</label>
                    <input
                      type="tel"
                      placeholder="Enter your phone number"
                      className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">Business Category</label>
                    <select className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                      <option>Select your business category</option>
                      <option>Fashion & Clothing</option>
                      <option>Electronics</option>
                      <option>Home & Garden</option>
                      <option>Health & Beauty</option>
                      <option>Sports & Outdoors</option>
                      <option>Books & Media</option>
                      <option>Food & Beverages</option>
                      <option>Other</option>
                    </select>
                  </div>
                  <button
                    type="button"
                    onClick={handleGetStarted}
                    className="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition-colors font-semibold"
                  >
                    Create Seller Account
                  </button>
                </form>

                <p className="text-xs text-gray-500 text-center mt-4">
                  By signing up, you agree to our Terms of Service and Privacy Policy
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Stats Section */}
      <div className="bg-white py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900 mb-4">Trusted by Thousands</h2>
            <p className="text-gray-600">Join a thriving community of successful sellers</p>
          </div>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            {stats.map((stat, index) => {
              const Icon = stat.icon;
              return (
                <div key={index} className="text-center">
                  <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
                    <Icon className="h-8 w-8 text-blue-600" />
                  </div>
                  <div className="text-3xl font-bold text-gray-900 mb-2">{stat.value}</div>
                  <div className="text-gray-600">{stat.label}</div>
                </div>
              );
            })}
          </div>
        </div>
      </div>

      {/* Benefits Section */}
      <div className="bg-gray-50 py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900 mb-4">Why Sell on BD Bazar?</h2>
            <p className="text-gray-600 max-w-2xl mx-auto">
              We provide everything you need to build and grow a successful online business
            </p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {benefits.map((benefit, index) => {
              const Icon = benefit.icon;
              return (
                <div key={index} className="bg-white rounded-xl p-6 shadow-md hover:shadow-lg transition-shadow">
                  <div className="flex items-center justify-between mb-4">
                    <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
                      <Icon className="h-6 w-6 text-blue-600" />
                    </div>
                    <span className="bg-green-100 text-green-800 px-3 py-1 rounded-full text-sm font-medium">
                      {benefit.highlight}
                    </span>
                  </div>
                  <h3 className="text-xl font-semibold text-gray-900 mb-2">{benefit.title}</h3>
                  <p className="text-gray-600">{benefit.description}</p>
                </div>
              );
            })}
          </div>
        </div>
      </div>

      {/* Testimonials Section */}
      <div className="bg-white py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900 mb-4">Success Stories</h2>
            <p className="text-gray-600">Hear from sellers who transformed their businesses with BD Bazar</p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            {testimonials.map((testimonial, index) => (
              <div key={index} className="bg-gray-50 rounded-xl p-6 border border-gray-200">
                <div className="flex items-start space-x-4 mb-4">
                  <img
                    src={testimonial.image}
                    alt={testimonial.name}
                    className="w-16 h-16 rounded-full object-cover"
                  />
                  <div className="flex-1">
                    <h4 className="font-semibold text-gray-900">{testimonial.name}</h4>
                    <p className="text-blue-600 font-medium">{testimonial.business}</p>
                    <p className="text-gray-500 text-sm">{testimonial.location}</p>
                    <div className="flex items-center space-x-1 mt-1">
                      {[...Array(5)].map((_, i) => (
                        <Star
                          key={i}
                          className={`h-4 w-4 ${
                            i < Math.floor(testimonial.rating) ? 'text-yellow-400 fill-current' : 'text-gray-300'
                          }`}
                        />
                      ))}
                      <span className="text-sm text-gray-600 ml-1">{testimonial.rating}</span>
                    </div>
                  </div>
                </div>
                <blockquote className="text-gray-700 italic mb-4">
                  "{testimonial.quote}"
                </blockquote>
                <div className="flex justify-between text-sm">
                  <span className="text-green-600 font-medium">{testimonial.sales}</span>
                  <span className="text-gray-500">{testimonial.products} products</span>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* How to Start Section */}
      <div className="bg-gradient-to-r from-blue-600 to-green-600 py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-white mb-4">5 Simple Steps to Start Selling</h2>
            <p className="text-blue-100 max-w-2xl mx-auto">
              Get your business online in minutes with our easy setup process
            </p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-5 gap-8">
            {steps.map((step, index) => {
              const Icon = step.icon;
              return (
                <div key={index} className="text-center">
                  <div className="relative mb-6">
                    <div className="w-16 h-16 bg-white rounded-full flex items-center justify-center mx-auto shadow-lg">
                      <Icon className="h-8 w-8 text-blue-600" />
                    </div>
                    <div className="absolute -top-2 -right-2 w-8 h-8 bg-green-500 rounded-full flex items-center justify-center text-white font-bold text-sm">
                      {step.number}
                    </div>
                    {index < steps.length - 1 && (
                      <div className="hidden md:block absolute top-8 left-full w-full h-0.5 bg-white/30 -translate-x-1/2"></div>
                    )}
                  </div>
                  <h3 className="text-xl font-semibold text-white mb-2">{step.title}</h3>
                  <p className="text-blue-100 text-sm mb-4">{step.description}</p>
                  <ul className="text-xs text-blue-200 space-y-1">
                    {step.details.map((detail, detailIndex) => (
                      <li key={detailIndex} className="flex items-center justify-center space-x-1">
                        <CheckCircle className="h-3 w-3 flex-shrink-0" />
                        <span>{detail}</span>
                      </li>
                    ))}
                  </ul>
                </div>
              );
            })}
          </div>
        </div>
      </div>

      {/* CTA Section */}
      <div className="bg-white py-16">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl font-bold text-gray-900 mb-4">Ready to Start Your Success Story?</h2>
          <p className="text-xl text-gray-600 mb-8">
            Join thousands of sellers who are already growing their businesses on BD Bazar
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <button
              onClick={handleGetStarted}
              className="bg-blue-600 text-white px-8 py-4 rounded-lg hover:bg-blue-700 transition-colors font-semibold text-lg flex items-center justify-center space-x-2"
            >
              <span>Start Selling Now</span>
              <ArrowRight className="h-5 w-5" />
            </button>
            <button className="border border-gray-300 text-gray-700 px-8 py-4 rounded-lg hover:bg-gray-50 transition-colors font-semibold">
              Contact Sales Team
            </button>
          </div>
          <p className="text-sm text-gray-500 mt-4">
            No setup fees • No monthly charges • Only pay when you sell
          </p>
        </div>
      </div>

      {/* Auth Modal */}
      {showAuthModal && (
        <AuthModal
          onClose={() => setShowAuthModal(false)}
          onLogin={onLogin}
        />
      )}
    </div>
  );
};
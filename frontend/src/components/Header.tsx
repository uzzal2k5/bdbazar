import React from 'react';
import { ShoppingCart, User, Search, Store, Home, Package } from 'lucide-react';
import { User as UserType } from '../types';

interface HeaderProps {
  cartItemCount: number;
  onCartClick: () => void;
  onAuthClick: () => void;
  onOrderHistoryClick: () => void;
  currentUser: UserType | null;
  onPageChange: (page: 'home' | 'seller') => void;
  currentPage: 'home' | 'seller';
  onSearch: (query: string) => void;
  searchQuery: string;
}

export const Header: React.FC<HeaderProps> = ({
  cartItemCount,
  onCartClick,
  onAuthClick,
  onOrderHistoryClick,
  currentUser,
  onPageChange,
  currentPage,
  onSearch,
  searchQuery,
}) => {
  return (
    <header className="fixed top-0 left-0 right-0 bg-white shadow-lg z-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center space-x-8">
            <div className="flex items-center space-x-2">
              <Store className="h-8 w-8 text-blue-600" />
              <h1 className="text-xl font-bold text-gray-900">BD Bazar</h1>
            </div>

            <nav className="hidden md:flex space-x-6">
              <button
                onClick={() => onPageChange('home')}
                className={`flex items-center space-x-2 px-3 py-2 rounded-md transition-colors ${
                  currentPage === 'home'
                    ? 'bg-blue-100 text-blue-600'
                    : 'text-gray-700 hover:bg-gray-100'
                }`}
              >
                <Home className="h-4 w-4" />
                <span>Shop</span>
              </button>

              {currentUser?.isSeller && (
                <button
                  onClick={() => onPageChange('seller')}
                  className={`flex items-center space-x-2 px-3 py-2 rounded-md transition-colors ${
                    currentPage === 'seller'
                      ? 'bg-blue-100 text-blue-600'
                      : 'text-gray-700 hover:bg-gray-100'
                  }`}
                >
                  <Store className="h-4 w-4" />
                  <span>Sell</span>
                </button>
              )}
            </nav>
          </div>

          <div className="flex-1 max-w-lg mx-8">
            <div className="relative">
              <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <input
                type="text"
                placeholder="Search products..."
                value={searchQuery}
                onChange={(e) => onSearch(e.target.value)}
                className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
          </div>

          <div className="flex items-center space-x-4">
            {currentUser && (
              <button
                onClick={onOrderHistoryClick}
                className="relative p-2 text-gray-700 hover:text-blue-600 transition-colors"
                title="Order History"
              >
                <Package className="h-6 w-6" />
              </button>
            )}

            <button
              onClick={onCartClick}
              className="relative p-2 text-gray-700 hover:text-blue-600 transition-colors"
            >
              <ShoppingCart className="h-6 w-6" />
              {cartItemCount > 0 && (
                <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">
                  {cartItemCount}
                </span>
              )}
            </button>

            <button
              onClick={onAuthClick}
              className="flex items-center space-x-2 p-2 text-gray-700 hover:text-blue-600 transition-colors"
            >
              <User className="h-6 w-6" />
              {currentUser && <span className="hidden sm:inline">{currentUser.name}</span>}
            </button>
          </div>
        </div>
      </div>
    </header>
  );
};
import React from 'react';
import { Store, Mail, Phone, MapPin } from 'lucide-react';
interface FooterProps {
  onNavigate: (page: 'home' | 'seller' | 'seller-login') => void;
}

// export const Footer: React.FC = () => {
export const Footer: React.FC<FooterProps> = ({ onNavigate }) => {
  return (
    <footer className="bg-gray-900 text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div className="space-y-4">
            <div className="flex items-center space-x-2">
              <Store className="h-8 w-8 text-blue-400" />
              <span className="text-xl font-bold">BD Bazar</span>
            </div>
            <p className="text-gray-400">
              The Smart Way to Buy and Sell in Bangladesh — With a Partner You Trust.
            </p>
          </div>

          <div>
            <h3 className="font-semibold mb-4">Quick Links</h3>
            <ul className="space-y-2 text-gray-400">
              <li><a href="#" className="hover:text-white transition-colors">About Us</a></li>
              <li><a href="#" className="hover:text-white transition-colors">How It Works</a></li>
              <li>
              <button
                 onClick={() => onNavigate('seller-login')}
                 className="hover:text-white transition-colors text-left"
              >
               Become Seller
              </button></li>
              <li><a href="#" className="hover:text-white transition-colors">Seller Guide</a></li>
              <li><a href="#" className="hover:text-white transition-colors">FAQ</a></li>
            </ul>
          </div>

          <div>
            <h3 className="font-semibold mb-4">Support</h3>
            <ul className="space-y-2 text-gray-400">
              <li><a href="#" className="hover:text-white transition-colors">Help Center</a></li>
              <li><a href="#" className="hover:text-white transition-colors">Contact Us</a></li>
              <li><a href="#" className="hover:text-white transition-colors">Shipping Info</a></li>
              <li><a href="#" className="hover:text-white transition-colors">Returns</a></li>
            </ul>
          </div>

          <div>
            <h3 className="font-semibold mb-4">Contact Info</h3>
            <div className="space-y-2 text-gray-400">
              <div className="flex items-center space-x-2">
                <Mail className="h-4 w-4" />
                <span>uzzal2k5@gmail.com</span>
              </div>
              <div className="flex items-center space-x-2">
                <Phone className="h-4 w-4" />
                <span>+880-1715-519132</span>
              </div>
              <div className="flex items-center space-x-2">
                <MapPin className="h-4 w-4" />
                <span>SSK Road, Feni-3900, Chattogram</span>
              </div>
            </div>
          </div>
        </div>

        <div className="border-t border-gray-800 mt-8 pt-8 text-center text-gray-400">
          <p>&copy; 2025 BDBazar. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
};
import React, { useState, useEffect } from 'react';
import { Clock, Star, ShoppingCart, Zap } from 'lucide-react';
import { Product } from '../types';

interface FlashSaleProps {
  products: Product[];
  onProductClick: (product: Product) => void;
  onAddToCart: (product: Product) => void;
}

export const FlashSale: React.FC<FlashSaleProps> = ({
  products,
  onProductClick,
  onAddToCart,
}) => {
  const [timeLeft, setTimeLeft] = useState({
    hours: 23,
    minutes: 45,
    seconds: 30,
  });

  // Flash sale products (mock data with discounts)
  const flashSaleProducts = products.slice(0, 6).map(product => ({
    ...product,
    originalPrice: product.price * 1.4, // 30% discount
    discount: 30,
    flashSaleStock: Math.floor(product.stock * 0.3), // Limited stock for flash sale
  }));

  useEffect(() => {
    const timer = setInterval(() => {
      setTimeLeft(prev => {
        let { hours, minutes, seconds } = prev;

        if (seconds > 0) {
          seconds--;
        } else if (minutes > 0) {
          minutes--;
          seconds = 59;
        } else if (hours > 0) {
          hours--;
          minutes = 59;
          seconds = 59;
        } else {
          // Reset timer for next flash sale
          hours = 23;
          minutes = 59;
          seconds = 59;
        }

        return { hours, minutes, seconds };
      });
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  const formatTime = (time: number) => time.toString().padStart(2, '0');

  return (
    <div className="bg-gradient-to-r from-red-500 to-pink-600 rounded-xl p-6 mb-8">
      <div className="flex items-center justify-between mb-6">
        <div className="flex items-center space-x-3">
          <div className="w-12 h-12 bg-white rounded-full flex items-center justify-center">
            <Zap className="h-6 w-6 text-red-500" />
          </div>
          <div>
            <h2 className="text-2xl font-bold text-white">Flash Sale</h2>
            <p className="text-red-100">Limited time offers - Up to 50% off!</p>
          </div>
        </div>

        {/* Countdown Timer */}
        <div className="bg-white bg-opacity-20 rounded-lg p-4">
          <div className="flex items-center space-x-2 text-white">
            <Clock className="h-5 w-5" />
            <span className="text-sm font-medium">Ends in:</span>
          </div>
          <div className="flex items-center space-x-2 mt-2">
            <div className="bg-white text-red-600 px-2 py-1 rounded font-bold text-lg">
              {formatTime(timeLeft.hours)}
            </div>
            <span className="text-white font-bold">:</span>
            <div className="bg-white text-red-600 px-2 py-1 rounded font-bold text-lg">
              {formatTime(timeLeft.minutes)}
            </div>
            <span className="text-white font-bold">:</span>
            <div className="bg-white text-red-600 px-2 py-1 rounded font-bold text-lg">
              {formatTime(timeLeft.seconds)}
            </div>
          </div>
          <div className="flex justify-between text-xs text-red-100 mt-1">
            <span>Hours</span>
            <span>Minutes</span>
            <span>Seconds</span>
          </div>
        </div>
      </div>

      {/* Flash Sale Products */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-6 gap-4">
        {flashSaleProducts.map((product) => (
          <div
            key={product.id}
            className="bg-white rounded-lg overflow-hidden shadow-lg hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1"
          >
            <div className="relative">
              <img
                src={product.image}
                alt={product.name}
                className="w-full h-32 object-cover cursor-pointer"
                onClick={() => onProductClick(product)}
              />
              <div className="absolute top-2 left-2 bg-red-500 text-white px-2 py-1 rounded-full text-xs font-bold">
                -{product.discount}%
              </div>
              <div className="absolute top-2 right-2 bg-white rounded-full px-2 py-1 flex items-center space-x-1">
                <Star className="h-3 w-3 text-yellow-400 fill-current" />
                <span className="text-xs font-medium">{product.rating}</span>
              </div>
            </div>

            <div className="p-3">
              <h3
                className="font-semibold text-sm mb-2 cursor-pointer hover:text-blue-600 transition-colors line-clamp-2"
                onClick={() => onProductClick(product)}
              >
                {product.name}
              </h3>

              <div className="space-y-2">
                <div className="flex items-center space-x-2">
                  <span className="text-lg font-bold text-red-600">${product.price}</span>
                  <span className="text-sm text-gray-500 line-through">${product.originalPrice.toFixed(2)}</span>
                </div>

                <div className="flex items-center justify-between text-xs text-gray-600">
                  <span>{product.flashSaleStock} left</span>
                  <span>by {product.sellerName}</span>
                </div>

                {/* Stock Progress Bar */}
                <div className="w-full bg-gray-200 rounded-full h-2">
                  <div
                    className="bg-red-500 h-2 rounded-full transition-all duration-300"
                    style={{
                      width: `${Math.max(10, (product.flashSaleStock / (product.stock * 0.3)) * 100)}%`
                    }}
                  ></div>
                </div>

                <button
                  onClick={() => onAddToCart(product)}
                  className="w-full bg-red-500 text-white py-2 rounded-lg hover:bg-red-600 transition-colors flex items-center justify-center space-x-2 text-sm font-medium"
                  disabled={product.flashSaleStock === 0}
                >
                  <ShoppingCart className="h-4 w-4" />
                  <span>{product.flashSaleStock === 0 ? 'Sold Out' : 'Add to Cart'}</span>
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="text-center mt-6">
        <button className="bg-white text-red-600 px-8 py-3 rounded-lg font-semibold hover:bg-red-50 transition-colors">
          View All Flash Sale Items
        </button>
      </div>
    </div>
  );
};
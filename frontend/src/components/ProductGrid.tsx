import React from 'react';
import { Star, ShoppingCart } from 'lucide-react';
import { Product } from '../types';

interface ProductGridProps {
  products: Product[];
  onProductClick: (product: Product) => void;
  onAddToCart: (product: Product) => void;
}

export const ProductGrid: React.FC<ProductGridProps> = ({
  products,
  onProductClick,
  onAddToCart,
}) => {
  if (products.length === 0) {
    return (
      <div className="text-center py-16">
        <p className="text-gray-500 text-lg">No products found</p>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      {products.map((product) => (
        <div
          key={product.id}
          className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1"
        >
          <div className="relative">
            <img
              src={product.image}
              alt={product.name}
              className="w-full h-48 object-cover cursor-pointer"
              onClick={() => onProductClick(product)}
            />
            <div className="absolute top-2 right-2 bg-white rounded-full px-2 py-1 flex items-center space-x-1">
              <Star className="h-4 w-4 text-yellow-400 fill-current" />
              <span className="text-sm font-medium">{product.rating}</span>
            </div>
          </div>

          <div className="p-4">
            <h3
              className="font-semibold text-lg mb-2 cursor-pointer hover:text-blue-600 transition-colors"
              onClick={() => onProductClick(product)}
            >
              {product.name}
            </h3>
            <p className="text-gray-600 text-sm mb-3 line-clamp-2">{product.description}</p>

            <div className="flex items-center justify-between mb-3">
              <span className="text-2xl font-bold text-blue-600">${product.price}</span>
              <span className="text-sm text-gray-500">by {product.sellerName}</span>
            </div>

            <div className="flex items-center justify-between">
              <span className="text-sm text-gray-500">{product.stock} in stock</span>
              <button
                onClick={() => onAddToCart(product)}
                className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors flex items-center space-x-2"
              >
                <ShoppingCart className="h-4 w-4" />
                <span>Add</span>
              </button>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};
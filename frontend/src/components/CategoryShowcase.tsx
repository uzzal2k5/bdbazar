import React from 'react';
import { ArrowRight, Smartphone, Shirt, Home, Book, Dumbbell, Gamepad2, Heart, Car } from 'lucide-react';
import { Product } from '../types';

interface CategoryShowcaseProps {
  products: Product[];
  onCategoryClick: (category: string) => void;
  onProductClick: (product: Product) => void;
}

export const CategoryShowcase: React.FC<CategoryShowcaseProps> = ({
  products,
  onCategoryClick,
  onProductClick,
}) => {
  const categoryIcons = {
    electronics: Smartphone,
    clothing: Shirt,
    home: Home,
    books: Book,
    sports: Dumbbell,
    toys: Gamepad2,
    beauty: Heart,
    automotive: Car,
  };

  const categories = Array.from(new Set(products.map(p => p.category))).map(category => {
    const categoryProducts = products.filter(p => p.category === category);
    const Icon = categoryIcons[category as keyof typeof categoryIcons] || Smartphone;

    return {
      name: category,
      icon: Icon,
      productCount: categoryProducts.length,
      featuredProducts: categoryProducts.slice(0, 4),
      averagePrice: categoryProducts.reduce((sum, p) => sum + p.price, 0) / categoryProducts.length,
    };
  });

  return (
    <div className="mb-8">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-2xl font-bold text-gray-900">Shop by Category</h2>
        <button className="text-blue-600 hover:text-blue-700 font-medium flex items-center space-x-1">
          <span>View All Categories</span>
          <ArrowRight className="h-4 w-4" />
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {categories.map((category) => {
          const Icon = category.icon;
          return (
            <div
              key={category.name}
              className="bg-white rounded-xl shadow-md hover:shadow-xl transition-all duration-300 overflow-hidden cursor-pointer transform hover:-translate-y-1"
              onClick={() => onCategoryClick(category.name)}
            >
              {/* Category Header */}
              <div className="bg-gradient-to-r from-blue-500 to-purple-600 p-4 text-white">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-white bg-opacity-20 rounded-full flex items-center justify-center">
                    <Icon className="h-5 w-5" />
                  </div>
                  <div>
                    <h3 className="font-semibold text-lg capitalize">{category.name}</h3>
                    <p className="text-blue-100 text-sm">{category.productCount} products</p>
                  </div>
                </div>
                <div className="mt-2 text-right">
                  <p className="text-blue-100 text-sm">From ${category.averagePrice.toFixed(0)}</p>
                </div>
              </div>

              {/* Featured Products Grid */}
              <div className="p-4">
                <div className="grid grid-cols-2 gap-2">
                  {category.featuredProducts.map((product, index) => (
                    <div
                      key={product.id}
                      className="relative group"
                      onClick={(e) => {
                        e.stopPropagation();
                        onProductClick(product);
                      }}
                    >
                      <img
                        src={product.image}
                        alt={product.name}
                        className="w-full h-20 object-cover rounded-lg group-hover:opacity-80 transition-opacity"
                      />
                      <div className="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-20 rounded-lg transition-all duration-200 flex items-center justify-center">
                        <span className="text-white text-xs font-medium opacity-0 group-hover:opacity-100 transition-opacity">
                          ${product.price}
                        </span>
                      </div>
                    </div>
                  ))}
                </div>

                <button
                  onClick={(e) => {
                    e.stopPropagation();
                    onCategoryClick(category.name);
                  }}
                  className="w-full mt-3 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg transition-colors flex items-center justify-center space-x-2"
                >
                  <span>Explore {category.name}</span>
                  <ArrowRight className="h-4 w-4" />
                </button>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};
import React, { useState, useEffect } from 'react';
import { Star, ShoppingCart, Heart, Eye, RefreshCw } from 'lucide-react';
import { Product } from '../types';

interface ProductsForYouProps {
  products: Product[];
  onProductClick: (product: Product) => void;
  onAddToCart: (product: Product) => void;
  currentUser?: any;
}

export const ProductsForYou: React.FC<ProductsForYouProps> = ({
  products,
  onProductClick,
  onAddToCart,
  currentUser,
}) => {
  const [recommendedProducts, setRecommendedProducts] = useState<Product[]>([]);
  const [isRefreshing, setIsRefreshing] = useState(false);

  // Mock recommendation algorithm
  const generateRecommendations = () => {
    let recommended: Product[] = [];

    if (currentUser) {
      // Personalized recommendations based on user behavior (mock)
      const userPreferences = ['electronics', 'clothing']; // Mock user preferences
      const preferredProducts = products.filter(p =>
        userPreferences.includes(p.category)
      );

      // Add highly rated products from preferred categories
      const highRatedPreferred = preferredProducts
        .filter(p => p.rating >= 4.0)
        .sort((a, b) => b.rating - a.rating)
        .slice(0, 4);

      recommended = [...highRatedPreferred];
    }

    // Add trending products (high rating + good sales)
    const trending = products
      .filter(p => p.rating >= 4.2 && p.stock < 30) // Low stock indicates good sales
      .sort((a, b) => b.rating - a.rating)
      .slice(0, 4);

    // Add new arrivals (mock - using higher IDs as newer)
    const newArrivals = products
      .sort((a, b) => parseInt(b.id) - parseInt(a.id))
      .slice(0, 4);

    // Combine and remove duplicates
    const combined = [...recommended, ...trending, ...newArrivals];
    const unique = combined.filter((product, index, self) =>
      index === self.findIndex(p => p.id === product.id)
    );

    return unique.slice(0, 8);
  };

  useEffect(() => {
    setRecommendedProducts(generateRecommendations());
  }, [products, currentUser]);

  const refreshRecommendations = () => {
    setIsRefreshing(true);
    setTimeout(() => {
      setRecommendedProducts(generateRecommendations());
      setIsRefreshing(false);
    }, 1000);
  };

  const getRecommendationReason = (product: Product): string => {
    if (product.rating >= 4.5) return 'Highly Rated';
    if (product.stock < 20) return 'Trending';
    if (parseInt(product.id) > 5) return 'New Arrival';
    return 'Recommended';
  };

  return (
    <div className="mb-8">
      <div className="flex items-center justify-between mb-6">
        <div>
          <h2 className="text-2xl font-bold text-gray-900">
            {currentUser ? `Products for You, ${currentUser.name}` : 'Recommended Products'}
          </h2>
          <p className="text-gray-600 mt-1">
            {currentUser
              ? 'Curated based on your preferences and browsing history'
              : 'Popular and trending items you might like'
            }
          </p>
        </div>
        <button
          onClick={refreshRecommendations}
          disabled={isRefreshing}
          className="flex items-center space-x-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50"
        >
          <RefreshCw className={`h-4 w-4 ${isRefreshing ? 'animate-spin' : ''}`} />
          <span>Refresh</span>
        </button>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        {recommendedProducts.map((product) => (
          <div
            key={product.id}
            className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 group"
          >
            <div className="relative">
              <img
                src={product.image}
                alt={product.name}
                className="w-full h-48 object-cover cursor-pointer group-hover:scale-105 transition-transform duration-300"
                onClick={() => onProductClick(product)}
              />

              {/* Recommendation Badge */}
              <div className="absolute top-2 left-2 bg-blue-600 text-white px-2 py-1 rounded-full text-xs font-medium">
                {getRecommendationReason(product)}
              </div>

              {/* Rating Badge */}
              <div className="absolute top-2 right-2 bg-white rounded-full px-2 py-1 flex items-center space-x-1 shadow-md">
                <Star className="h-3 w-3 text-yellow-400 fill-current" />
                <span className="text-xs font-medium">{product.rating}</span>
              </div>

              {/* Quick Actions */}
              <div className="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-20 transition-all duration-300 flex items-center justify-center opacity-0 group-hover:opacity-100">
                <div className="flex space-x-2">
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                      onProductClick(product);
                    }}
                    className="w-10 h-10 bg-white rounded-full flex items-center justify-center hover:bg-gray-100 transition-colors"
                    title="Quick View"
                  >
                    <Eye className="h-4 w-4 text-gray-700" />
                  </button>
                  <button
                    className="w-10 h-10 bg-white rounded-full flex items-center justify-center hover:bg-gray-100 transition-colors"
                    title="Add to Wishlist"
                  >
                    <Heart className="h-4 w-4 text-gray-700" />
                  </button>
                </div>
              </div>
            </div>

            <div className="p-4">
              <h3
                className="font-semibold text-lg mb-2 cursor-pointer hover:text-blue-600 transition-colors line-clamp-2"
                onClick={() => onProductClick(product)}
              >
                {product.name}
              </h3>

              <p className="text-gray-600 text-sm mb-3 line-clamp-2">
                {product.description}
              </p>

              <div className="flex items-center justify-between mb-3">
                <span className="text-2xl font-bold text-blue-600">${product.price}</span>
                <span className="text-sm text-gray-500">by {product.sellerName}</span>
              </div>

              <div className="flex items-center justify-between mb-4">
                <div className="flex items-center space-x-1">
                  <div className="flex items-center">
                    {[...Array(5)].map((_, i) => (
                      <Star
                        key={i}
                        className={`h-4 w-4 ${
                          i < Math.floor(product.rating)
                            ? 'text-yellow-400 fill-current'
                            : 'text-gray-300'
                        }`}
                      />
                    ))}
                  </div>
                  <span className="text-sm text-gray-600">({product.reviews.length})</span>
                </div>
                <span className="text-sm text-green-600 font-medium">
                  {product.stock} in stock
                </span>
              </div>

              {/* Tags */}
              <div className="flex flex-wrap gap-1 mb-4">
                {product.tags.slice(0, 3).map((tag, index) => (
                  <span
                    key={index}
                    className="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded-full"
                  >
                    {tag}
                  </span>
                ))}
              </div>

              <button
                onClick={() => onAddToCart(product)}
                className="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition-colors flex items-center justify-center space-x-2 font-medium"
              >
                <ShoppingCart className="h-4 w-4" />
                <span>Add to Cart</span>
              </button>
            </div>
          </div>
        ))}
      </div>

      {/* Load More Button */}
      <div className="text-center mt-8">
        <button className="px-8 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors font-medium">
          Load More Recommendations
        </button>
      </div>
    </div>
  );
};
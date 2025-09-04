import React, { useState, useEffect } from 'react';
import { Header } from './components/Header';
import { ProductGrid } from './components/ProductGrid';
import { ProductDetail } from './components/ProductDetail';
import { Cart } from './components/Cart';
import { Checkout } from './components/Checkout';
import { OrderConfirmation } from './components/OrderConfirmation';
import { OrderHistory } from './components/OrderHistory';
import { SellerDashboard } from './components/SellerDashboard';
import { SellerLoginPage } from './components/SellerLoginPage';
import { AuthModal } from './components/AuthModal';
import { Footer } from './components/Footer';
import { FlashSale } from './components/FlashSale';
import { CategoryShowcase } from './components/CategoryShowcase';
import { ProductsForYou } from './components/ProductsForYou';
import { Product, CartItem, User, SellerProfile, CheckoutData, Order } from './types';
import { generateMockProducts } from './utils/mockData';
import { generateTrackingNumber, generateOrderNumber, calculateEstimatedDelivery, createMockTrackingEvents, createMockOrders } from './utils/orderUtils';

function App() {
  const [products, setProducts] = useState<Product[]>([]);
  const [filteredProducts, setFilteredProducts] = useState<Product[]>([]);
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);
  const [cartItems, setCartItems] = useState<CartItem[]>([]);
  const [isCartOpen, setIsCartOpen] = useState(false);
  const [isCheckoutOpen, setIsCheckoutOpen] = useState(false);
  const [isAuthModalOpen, setIsAuthModalOpen] = useState(false);
  const [isOrderHistoryOpen, setIsOrderHistoryOpen] = useState(false);
  const [orderConfirmation, setOrderConfirmation] = useState<{
    data: CheckoutData;
    orderNumber: string;
  } | null>(null);
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [currentPage, setCurrentPage] = useState<'home' | 'seller' | 'seller-login'>('home');
  const [selectedCategory, setSelectedCategory] = useState<string>('all');
  const [searchQuery, setSearchQuery] = useState('');
  const [bannerSearchQuery, setBannerSearchQuery] = useState('');
  const [userOrders, setUserOrders] = useState<Order[]>([]);

  useEffect(() => {
    const mockProducts = generateMockProducts();
    setProducts(mockProducts);
    setFilteredProducts(mockProducts);
  }, []);

  useEffect(() => {
    let filtered = products;

    if (selectedCategory !== 'all') {
      filtered = filtered.filter(product => product.category === selectedCategory);
    }

    if (searchQuery) {
      filtered = filtered.filter(product =>
        product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        product.description.toLowerCase().includes(searchQuery.toLowerCase())
      );
    }

    setFilteredProducts(filtered);
  }, [products, selectedCategory, searchQuery]);

  useEffect(() => {
    if (currentUser) {
      const mockOrders = createMockOrders(currentUser.id);
      setUserOrders(mockOrders);
    }
  }, [currentUser]);

  const addToCart = (product: Product, quantity: number = 1) => {
    setCartItems(prev => {
      const existingItem = prev.find(item => item.product.id === product.id);
      if (existingItem) {
        return prev.map(item =>
          item.product.id === product.id
            ? { ...item, quantity: item.quantity + quantity }
            : item
        );
      }
      return [...prev, { product, quantity }];
    });
  };

  const updateCartQuantity = (productId: string, quantity: number) => {
    if (quantity <= 0) {
      setCartItems(prev => prev.filter(item => item.product.id !== productId));
    } else {
      setCartItems(prev =>
        prev.map(item =>
          item.product.id === productId ? { ...item, quantity } : item
        )
      );
    }
  };

  const removeFromCart = (productId: string) => {
    setCartItems(prev => prev.filter(item => item.product.id !== productId));
  };

  const addProduct = (newProduct: Omit<Product, 'id'>) => {
    const product: Product = {
      ...newProduct,
      id: Date.now().toString(),
    };
    setProducts(prev => [product, ...prev]);
  };

  const updateUser = (updatedUser: User) => {
    setCurrentUser(updatedUser);
  };

  const handleLogin = (user: User) => {
    // Add default seller profile for sellers
    if (user.isSeller && !user.sellerProfile) {
      const defaultProfile: SellerProfile = {
        businessName: user.name + "'s Store",
        description: 'Welcome to my store! I offer quality products with excellent customer service.',
        location: '',
        phone: '',
        rating: 0,
        totalSales: 0,
        joinDate: new Date().toISOString(),
        verified: false,
      };
      user.sellerProfile = defaultProfile;
    }
    setCurrentUser(user);

    // Redirect to appropriate page after login
    if (user.isSeller && currentPage === 'seller-login') {
      setCurrentPage('seller');
    } else if (currentPage === 'seller-login') {
      setCurrentPage('home');
    }
  };

  const handleCheckout = () => {
    setIsCheckoutOpen(true);
  };

  const handleOrderComplete = (orderData: CheckoutData) => {
    const orderNumber = generateOrderNumber();
    const trackingNumber = generateTrackingNumber();
    const estimatedDelivery = calculateEstimatedDelivery(orderData.shippingMethod);

    // Create new order
    const newOrder: Order = {
      id: Date.now().toString(),
      orderNumber,
      buyerId: currentUser?.id || 'guest',
      sellerId: orderData.items[0]?.product.sellerId || '',
      products: orderData.items,
      total: orderData.total,
      subtotal: orderData.subtotal,
      tax: orderData.tax,
      shipping: orderData.shipping,
      status: 'confirmed',
      orderDate: new Date().toISOString(),
      shippingAddress: orderData.shippingAddress,
      billingAddress: orderData.billingAddress,
      paymentMethod: orderData.paymentMethod,
      shippingMethod: orderData.shippingMethod,
      trackingNumber,
      estimatedDelivery,
      trackingEvents: createMockTrackingEvents('confirmed'),
    };

    // Add order to user's order history
    setUserOrders(prev => [newOrder, ...prev]);

    setOrderConfirmation({ data: orderData, orderNumber });
    setIsCheckoutOpen(false);
    setCartItems([]); // Clear cart after successful order
  };

  const handleOrderConfirmationClose = () => {
    setOrderConfirmation(null);
  };

  const handlePageChange = (page: 'home' | 'seller') => {
    if (page === 'seller' && (!currentUser || !currentUser.isSeller)) {
      setCurrentPage('seller-login');
    } else {
      setCurrentPage(page);
    }
  };

  const handleCategoryClick = (category: string) => {
    setSelectedCategory(category);
  };

  const handleBannerSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setSearchQuery(bannerSearchQuery);
    // Scroll to products section
    const productsSection = document.getElementById('products-section');
    if (productsSection) {
      productsSection.scrollIntoView({ behavior: 'smooth' });
    }
  };

  const totalItems = cartItems.reduce((sum, item) => sum + item.quantity, 0);
  const totalPrice = cartItems.reduce((sum, item) => sum + (item.product.price * item.quantity), 0);

  const categories = ['all', ...Array.from(new Set(products.map(p => p.category)))];

  // Set URL
  useEffect(() => {
      if (currentPage === 'seller-login') {
        window.history.pushState({}, '', '/become-seller');
      } else if (currentPage === 'home') {
        window.history.pushState({}, '', '/');
      } else if (currentPage === 'seller') {
        window.history.pushState({}, '', '/dashboard'); // optional
      }
  }, [currentPage]);

  // Render seller login page
  if (currentPage === 'seller-login') {
    return <SellerLoginPage onLogin={handleLogin} />;
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <Header
        cartItemCount={totalItems}
        onCartClick={() => setIsCartOpen(true)}
        onAuthClick={() => setIsAuthModalOpen(true)}
        onOrderHistoryClick={() => setIsOrderHistoryOpen(true)}
        currentUser={currentUser}
        onPageChange={handlePageChange}
        currentPage={currentPage}
        onSearch={setSearchQuery}
        searchQuery={searchQuery}
        products={products}
        onFilteredProducts={setFilteredProducts}
      />

      <main className="pt-16">
        {currentPage === 'home' ? (
          <>
            <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white">
              <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
                <div className="text-center">
                  <h1 className="text-4xl md:text-6xl font-bold mb-4">
                    Discover Amazing Products
                  </h1>
                  <p className="text-xl md:text-2xl mb-8 text-blue-100">
                    Shop from thousands of sellers worldwide
                  </p>

                  {/* Banner Search Field */}
                  <div className="max-w-2xl mx-auto">
                    <form onSubmit={handleBannerSearch} className="relative">
                      <div className="relative">
                        <input
                          type="text"
                          placeholder="What are you looking for today?"
                          value={bannerSearchQuery}
                          onChange={(e) => setBannerSearchQuery(e.target.value)}
                          className="w-full px-6 py-4 pr-16 text-gray-900 rounded-full text-lg focus:outline-none focus:ring-4 focus:ring-blue-300 shadow-lg"
                        />
                        <button
                          type="submit"
                          className="absolute right-2 top-1/2 transform -translate-y-1/2 bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-full transition-colors font-medium"
                        >
                          Search
                        </button>
                      </div>
                    </form>

                    {/* Popular Search Suggestions */}
                    <div className="mt-4 flex flex-wrap justify-center gap-2">
                      <span className="text-blue-200 text-sm">Popular searches:</span>
                      {['headphones', 'laptop', 'fashion', 'home decor', 'fitness'].map((term) => (
                        <button
                          key={term}
                          onClick={() => {
                            setBannerSearchQuery(term);
                            setSearchQuery(term);
                          }}
                          className="text-blue-200 hover:text-white text-sm underline hover:no-underline transition-colors"
                        >
                          {term}
                        </button>
                      ))}
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div id="products-section" className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
              {/* Flash Sale Section */}
              <FlashSale
                products={products}
                onProductClick={setSelectedProduct}
                onAddToCart={addToCart}
              />

              {/* Category Showcase */}
              <CategoryShowcase
                products={products}
                onCategoryClick={handleCategoryClick}
                onProductClick={setSelectedProduct}
              />

              {/* Products for You */}
              <ProductsForYou
                products={products}
                onProductClick={setSelectedProduct}
                onAddToCart={addToCart}
                currentUser={currentUser}
              />

              {/* Category Filter Menu */}
              <div className="mb-8">
                <div className="flex flex-wrap gap-4 justify-center">
                  {categories.map(category => (
                    <button
                      key={category}
                      onClick={() => setSelectedCategory(category)}
                      className={`px-6 py-3 rounded-full font-medium transition-all ${
                        selectedCategory === category
                          ? 'bg-blue-600 text-white shadow-lg'
                          : 'bg-white text-gray-700 hover:bg-gray-100 shadow-md'
                      }`}
                    >
                      {category.charAt(0).toUpperCase() + category.slice(1)}
                    </button>
                  ))}
                </div>
              </div>

              {/* Main Product Grid */}
              <div className="mb-8">
                <h2 className="text-2xl font-bold text-gray-900 mb-6">
                  {selectedCategory === 'all' ? 'All Products' : `${selectedCategory.charAt(0).toUpperCase() + selectedCategory.slice(1)} Products`}
                  <span className="text-gray-500 text-lg ml-2">({filteredProducts.length})</span>
                </h2>
                <ProductGrid
                  products={filteredProducts}
                  onProductClick={setSelectedProduct}
                  onAddToCart={addToCart}
                />
              </div>
            </div>
          </>
        ) : (
          <SellerDashboard
            currentUser={currentUser}
            onAddProduct={addProduct}
            products={products.filter(p => p.sellerId === currentUser?.id)}
            onUpdateUser={updateUser}
          />
        )}
      </main>

      <Footer onNavigate={setCurrentPage} />

      {selectedProduct && (
        <ProductDetail
          product={selectedProduct}
          onClose={() => setSelectedProduct(null)}
          onAddToCart={addToCart}
        />
      )}

      {isCartOpen && (
        <Cart
          items={cartItems}
          onClose={() => setIsCartOpen(false)}
          onUpdateQuantity={updateCartQuantity}
          onRemoveItem={removeFromCart}
          onCheckout={handleCheckout}
          totalPrice={totalPrice}
        />
      )}

      {isCheckoutOpen && (
        <Checkout
          items={cartItems}
          onClose={() => setIsCheckoutOpen(false)}
          onOrderComplete={handleOrderComplete}
        />
      )}

      {orderConfirmation && (
        <OrderConfirmation
          orderData={orderConfirmation.data}
          orderNumber={orderConfirmation.orderNumber}
          onClose={handleOrderConfirmationClose}
        />
      )}

      {isOrderHistoryOpen && currentUser && (
        <OrderHistory
          orders={userOrders}
          onClose={() => setIsOrderHistoryOpen(false)}
        />
      )}

      {isAuthModalOpen && (
        <AuthModal
          onClose={() => setIsAuthModalOpen(false)}
          onLogin={handleLogin}
        />
      )}
    </div>
  );
}

export default App;
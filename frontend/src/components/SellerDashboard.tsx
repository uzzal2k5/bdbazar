import React, { useState } from 'react';
import { Plus, Package, TrendingUp, DollarSign, Users, Eye, Edit, BarChart3, Settings } from 'lucide-react';
import { Product, User, Order } from '../types';
import { AddProductForm } from './AddProductForm';
import { SellerProfile } from './SellerProfile';

interface SellerDashboardProps {
  currentUser: User | null;
  onAddProduct: (product: Omit<Product, 'id'>) => void;
  products: Product[];
  onUpdateUser?: (user: User) => void;
}

export const SellerDashboard: React.FC<SellerDashboardProps> = ({
  currentUser,
  onAddProduct,
  products,
  onUpdateUser,
}) => {
  const [showAddForm, setShowAddForm] = useState(false);
  const [showProfile, setShowProfile] = useState(false);
  const [activeTab, setActiveTab] = useState<'overview' | 'products' | 'orders' | 'analytics'>('overview');

  if (!currentUser) {
    return (
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="text-center">
          <h1 className="text-3xl font-bold text-gray-900 mb-4">Please log in to access seller dashboard</h1>
        </div>
      </div>
    );
  }

  const totalRevenue = products.reduce((sum, product) => sum + (product.price * 10), 0);
  const totalProducts = products.length;
  const averageRating = products.length > 0 ? products.reduce((sum, product) => sum + product.rating, 0) / products.length : 0;
  const totalViews = products.reduce((sum) => sum + Math.floor(Math.random() * 1000), 0);

  const handleUpdateProfile = (profile: any) => {
    if (onUpdateUser && currentUser) {
      const updatedUser = {
        ...currentUser,
        sellerProfile: profile,
      };
      onUpdateUser(updatedUser);
    }
  };

  const mockOrders: Order[] = [
    {
      id: '1',
      buyerId: 'buyer1',
      sellerId: currentUser.id,
      products: products.slice(0, 2).map(p => ({ product: p, quantity: 1 })),
      total: 124.98,
      status: 'processing',
      orderDate: '2024-01-20',
      shippingAddress: '123 Main St, City, State 12345',
    },
    {
      id: '2',
      buyerId: 'buyer2',
      sellerId: currentUser.id,
      products: products.slice(1, 3).map(p => ({ product: p, quantity: 2 })),
      total: 299.97,
      status: 'shipped',
      orderDate: '2024-01-19',
      shippingAddress: '456 Oak Ave, City, State 67890',
    },
  ];

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'pending': return 'bg-yellow-100 text-yellow-800';
      case 'processing': return 'bg-blue-100 text-blue-800';
      case 'shipped': return 'bg-purple-100 text-purple-800';
      case 'delivered': return 'bg-green-100 text-green-800';
      case 'cancelled': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8">
        <div className="flex justify-between items-start">
          <div>
            <h1 className="text-3xl font-bold text-gray-900 mb-2">Seller Dashboard</h1>
            <p className="text-gray-600">Welcome back, {currentUser.name}!</p>
          </div>
          <button
            onClick={() => setShowProfile(true)}
            className="flex items-center space-x-2 px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors"
          >
            <Settings className="h-4 w-4" />
            <span>Manage Profile</span>
          </button>
        </div>
      </div>

      {/* Navigation Tabs */}
      <div className="border-b border-gray-200 mb-8">
        <nav className="flex space-x-8">
          {[
            { id: 'overview', label: 'Overview', icon: BarChart3 },
            { id: 'products', label: 'Products', icon: Package },
            { id: 'orders', label: 'Orders', icon: Users },
            { id: 'analytics', label: 'Analytics', icon: TrendingUp },
          ].map((tab) => {
            const Icon = tab.icon;
            return (
              <button
                key={tab.id}
                onClick={() => setActiveTab(tab.id as any)}
                className={`flex items-center space-x-2 py-4 px-1 border-b-2 font-medium text-sm transition-colors ${
                  activeTab === tab.id
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
              >
                <Icon className="h-4 w-4" />
                <span>{tab.label}</span>
              </button>
            );
          })}
        </nav>
      </div>

      {/* Overview Tab */}
      {activeTab === 'overview' && (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
            <div className="bg-white rounded-xl shadow-md p-6 border border-gray-100">
              <div className="flex items-center">
                <DollarSign className="h-8 w-8 text-green-600" />
                <div className="ml-4">
                  <p className="text-sm text-gray-600">Total Revenue</p>
                  <p className="text-2xl font-bold text-gray-900">${totalRevenue.toFixed(2)}</p>
                </div>
              </div>
            </div>

            <div className="bg-white rounded-xl shadow-md p-6 border border-gray-100">
              <div className="flex items-center">
                <Package className="h-8 w-8 text-blue-600" />
                <div className="ml-4">
                  <p className="text-sm text-gray-600">Total Products</p>
                  <p className="text-2xl font-bold text-gray-900">{totalProducts}</p>
                </div>
              </div>
            </div>

            <div className="bg-white rounded-xl shadow-md p-6 border border-gray-100">
              <div className="flex items-center">
                <TrendingUp className="h-8 w-8 text-purple-600" />
                <div className="ml-4">
                  <p className="text-sm text-gray-600">Average Rating</p>
                  <p className="text-2xl font-bold text-gray-900">{averageRating.toFixed(1)}</p>
                </div>
              </div>
            </div>

            <div className="bg-white rounded-xl shadow-md p-6 border border-gray-100">
              <div className="flex items-center">
                <Eye className="h-8 w-8 text-orange-600" />
                <div className="ml-4">
                  <p className="text-sm text-gray-600">Total Views</p>
                  <p className="text-2xl font-bold text-gray-900">{totalViews.toLocaleString()}</p>
                </div>
              </div>
            </div>
          </div>

          {/* Recent Orders */}
          <div className="bg-white rounded-xl shadow-md border border-gray-100 mb-8">
            <div className="border-b border-gray-200 p-6">
              <h2 className="text-xl font-semibold">Recent Orders</h2>
            </div>
            <div className="p-6">
              <div className="space-y-4">
                {mockOrders.map((order) => (
                  <div key={order.id} className="flex items-center justify-between p-4 border border-gray-200 rounded-lg">
                    <div>
                      <p className="font-medium">Order #{order.id}</p>
                      <p className="text-sm text-gray-600">{order.products.length} items • ${order.total}</p>
                      <p className="text-sm text-gray-500">{order.orderDate}</p>
                    </div>
                    <span className={`px-3 py-1 rounded-full text-sm font-medium ${getStatusColor(order.status)}`}>
                      {order.status.charAt(0).toUpperCase() + order.status.slice(1)}
                    </span>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </>
      )}

      {/* Products Tab */}
      {activeTab === 'products' && (
        <div className="bg-white rounded-xl shadow-md border border-gray-100">
          <div className="border-b border-gray-200 p-6">
            <div className="flex justify-between items-center">
              <h2 className="text-xl font-semibold">Your Products</h2>
              <button
                onClick={() => setShowAddForm(true)}
                className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors flex items-center space-x-2"
              >
                <Plus className="h-4 w-4" />
                <span>Add Product</span>
              </button>
            </div>
          </div>

          <div className="p-6">
            {products.length === 0 ? (
              <div className="text-center py-12">
                <Package className="h-16 w-16 text-gray-400 mx-auto mb-4" />
                <p className="text-gray-500 text-lg mb-4">No products yet</p>
                <p className="text-gray-400 mb-6">Add your first product to get started selling!</p>
                <button
                  onClick={() => setShowAddForm(true)}
                  className="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors"
                >
                  Add Your First Product
                </button>
              </div>
            ) : (
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {products.map((product) => (
                  <div key={product.id} className="border border-gray-200 rounded-lg overflow-hidden hover:shadow-lg transition-shadow">
                    <img
                      src={product.image}
                      alt={product.name}
                      className="w-full h-48 object-cover"
                    />
                    <div className="p-4">
                      <h3 className="font-semibold mb-2">{product.name}</h3>
                      <p className="text-gray-600 text-sm mb-3 line-clamp-2">{product.description}</p>
                      <div className="flex justify-between items-center mb-3">
                        <span className="text-lg font-bold text-blue-600">${product.price}</span>
                        <span className="text-sm text-gray-500">{product.stock} in stock</span>
                      </div>
                      <div className="flex justify-between items-center">
                        <div className="flex items-center space-x-1">
                          <span className="text-sm text-gray-600">Rating:</span>
                          <span className="text-sm font-medium">{product.rating}</span>
                        </div>
                        <button className="p-2 text-gray-600 hover:text-blue-600 transition-colors">
                          <Edit className="h-4 w-4" />
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
      )}

      {/* Orders Tab */}
      {activeTab === 'orders' && (
        <div className="bg-white rounded-xl shadow-md border border-gray-100">
          <div className="border-b border-gray-200 p-6">
            <h2 className="text-xl font-semibold">Order Management</h2>
          </div>
          <div className="p-6">
            <div className="space-y-6">
              {mockOrders.map((order) => (
                <div key={order.id} className="border border-gray-200 rounded-lg p-6">
                  <div className="flex justify-between items-start mb-4">
                    <div>
                      <h3 className="font-semibold text-lg">Order #{order.id}</h3>
                      <p className="text-gray-600">Placed on {order.orderDate}</p>
                    </div>
                    <span className={`px-3 py-1 rounded-full text-sm font-medium ${getStatusColor(order.status)}`}>
                      {order.status.charAt(0).toUpperCase() + order.status.slice(1)}
                    </span>
                  </div>

                  <div className="space-y-3 mb-4">
                    {order.products.map((item, index) => (
                      <div key={index} className="flex items-center space-x-4">
                        <img
                          src={item.product.image}
                          alt={item.product.name}
                          className="w-12 h-12 object-cover rounded"
                        />
                        <div className="flex-1">
                          <p className="font-medium">{item.product.name}</p>
                          <p className="text-sm text-gray-600">Quantity: {item.quantity}</p>
                        </div>
                        <p className="font-medium">${(item.product.price * item.quantity).toFixed(2)}</p>
                      </div>
                    ))}
                  </div>

                  <div className="border-t border-gray-200 pt-4">
                    <div className="flex justify-between items-center">
                      <div>
                        <p className="text-sm text-gray-600">Shipping Address:</p>
                        <p className="text-sm">{order.shippingAddress}</p>
                      </div>
                      <div className="text-right">
                        <p className="text-lg font-bold">Total: ${order.total}</p>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}

      {/* Analytics Tab */}
      {activeTab === 'analytics' && (
        <div className="space-y-6">
          <div className="bg-white rounded-xl shadow-md border border-gray-100 p-6">
            <h2 className="text-xl font-semibold mb-4">Sales Analytics</h2>
            <div className="text-center py-12">
              <BarChart3 className="h-16 w-16 text-gray-400 mx-auto mb-4" />
              <p className="text-gray-500">Advanced analytics coming soon!</p>
              <p className="text-gray-400 text-sm">Track your sales performance, customer insights, and growth metrics.</p>
            </div>
          </div>
        </div>
      )}

      {showAddForm && (
        <AddProductForm
          onClose={() => setShowAddForm(false)}
          onAddProduct={onAddProduct}
          sellerId={currentUser.id}
          sellerName={currentUser.name}
        />
      )}

      {showProfile && (
        <SellerProfile
          seller={currentUser}
          onClose={() => setShowProfile(false)}
          onUpdateProfile={handleUpdateProfile}
          isOwnProfile={true}
        />
      )}
    </div>
  );
};
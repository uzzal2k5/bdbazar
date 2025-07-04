import React, { useState } from 'react';
import { Package, Eye, Download, RefreshCw } from 'lucide-react';
import { Order } from '../types';
import { OrderTracking } from './OrderTracking';

interface OrderHistoryProps {
  orders: Order[];
  onClose: () => void;
}

export const OrderHistory: React.FC<OrderHistoryProps> = ({ orders, onClose }) => {
  const [selectedOrder, setSelectedOrder] = useState<Order | null>(null);
  const [statusFilter, setStatusFilter] = useState<string>('all');

  const filteredOrders = orders.filter(order =>
    statusFilter === 'all' || order.status === statusFilter
  );

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'pending': return 'bg-yellow-100 text-yellow-800';
      case 'confirmed':
      case 'processing': return 'bg-blue-100 text-blue-800';
      case 'shipped':
      case 'out_for_delivery': return 'bg-purple-100 text-purple-800';
      case 'delivered': return 'bg-green-100 text-green-800';
      case 'cancelled':
      case 'returned': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    });
  };

  return (
    <>
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
        <div className="bg-white rounded-2xl max-w-6xl w-full max-h-[90vh] overflow-y-auto">
          <div className="sticky top-0 bg-white border-b border-gray-200 p-6">
            <div className="flex justify-between items-center">
              <h2 className="text-2xl font-bold text-gray-900">Order History</h2>
              <button
                onClick={onClose}
                className="p-2 hover:bg-gray-100 rounded-full transition-colors"
              >
                ✕
              </button>
            </div>

            {/* Status Filter */}
            <div className="mt-4 flex flex-wrap gap-2">
              {['all', 'pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled'].map((status) => (
                <button
                  key={status}
                  onClick={() => setStatusFilter(status)}
                  className={`px-4 py-2 rounded-full text-sm font-medium transition-colors ${
                    statusFilter === status
                      ? 'bg-blue-600 text-white'
                      : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
                  }`}
                >
                  {status.charAt(0).toUpperCase() + status.slice(1).replace('_', ' ')}
                </button>
              ))}
            </div>
          </div>

          <div className="p-6">
            {filteredOrders.length === 0 ? (
              <div className="text-center py-12">
                <Package className="h-16 w-16 text-gray-400 mx-auto mb-4" />
                <p className="text-gray-500 text-lg mb-2">No orders found</p>
                <p className="text-gray-400">
                  {statusFilter === 'all'
                    ? "You haven't placed any orders yet."
                    : `No orders with status "${statusFilter}".`
                  }
                </p>
              </div>
            ) : (
              <div className="space-y-4">
                {filteredOrders.map((order) => (
                  <div key={order.id} className="border border-gray-200 rounded-xl p-6 hover:shadow-lg transition-shadow">
                    <div className="flex justify-between items-start mb-4">
                      <div>
                        <h3 className="font-semibold text-lg">Order #{order.orderNumber}</h3>
                        <p className="text-gray-600">Placed on {formatDate(order.orderDate)}</p>
                        {order.trackingNumber && (
                          <p className="text-sm text-gray-500 font-mono">
                            Tracking: {order.trackingNumber}
                          </p>
                        )}
                      </div>
                      <div className="text-right">
                        <span className={`px-3 py-1 rounded-full text-sm font-medium ${getStatusColor(order.status)}`}>
                          {order.status.replace('_', ' ').toUpperCase()}
                        </span>
                        <p className="text-lg font-bold mt-2">${order.total.toFixed(2)}</p>
                      </div>
                    </div>

                    {/* Order Items Preview */}
                    <div className="mb-4">
                      <div className="flex items-center space-x-4 overflow-x-auto pb-2">
                        {order.products.slice(0, 3).map((item, index) => (
                          <div key={index} className="flex items-center space-x-3 min-w-0">
                            <img
                              src={item.product.image}
                              alt={item.product.name}
                              className="w-12 h-12 object-cover rounded-lg flex-shrink-0"
                            />
                            <div className="min-w-0">
                              <p className="font-medium text-sm truncate">{item.product.name}</p>
                              <p className="text-xs text-gray-600">Qty: {item.quantity}</p>
                            </div>
                          </div>
                        ))}
                        {order.products.length > 3 && (
                          <div className="text-sm text-gray-500 flex-shrink-0">
                            +{order.products.length - 3} more items
                          </div>
                        )}
                      </div>
                    </div>

                    {/* Shipping Info */}
                    <div className="bg-gray-50 rounded-lg p-4 mb-4">
                      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm">
                        <div>
                          <p className="text-gray-600">Shipping Method</p>
                          <p className="font-medium">{order.shippingMethod.name}</p>
                        </div>
                        <div>
                          <p className="text-gray-600">Shipping Address</p>
                          <p className="font-medium">
                            {order.shippingAddress.city}, {order.shippingAddress.state}
                          </p>
                        </div>
                        {order.estimatedDelivery && (
                          <div>
                            <p className="text-gray-600">Estimated Delivery</p>
                            <p className="font-medium">{formatDate(order.estimatedDelivery)}</p>
                          </div>
                        )}
                      </div>
                    </div>

                    {/* Action Buttons */}
                    <div className="flex flex-wrap gap-3">
                      <button
                        onClick={() => setSelectedOrder(order)}
                        className="flex items-center space-x-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                      >
                        <Eye className="h-4 w-4" />
                        <span>Track Order</span>
                      </button>

                      <button className="flex items-center space-x-2 px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors">
                        <Download className="h-4 w-4" />
                        <span>Download Invoice</span>
                      </button>

                      {order.status === 'delivered' && (
                        <button className="flex items-center space-x-2 px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors">
                          <RefreshCw className="h-4 w-4" />
                          <span>Reorder</span>
                        </button>
                      )}

                      {['pending', 'confirmed'].includes(order.status) && (
                        <button className="flex items-center space-x-2 px-4 py-2 border border-red-300 text-red-700 rounded-lg hover:bg-red-50 transition-colors">
                          <span>Cancel Order</span>
                        </button>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
      </div>

      {selectedOrder && (
        <OrderTracking
          order={selectedOrder}
          onClose={() => setSelectedOrder(null)}
        />
      )}
    </>
  );
};
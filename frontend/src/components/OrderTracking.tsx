import React from 'react';
import { Package, Truck, MapPin, CheckCircle, Clock, AlertCircle } from 'lucide-react';
import { Order, TrackingEvent } from '../types';

interface OrderTrackingProps {
  order: Order;
  onClose: () => void;
}

export const OrderTracking: React.FC<OrderTrackingProps> = ({ order, onClose }) => {
  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'confirmed':
      case 'processing':
        return Package;
      case 'shipped':
      case 'out_for_delivery':
        return Truck;
      case 'delivered':
        return CheckCircle;
      default:
        return Clock;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'pending':
        return 'text-yellow-600 bg-yellow-100';
      case 'confirmed':
      case 'processing':
        return 'text-blue-600 bg-blue-100';
      case 'shipped':
      case 'out_for_delivery':
        return 'text-purple-600 bg-purple-100';
      case 'delivered':
        return 'text-green-600 bg-green-100';
      case 'cancelled':
      case 'returned':
        return 'text-red-600 bg-red-100';
      default:
        return 'text-gray-600 bg-gray-100';
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      weekday: 'short',
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-2xl max-w-4xl w-full max-h-[90vh] overflow-y-auto">
        <div className="sticky top-0 bg-white border-b border-gray-200 p-6">
          <div className="flex justify-between items-start">
            <div>
              <h2 className="text-2xl font-bold text-gray-900">Order Tracking</h2>
              <p className="text-gray-600">Order #{order.orderNumber}</p>
            </div>
            <button
              onClick={onClose}
              className="p-2 hover:bg-gray-100 rounded-full transition-colors"
            >
              ✕
            </button>
          </div>
        </div>

        <div className="p-6">
          {/* Order Status Header */}
          <div className="bg-gradient-to-r from-blue-50 to-purple-50 rounded-xl p-6 mb-8">
            <div className="flex items-center justify-between mb-4">
              <div>
                <h3 className="text-xl font-semibold text-gray-900">Current Status</h3>
                <p className={`inline-flex items-center px-3 py-1 rounded-full text-sm font-medium mt-2 ${getStatusColor(order.status)}`}>
                  {order.status.replace('_', ' ').toUpperCase()}
                </p>
              </div>
              {order.trackingNumber && (
                <div className="text-right">
                  <p className="text-sm text-gray-600">Tracking Number</p>
                  <p className="font-mono font-semibold text-lg">{order.trackingNumber}</p>
                </div>
              )}
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm">
              <div>
                <p className="text-gray-600">Order Date</p>
                <p className="font-medium">{formatDate(order.orderDate)}</p>
              </div>
              {order.estimatedDelivery && (
                <div>
                  <p className="text-gray-600">Estimated Delivery</p>
                  <p className="font-medium">{formatDate(order.estimatedDelivery)}</p>
                </div>
              )}
              <div>
                <p className="text-gray-600">Shipping Method</p>
                <p className="font-medium">{order.shippingMethod.name}</p>
              </div>
            </div>
          </div>

          {/* Tracking Timeline */}
          <div className="mb-8">
            <h3 className="text-lg font-semibold mb-6">Tracking History</h3>
            <div className="space-y-4">
              {order.trackingEvents.map((event, index) => {
                const Icon = getStatusIcon(event.status);
                const isLast = index === order.trackingEvents.length - 1;

                return (
                  <div key={event.id} className="flex items-start space-x-4">
                    <div className="flex flex-col items-center">
                      <div className={`w-10 h-10 rounded-full flex items-center justify-center ${
                        event.isCompleted ? 'bg-green-100 text-green-600' : 'bg-gray-100 text-gray-400'
                      }`}>
                        <Icon className="h-5 w-5" />
                      </div>
                      {!isLast && (
                        <div className={`w-0.5 h-12 mt-2 ${
                          event.isCompleted ? 'bg-green-200' : 'bg-gray-200'
                        }`} />
                      )}
                    </div>

                    <div className="flex-1 pb-8">
                      <div className="flex justify-between items-start mb-1">
                        <h4 className={`font-medium ${
                          event.isCompleted ? 'text-gray-900' : 'text-gray-500'
                        }`}>
                          {event.status.replace('_', ' ').toUpperCase()}
                        </h4>
                        <span className="text-sm text-gray-500">
                          {formatDate(event.timestamp)}
                        </span>
                      </div>

                      <p className={`text-sm ${
                        event.isCompleted ? 'text-gray-700' : 'text-gray-500'
                      }`}>
                        {event.description}
                      </p>

                      {event.location && (
                        <div className="flex items-center space-x-1 mt-1">
                          <MapPin className="h-3 w-3 text-gray-400" />
                          <span className="text-xs text-gray-500">{event.location}</span>
                        </div>
                      )}
                    </div>
                  </div>
                );
              })}
            </div>
          </div>

          {/* Order Items */}
          <div className="bg-gray-50 rounded-xl p-6 mb-6">
            <h3 className="font-semibold mb-4">Order Items</h3>
            <div className="space-y-4">
              {order.products.map((item, index) => (
                <div key={index} className="flex items-center space-x-4">
                  <img
                    src={item.product.image}
                    alt={item.product.name}
                    className="w-16 h-16 object-cover rounded-lg"
                  />
                  <div className="flex-1">
                    <h4 className="font-medium">{item.product.name}</h4>
                    <p className="text-sm text-gray-600">Quantity: {item.quantity}</p>
                    <p className="text-sm text-gray-600">Seller: {item.product.sellerName}</p>
                  </div>
                  <p className="font-medium">${(item.product.price * item.quantity).toFixed(2)}</p>
                </div>
              ))}
            </div>
          </div>

          {/* Shipping Address */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="bg-gray-50 rounded-xl p-6">
              <h3 className="font-semibold mb-3 flex items-center space-x-2">
                <MapPin className="h-5 w-5" />
                <span>Shipping Address</span>
              </h3>
              <div className="text-sm text-gray-700">
                <p className="font-medium">
                  {order.shippingAddress.firstName} {order.shippingAddress.lastName}
                </p>
                <p>{order.shippingAddress.address1}</p>
                {order.shippingAddress.address2 && <p>{order.shippingAddress.address2}</p>}
                <p>
                  {order.shippingAddress.city}, {order.shippingAddress.state} {order.shippingAddress.zipCode}
                </p>
                {order.shippingAddress.phone && <p>Phone: {order.shippingAddress.phone}</p>}
              </div>
            </div>

            <div className="bg-gray-50 rounded-xl p-6">
              <h3 className="font-semibold mb-3 flex items-center space-x-2">
                <Truck className="h-5 w-5" />
                <span>Shipping Details</span>
              </h3>
              <div className="text-sm text-gray-700 space-y-2">
                <div className="flex justify-between">
                  <span>Method:</span>
                  <span className="font-medium">{order.shippingMethod.name}</span>
                </div>
                <div className="flex justify-between">
                  <span>Courier:</span>
                  <span className="font-medium">{order.shippingMethod.courier}</span>
                </div>
                <div className="flex justify-between">
                  <span>Cost:</span>
                  <span className="font-medium">
                    {order.shipping === 0 ? 'Free' : `$${order.shipping.toFixed(2)}`}
                  </span>
                </div>
                {order.estimatedDelivery && (
                  <div className="flex justify-between">
                    <span>Est. Delivery:</span>
                    <span className="font-medium">
                      {new Date(order.estimatedDelivery).toLocaleDateString()}
                    </span>
                  </div>
                )}
              </div>
            </div>
          </div>

          {/* Order Summary */}
          <div className="mt-6 bg-gray-50 rounded-xl p-6">
            <h3 className="font-semibold mb-4">Order Summary</h3>
            <div className="space-y-2 text-sm">
              <div className="flex justify-between">
                <span>Subtotal:</span>
                <span>${order.subtotal.toFixed(2)}</span>
              </div>
              <div className="flex justify-between">
                <span>Shipping:</span>
                <span>{order.shipping === 0 ? 'Free' : `$${order.shipping.toFixed(2)}`}</span>
              </div>
              <div className="flex justify-between">
                <span>Tax:</span>
                <span>${order.tax.toFixed(2)}</span>
              </div>
              <div className="border-t border-gray-300 pt-2 mt-2">
                <div className="flex justify-between font-semibold text-lg">
                  <span>Total:</span>
                  <span>${order.total.toFixed(2)}</span>
                </div>
              </div>
            </div>
          </div>

          {/* Action Buttons */}
          <div className="mt-6 flex flex-col sm:flex-row gap-3">
            <button
              onClick={onClose}
              className="flex-1 bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition-colors font-medium"
            >
              Close Tracking
            </button>

            {order.status === 'delivered' && (
              <button className="flex-1 border border-gray-300 text-gray-700 py-3 rounded-lg hover:bg-gray-50 transition-colors font-medium">
                Leave Review
              </button>
            )}

            {['pending', 'confirmed'].includes(order.status) && (
              <button className="flex-1 border border-red-300 text-red-700 py-3 rounded-lg hover:bg-red-50 transition-colors font-medium">
                Cancel Order
              </button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
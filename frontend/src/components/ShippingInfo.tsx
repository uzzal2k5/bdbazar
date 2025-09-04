import React from 'react';
import { Truck, Clock, MapPin, Package, Shield, DollarSign } from 'lucide-react';
import { ShippingMethod, Address } from '../types';

interface ShippingInfoProps {
  method: ShippingMethod;
  address: Address;
  trackingNumber?: string;
  estimatedDelivery?: string;
  weight?: number;
  showDetails?: boolean;
}

export const ShippingInfo: React.FC<ShippingInfoProps> = ({
  method,
  address,
  trackingNumber,
  estimatedDelivery,
  weight,
  showDetails = true,
}) => {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      weekday: 'long',
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
  };

  const getDeliveryIcon = () => {
    switch (method.icon) {
      case 'zap': return '⚡';
      case 'package': return '📦';
      default: return '🚚';
    }
  };

  return (
    <div className="bg-white rounded-lg border border-gray-200 p-6">
      <div className="flex items-center space-x-3 mb-4">
        <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center">
          <Truck className="h-6 w-6 text-blue-600" />
        </div>
        <div>
          <h3 className="text-lg font-semibold">{method.name}</h3>
          <p className="text-gray-600">{method.description}</p>
        </div>
        <div className="text-right">
          <span className="text-2xl">{getDeliveryIcon()}</span>
        </div>
      </div>

      {/* Shipping Details Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
        <div className="space-y-3">
          <div className="flex items-center space-x-3">
            <Clock className="h-5 w-5 text-gray-400" />
            <div>
              <p className="text-sm text-gray-600">Delivery Time</p>
              <p className="font-medium">{method.estimatedDays}</p>
            </div>
          </div>

          <div className="flex items-center space-x-3">
            <Package className="h-5 w-5 text-gray-400" />
            <div>
              <p className="text-sm text-gray-600">Courier</p>
              <p className="font-medium">{method.courier}</p>
            </div>
          </div>

          <div className="flex items-center space-x-3">
            <DollarSign className="h-5 w-5 text-gray-400" />
            <div>
              <p className="text-sm text-gray-600">Shipping Cost</p>
              <p className="font-medium">
                {method.price === 0 ? 'Free' : `$${method.price.toFixed(2)}`}
              </p>
            </div>
          </div>
        </div>

        <div className="space-y-3">
          {estimatedDelivery && (
            <div className="flex items-center space-x-3">
              <Clock className="h-5 w-5 text-green-500" />
              <div>
                <p className="text-sm text-gray-600">Estimated Delivery</p>
                <p className="font-medium text-green-700">{formatDate(estimatedDelivery)}</p>
              </div>
            </div>
          )}

          {trackingNumber && (
            <div className="flex items-center space-x-3">
              <Package className="h-5 w-5 text-blue-500" />
              <div>
                <p className="text-sm text-gray-600">Tracking Number</p>
                <p className="font-mono font-medium text-blue-700">{trackingNumber}</p>
              </div>
            </div>
          )}

          {weight && (
            <div className="flex items-center space-x-3">
              <Package className="h-5 w-5 text-gray-400" />
              <div>
                <p className="text-sm text-gray-600">Package Weight</p>
                <p className="font-medium">{weight.toFixed(1)} kg</p>
              </div>
            </div>
          )}
        </div>
      </div>

      {/* Shipping Address */}
      <div className="border-t border-gray-200 pt-4 mb-4">
        <div className="flex items-center space-x-2 mb-3">
          <MapPin className="h-5 w-5 text-gray-400" />
          <h4 className="font-medium">Shipping Address</h4>
        </div>
        <div className="bg-gray-50 rounded-lg p-3">
          <p className="font-medium">{address.firstName} {address.lastName}</p>
          {address.company && <p className="text-gray-600">{address.company}</p>}
          <p className="text-gray-600">{address.address1}</p>
          {address.address2 && <p className="text-gray-600">{address.address2}</p>}
          <p className="text-gray-600">
            {address.city}, {address.state} {address.zipCode}
          </p>
          <p className="text-gray-600">{address.country}</p>
          {address.phone && <p className="text-gray-600">Phone: {address.phone}</p>}
        </div>
      </div>

      {/* Features */}
      {showDetails && method.features.length > 0 && (
        <div className="border-t border-gray-200 pt-4">
          <div className="flex items-center space-x-2 mb-3">
            <Shield className="h-5 w-5 text-gray-400" />
            <h4 className="font-medium">Included Features</h4>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-2">
            {method.features.map((feature, index) => (
              <div key={index} className="flex items-center space-x-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                <span className="text-sm text-gray-600">{feature}</span>
              </div>
            ))}
          </div>
        </div>
      )}

      {/* COD Information */}
      {method.supportsCOD && (
        <div className="mt-4 p-3 bg-green-50 border border-green-200 rounded-lg">
          <div className="flex items-center space-x-2">
            <DollarSign className="h-4 w-4 text-green-600" />
            <span className="text-sm font-medium text-green-800">Cash on Delivery Available</span>
          </div>
          <p className="text-xs text-green-700 mt-1">
            You can pay in cash when your order is delivered
          </p>
        </div>
      )}
    </div>
  );
};
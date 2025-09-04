import React, { useState, useEffect } from 'react';
import { Calculator, MapPin, Package, Truck } from 'lucide-react';
import { CartItem, ShippingMethod, Address } from '../types';

interface ShippingCalculatorProps {
  items: CartItem[];
  destination?: Address;
  onShippingUpdate?: (cost: number, method: ShippingMethod) => void;
}

export const ShippingCalculator: React.FC<ShippingCalculatorProps> = ({
  items,
  destination,
  onShippingUpdate,
}) => {
  const [zipCode, setZipCode] = useState(destination?.zipCode || '');
  const [shippingMethods, setShippingMethods] = useState<ShippingMethod[]>([]);
  const [selectedMethod, setSelectedMethod] = useState<ShippingMethod | null>(null);
  const [isCalculating, setIsCalculating] = useState(false);

  const calculateWeight = () => {
    return items.reduce((total, item) => {
      const weight = item.product.weight || 0.5; // Default 0.5kg if not specified
      return total + (weight * item.quantity);
    }, 0);
  };

  const calculateDimensions = () => {
    // Simplified calculation - in reality, this would be more complex
    const totalVolume = items.reduce((total, item) => {
      const dims = item.product.dimensions || { length: 20, width: 15, height: 10 };
      const volume = dims.length * dims.width * dims.height * item.quantity;
      return total + volume;
    }, 0);

    // Convert to approximate box dimensions
    const side = Math.cbrt(totalVolume);
    return {
      length: Math.ceil(side),
      width: Math.ceil(side),
      height: Math.ceil(side * 0.8),
    };
  };

  const calculateShipping = async () => {
    if (!zipCode || zipCode.length < 5) return;

    setIsCalculating(true);

    // Simulate API call delay
    await new Promise(resolve => setTimeout(resolve, 1000));

    const weight = calculateWeight();
    const dimensions = calculateDimensions();
    const subtotal = items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0);

    // Mock shipping calculation based on weight, dimensions, and destination
    const baseRate = 5.99;
    const weightRate = weight * 2.50;
    const distanceMultiplier = getDistanceMultiplier(zipCode);
    const sizeMultiplier = getSizeMultiplier(dimensions);

    const methods: ShippingMethod[] = [
      {
        id: 'standard',
        name: 'Standard Shipping',
        description: 'Reliable delivery with tracking',
        price: subtotal > 50 ? 0 : Math.round((baseRate + weightRate) * distanceMultiplier * sizeMultiplier * 100) / 100,
        estimatedDays: '5-7 business days',
        courier: 'USPS',
        icon: 'truck',
        features: ['Tracking included', 'Insurance up to $100'],
        supportsCOD: true
      },
      {
        id: 'express',
        name: 'Express Shipping',
        description: 'Faster delivery for urgent orders',
        price: Math.round((baseRate + weightRate + 15) * distanceMultiplier * sizeMultiplier * 100) / 100,
        estimatedDays: '2-3 business days',
        courier: 'FedEx',
        icon: 'zap',
        features: ['Priority handling', 'Tracking included', 'Insurance up to $500'],
        supportsCOD: true
      },
      {
        id: 'overnight',
        name: 'Overnight Delivery',
        description: 'Next business day delivery',
        price: Math.round((baseRate + weightRate + 35) * distanceMultiplier * sizeMultiplier * 100) / 100,
        estimatedDays: '1 business day',
        courier: 'UPS',
        icon: 'package',
        features: ['Next day delivery', 'Signature required', 'Insurance up to $1000'],
        supportsCOD: false
      }
    ];

    setShippingMethods(methods);
    setSelectedMethod(methods[0]);
    setIsCalculating(false);
  };

  const getDistanceMultiplier = (zip: string): number => {
    // Mock distance calculation based on zip code
    const firstDigit = parseInt(zip.charAt(0));
    if (firstDigit <= 2) return 1.0; // East Coast
    if (firstDigit <= 5) return 1.2; // Central
    if (firstDigit <= 7) return 1.4; // Mountain
    return 1.6; // West Coast
  };

  const getSizeMultiplier = (dims: { length: number; width: number; height: number }): number => {
    const volume = dims.length * dims.width * dims.height;
    if (volume < 1000) return 1.0;
    if (volume < 5000) return 1.2;
    if (volume < 10000) return 1.5;
    return 2.0;
  };

  const handleMethodSelect = (method: ShippingMethod) => {
    setSelectedMethod(method);
    if (onShippingUpdate) {
      onShippingUpdate(method.price, method);
    }
  };

  useEffect(() => {
    if (destination?.zipCode) {
      setZipCode(destination.zipCode);
      calculateShipping();
    }
  }, [destination, items]);

  return (
    <div className="bg-white rounded-lg border border-gray-200 p-6">
      <div className="flex items-center space-x-2 mb-4">
        <Calculator className="h-5 w-5 text-blue-600" />
        <h3 className="text-lg font-semibold">Shipping Calculator</h3>
      </div>

      {/* Package Info */}
      <div className="bg-gray-50 rounded-lg p-4 mb-4">
        <h4 className="font-medium mb-2 flex items-center space-x-2">
          <Package className="h-4 w-4" />
          <span>Package Details</span>
        </h4>
        <div className="grid grid-cols-2 gap-4 text-sm">
          <div>
            <span className="text-gray-600">Weight:</span>
            <span className="ml-2 font-medium">{calculateWeight().toFixed(1)} kg</span>
          </div>
          <div>
            <span className="text-gray-600">Items:</span>
            <span className="ml-2 font-medium">{items.reduce((sum, item) => sum + item.quantity, 0)}</span>
          </div>
        </div>
      </div>

      {/* Destination Input */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          <MapPin className="h-4 w-4 inline mr-1" />
          Destination ZIP Code
        </label>
        <div className="flex space-x-2">
          <input
            type="text"
            value={zipCode}
            onChange={(e) => setZipCode(e.target.value)}
            placeholder="Enter ZIP code"
            maxLength={5}
            className="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            onClick={calculateShipping}
            disabled={isCalculating || zipCode.length < 5}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {isCalculating ? 'Calculating...' : 'Calculate'}
          </button>
        </div>
      </div>

      {/* Shipping Methods */}
      {shippingMethods.length > 0 && (
        <div className="space-y-3">
          <h4 className="font-medium flex items-center space-x-2">
            <Truck className="h-4 w-4" />
            <span>Available Shipping Options</span>
          </h4>

          {shippingMethods.map((method) => (
            <label
              key={method.id}
              className={`flex items-center p-3 border rounded-lg cursor-pointer transition-all ${
                selectedMethod?.id === method.id
                  ? 'border-blue-500 bg-blue-50'
                  : 'border-gray-200 hover:border-gray-300'
              }`}
            >
              <input
                type="radio"
                name="shippingMethod"
                value={method.id}
                checked={selectedMethod?.id === method.id}
                onChange={() => handleMethodSelect(method)}
                className="sr-only"
              />

              <div className="flex-1">
                <div className="flex justify-between items-start mb-1">
                  <h5 className="font-medium">{method.name}</h5>
                  <span className="font-bold text-lg">
                    {method.price === 0 ? 'Free' : `$${method.price.toFixed(2)}`}
                  </span>
                </div>
                <p className="text-sm text-gray-600 mb-1">{method.description}</p>
                <div className="flex items-center space-x-2 text-xs text-gray-500">
                  <span>{method.estimatedDays}</span>
                  <span>•</span>
                  <span>{method.courier}</span>
                  {method.supportsCOD && (
                    <>
                      <span>•</span>
                      <span className="text-green-600">COD Available</span>
                    </>
                  )}
                </div>
              </div>
            </label>
          ))}
        </div>
      )}

      {/* Free Shipping Threshold */}
      {items.length > 0 && (
        <div className="mt-4 p-3 bg-green-50 border border-green-200 rounded-lg">
          <div className="text-sm">
            {items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0) >= 50 ? (
              <p className="text-green-700 font-medium">
                🎉 You qualify for free standard shipping!
              </p>
            ) : (
              <p className="text-green-700">
                Add ${(50 - items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0)).toFixed(2)} more for free standard shipping
              </p>
            )}
          </div>
        </div>
      )}
    </div>
  );
};
import React from 'react';
import { Truck, Zap, Package, Clock, DollarSign } from 'lucide-react';
import { ShippingMethod } from '../types';

interface ShippingMethodSelectorProps {
  methods: ShippingMethod[];
  selectedMethod: ShippingMethod | null;
  onSelect: (method: ShippingMethod) => void;
  cartWeight?: number;
}

export const ShippingMethodSelector: React.FC<ShippingMethodSelectorProps> = ({
  methods,
  selectedMethod,
  onSelect,
  cartWeight = 1,
}) => {
  const getIcon = (iconName: string) => {
    switch (iconName) {
      case 'truck': return Truck;
      case 'zap': return Zap;
      case 'package': return Package;
      default: return Truck;
    }
  };

  return (
    <div className="space-y-3">
      <h3 className="text-lg font-semibold mb-4">Choose Shipping Method</h3>
      {methods.map((method) => {
        const Icon = getIcon(method.icon);
        const isSelected = selectedMethod?.id === method.id;

        return (
          <label
            key={method.id}
            className={`flex items-center p-4 border-2 rounded-lg cursor-pointer transition-all ${
              isSelected
                ? 'border-blue-500 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300 hover:bg-gray-50'
            }`}
          >
            <input
              type="radio"
              name="shippingMethod"
              value={method.id}
              checked={isSelected}
              onChange={() => onSelect(method)}
              className="sr-only"
            />

            <div className={`w-12 h-12 rounded-full flex items-center justify-center mr-4 ${
              isSelected ? 'bg-blue-100 text-blue-600' : 'bg-gray-100 text-gray-600'
            }`}>
              <Icon className="h-6 w-6" />
            </div>

            <div className="flex-1">
              <div className="flex justify-between items-start mb-1">
                <h4 className="font-semibold text-gray-900">{method.name}</h4>
                <span className="font-bold text-lg">
                  {method.price === 0 ? 'Free' : `$${method.price.toFixed(2)}`}
                </span>
              </div>

              <p className="text-gray-600 text-sm mb-2">{method.description}</p>

              <div className="flex items-center space-x-4 text-sm">
                <div className="flex items-center space-x-1 text-gray-500">
                  <Clock className="h-4 w-4" />
                  <span>{method.estimatedDays}</span>
                </div>
                <span className="text-gray-400">•</span>
                <span className="text-gray-500">{method.courier}</span>
                {method.supportsCOD && (
                  <>
                    <span className="text-gray-400">•</span>
                    <div className="flex items-center space-x-1 text-green-600">
                      <DollarSign className="h-3 w-3" />
                      <span className="text-xs">COD Available</span>
                    </div>
                  </>
                )}
              </div>

              {method.features.length > 0 && (
                <div className="mt-2 flex flex-wrap gap-1">
                  {method.features.map((feature, index) => (
                    <span
                      key={index}
                      className="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded-full"
                    >
                      {feature}
                    </span>
                  ))}
                </div>
              )}
            </div>
          </label>
        );
      })}
    </div>
  );
};
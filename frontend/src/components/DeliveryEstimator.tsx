import React, { useState, useEffect } from 'react';
import { Calendar, MapPin, Clock, Truck, AlertCircle } from 'lucide-react';
import { ShippingMethod } from '../types';

interface DeliveryEstimatorProps {
  shippingMethod: ShippingMethod;
  destinationZip: string;
  orderDate?: Date;
  onEstimateUpdate?: (estimate: DeliveryEstimate) => void;
}

interface DeliveryEstimate {
  earliestDate: Date;
  latestDate: Date;
  businessDays: number;
  isHolidayAffected: boolean;
  weatherWarning?: string;
}

export const DeliveryEstimator: React.FC<DeliveryEstimatorProps> = ({
  shippingMethod,
  destinationZip,
  orderDate = new Date(),
  onEstimateUpdate,
}) => {
  const [estimate, setEstimate] = useState<DeliveryEstimate | null>(null);
  const [isCalculating, setIsCalculating] = useState(false);

  const calculateDeliveryEstimate = async (): Promise<DeliveryEstimate> => {
    setIsCalculating(true);

    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 800));

    const businessDaysMatch = shippingMethod.estimatedDays.match(/(\d+)(?:-(\d+))?/);
    const minDays = businessDaysMatch ? parseInt(businessDaysMatch[1]) : 5;
    const maxDays = businessDaysMatch && businessDaysMatch[2] ? parseInt(businessDaysMatch[2]) : minDays;

    const startDate = new Date(orderDate);
    const earliestDate = addBusinessDays(startDate, minDays);
    const latestDate = addBusinessDays(startDate, maxDays);

    // Check for holidays and weather
    const isHolidayAffected = checkHolidayImpact(earliestDate, latestDate);
    const weatherWarning = checkWeatherImpact(destinationZip);

    const deliveryEstimate: DeliveryEstimate = {
      earliestDate,
      latestDate,
      businessDays: minDays,
      isHolidayAffected,
      weatherWarning,
    };

    setIsCalculating(false);
    return deliveryEstimate;
  };

  const addBusinessDays = (startDate: Date, businessDays: number): Date => {
    const result = new Date(startDate);
    let addedDays = 0;

    while (addedDays < businessDays) {
      result.setDate(result.getDate() + 1);
      // Skip weekends
      if (result.getDay() !== 0 && result.getDay() !== 6) {
        addedDays++;
      }
    }

    return result;
  };

  const checkHolidayImpact = (start: Date, end: Date): boolean => {
    const holidays = getUpcomingHolidays();
    return holidays.some(holiday => holiday >= start && holiday <= end);
  };

  const getUpcomingHolidays = (): Date[] => {
    const currentYear = new Date().getFullYear();
    return [
      new Date(currentYear, 0, 1),   // New Year's Day
      new Date(currentYear, 6, 4),   // Independence Day
      new Date(currentYear, 10, 11), // Veterans Day
      new Date(currentYear, 11, 25), // Christmas
      // Add more holidays as needed
    ];
  };

  const checkWeatherImpact = (zipCode: string): string | undefined => {
    // Mock weather impact based on zip code and season
    const firstDigit = parseInt(zipCode.charAt(0));
    const currentMonth = new Date().getMonth();

    // Winter months (Dec, Jan, Feb) in northern regions
    if ((currentMonth === 11 || currentMonth <= 1) && firstDigit <= 5) {
      return 'Winter weather may cause delays in northern regions';
    }

    // Hurricane season in southeastern regions
    if (currentMonth >= 5 && currentMonth <= 10 && firstDigit === 3) {
      return 'Hurricane season may affect deliveries in southeastern regions';
    }

    return undefined;
  };

  const formatDateRange = (start: Date, end: Date): string => {
    const options: Intl.DateTimeFormatOptions = {
      weekday: 'short',
      month: 'short',
      day: 'numeric',
    };

    if (start.toDateString() === end.toDateString()) {
      return start.toLocaleDateString('en-US', options);
    }

    return `${start.toLocaleDateString('en-US', options)} - ${end.toLocaleDateString('en-US', options)}`;
  };

  const getDaysUntilDelivery = (date: Date): number => {
    const today = new Date();
    const diffTime = date.getTime() - today.getTime();
    return Math.ceil(diffTime / (1000 * 60 * 60 * 24));
  };

  useEffect(() => {
    if (shippingMethod && destinationZip) {
      calculateDeliveryEstimate().then(newEstimate => {
        setEstimate(newEstimate);
        if (onEstimateUpdate) {
          onEstimateUpdate(newEstimate);
        }
      });
    }
  }, [shippingMethod, destinationZip, orderDate]);

  if (isCalculating) {
    return (
      <div className="bg-white rounded-lg border border-gray-200 p-6">
        <div className="flex items-center space-x-3">
          <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
          <span className="text-gray-600">Calculating delivery estimate...</span>
        </div>
      </div>
    );
  }

  if (!estimate) {
    return null;
  }

  return (
    <div className="bg-white rounded-lg border border-gray-200 p-6">
      <div className="flex items-center space-x-2 mb-4">
        <Calendar className="h-5 w-5 text-blue-600" />
        <h3 className="text-lg font-semibold">Delivery Estimate</h3>
      </div>

      {/* Main Estimate */}
      <div className="bg-gradient-to-r from-blue-50 to-green-50 rounded-lg p-4 mb-4">
        <div className="flex items-center justify-between">
          <div>
            <p className="text-sm text-gray-600 mb-1">Estimated Delivery</p>
            <p className="text-xl font-bold text-gray-900">
              {formatDateRange(estimate.earliestDate, estimate.latestDate)}
            </p>
            <p className="text-sm text-gray-600">
              {getDaysUntilDelivery(estimate.earliestDate)} days from now
            </p>
          </div>
          <div className="text-4xl">
            {shippingMethod.icon === 'zap' ? '⚡' : shippingMethod.icon === 'package' ? '📦' : '🚚'}
          </div>
        </div>
      </div>

      {/* Delivery Details */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
        <div className="flex items-center space-x-3">
          <Truck className="h-5 w-5 text-gray-400" />
          <div>
            <p className="text-sm text-gray-600">Shipping Method</p>
            <p className="font-medium">{shippingMethod.name}</p>
          </div>
        </div>

        <div className="flex items-center space-x-3">
          <Clock className="h-5 w-5 text-gray-400" />
          <div>
            <p className="text-sm text-gray-600">Business Days</p>
            <p className="font-medium">{estimate.businessDays} days</p>
          </div>
        </div>

        <div className="flex items-center space-x-3">
          <MapPin className="h-5 w-5 text-gray-400" />
          <div>
            <p className="text-sm text-gray-600">Destination</p>
            <p className="font-medium">{destinationZip}</p>
          </div>
        </div>

        <div className="flex items-center space-x-3">
          <Calendar className="h-5 w-5 text-gray-400" />
          <div>
            <p className="text-sm text-gray-600">Order Date</p>
            <p className="font-medium">{orderDate.toLocaleDateString()}</p>
          </div>
        </div>
      </div>

      {/* Warnings and Notes */}
      {(estimate.isHolidayAffected || estimate.weatherWarning) && (
        <div className="space-y-2">
          {estimate.isHolidayAffected && (
            <div className="flex items-start space-x-2 p-3 bg-yellow-50 border border-yellow-200 rounded-lg">
              <AlertCircle className="h-4 w-4 text-yellow-600 mt-0.5 flex-shrink-0" />
              <div className="text-sm">
                <p className="font-medium text-yellow-800">Holiday Impact</p>
                <p className="text-yellow-700">
                  Delivery may be delayed due to holidays during the estimated delivery period.
                </p>
              </div>
            </div>
          )}

          {estimate.weatherWarning && (
            <div className="flex items-start space-x-2 p-3 bg-blue-50 border border-blue-200 rounded-lg">
              <AlertCircle className="h-4 w-4 text-blue-600 mt-0.5 flex-shrink-0" />
              <div className="text-sm">
                <p className="font-medium text-blue-800">Weather Advisory</p>
                <p className="text-blue-700">{estimate.weatherWarning}</p>
              </div>
            </div>
          )}
        </div>
      )}

      {/* Delivery Timeline */}
      <div className="mt-4 pt-4 border-t border-gray-200">
        <h4 className="font-medium mb-3">Delivery Timeline</h4>
        <div className="space-y-2">
          <div className="flex items-center space-x-3">
            <div className="w-3 h-3 bg-green-500 rounded-full"></div>
            <span className="text-sm text-gray-600">
              Order processed: {orderDate.toLocaleDateString()}
            </span>
          </div>
          <div className="flex items-center space-x-3">
            <div className="w-3 h-3 bg-blue-500 rounded-full"></div>
            <span className="text-sm text-gray-600">
              Shipped: Within 1-2 business days
            </span>
          </div>
          <div className="flex items-center space-x-3">
            <div className="w-3 h-3 bg-purple-500 rounded-full"></div>
            <span className="text-sm text-gray-600">
              Delivered: {formatDateRange(estimate.earliestDate, estimate.latestDate)}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
};
import { ShippingMethod, CartItem, Address } from '../types';

export interface ShippingCalculation {
methods: ShippingMethod[];
freeShippingThreshold: number;
currentSubtotal: number;
amountForFreeShipping: number;
}

export const calculateShippingOptions = (
items: CartItem[],
destination: Address
): ShippingCalculation => {
const subtotal = items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0);
const weight = calculateTotalWeight(items);
const dimensions = calculatePackageDimensions(items);
const distanceMultiplier = getDistanceMultiplier(destination.zipCode);
const sizeMultiplier = getSizeMultiplier(dimensions);

const freeShippingThreshold = 50;
const amountForFreeShipping = Math.max(0, freeShippingThreshold - subtotal);

const baseRate = 5.99;
const weightRate = weight * 2.50;
const calculatedRate = (baseRate + weightRate) * distanceMultiplier * sizeMultiplier;

const methods: ShippingMethod[] = [
{
id: 'standard',
name: 'Standard Shipping',
description: 'Reliable delivery with tracking',
price: subtotal >= freeShippingThreshold ? 0 : Math.round(calculatedRate * 100) / 100,
estimatedDays: getEstimatedDays('standard', destination.zipCode),
      courier: 'USPS',
      icon: 'truck',
      features: ['Tracking included', 'Insurance up to $100', 'Signature on delivery'],
      supportsCOD: true
    },
    {
      id: 'express',
      name: 'Express Shipping',
      description: 'Faster delivery for urgent orders',
      price: Math.round((calculatedRate + 15) * 100) / 100,
      estimatedDays: getEstimatedDays('express', destination.zipCode),
      courier: 'FedEx',
      icon: 'zap',
      features: ['Priority handling', 'Tracking included', 'Insurance up to $500', 'Signature required'],
      supportsCOD: true
    },
    {
      id: 'overnight',
      name: 'Overnight Delivery',
      description: 'Next business day delivery',
      price: Math.round((calculatedRate + 35) * 100) / 100,
      estimatedDays: '1 business day',
      courier: 'UPS',
      icon: 'package',
      features: ['Next day delivery', 'Signature required', 'Insurance up to $1000', 'Morning delivery'],
      supportsCOD: false
    },
    {
      id: 'same_day',
      name: 'Same Day Delivery',
      description: 'Delivery within hours (select areas)',
      price: Math.round((calculatedRate + 25) * 100) / 100,
      estimatedDays: '4-8 hours',
      courier: 'Local Courier',
      icon: 'zap',
      features: ['Same day delivery', 'Real-time tracking', 'Direct contact with driver'],
      supportsCOD: true
    }
  ];

  // Filter same day delivery based on location
  const filteredMethods = isSameDayAvailable(destination.zipCode)
    ? methods
    : methods.filter(m => m.id !== 'same_day');

  return {
    methods: filteredMethods,
    freeShippingThreshold,
    currentSubtotal: subtotal,
    amountForFreeShipping
  };
};

export const calculateTotalWeight = (items: CartItem[]): number => {
  return items.reduce((total, item) => {
    const weight = item.product.weight || 0.5; // Default 0.5kg if not specified
    return total + (weight * item.quantity);
  }, 0);
};

export const calculatePackageDimensions = (items: CartItem[]) => {
  // Simplified calculation - in reality, this would be more complex
  const totalVolume = items.reduce((total, item) => {
    const dims = item.product.dimensions || { length: 20, width: 15, height: 10 };
    const volume = dims.length * dims.width * dims.height * item.quantity;
    return total + volume;
  }, 0);

  // Convert to approximate box dimensions (in cm)
  const side = Math.cbrt(totalVolume);
  return {
    length: Math.ceil(side),
    width: Math.ceil(side),
    height: Math.ceil(side * 0.8),
  };
};

export const getDistanceMultiplier = (zipCode: string): number => {
  if (!zipCode || zipCode.length < 5) return 1.0;

  const firstDigit = parseInt(zipCode.charAt(0));

  // Distance-based multipliers
  if (firstDigit <= 2) return 1.0; // East Coast
  if (firstDigit <= 5) return 1.2; // Central
  if (firstDigit <= 7) return 1.4; // Mountain
  return 1.6; // West Coast
};

export const getSizeMultiplier = (dimensions: { length: number; width: number; height: number }): number => {
  const volume = dimensions.length * dimensions.width * dimensions.height;

  if (volume < 1000) return 1.0;    // Small package
  if (volume < 5000) return 1.2;    // Medium package
  if (volume < 10000) return 1.5;   // Large package
  return 2.0;                       // Oversized package
};

export const getEstimatedDays = (method: string, zipCode: string): string => {
  const distanceMultiplier = getDistanceMultiplier(zipCode);

  switch (method) {
    case 'standard':
      if (distanceMultiplier <= 1.0) return '2-4 business days';
      if (distanceMultiplier <= 1.2) return '3-5 business days';
      if (distanceMultiplier <= 1.4) return '4-6 business days';
      return '5-7 business days';

    case 'express':
      if (distanceMultiplier <= 1.0) return '1-2 business days';
      if (distanceMultiplier <= 1.2) return '2-3 business days';
      return '2-4 business days';

    default:
      return '5-7 business days';
  }
};

export const isSameDayAvailable = (zipCode: string): boolean => {
  // Mock same-day availability for major metropolitan areas
  const sameDayZips = [
    '10001', '10002', '10003', // NYC
    '90210', '90211', '90212', // LA
    '60601', '60602', '60603', // Chicago
    '94102', '94103', '94104', // San Francisco
    '02101', '02102', '02103', // Boston
  ];

  return sameDayZips.some(zip => zipCode.startsWith(zip.substring(0, 3)));
};

export const formatShippingCost = (cost: number): string => {
  return cost === 0 ? 'Free' : `$${cost.toFixed(2)}`;
};

export const calculateDeliveryDate = (
  shippingMethod: ShippingMethod,
  orderDate: Date = new Date()
): { earliest: Date; latest: Date } => {
  const businessDaysMatch = shippingMethod.estimatedDays.match(/(\d+)(?:-(\d+))?/);
  const minDays = businessDaysMatch ? parseInt(businessDaysMatch[1]) : 5;
  const maxDays = businessDaysMatch && businessDaysMatch[2] ? parseInt(businessDaysMatch[2]) : minDays;

  const earliest = addBusinessDays(orderDate, minDays);
  const latest = addBusinessDays(orderDate, maxDays);

  return { earliest, latest };
};

export const addBusinessDays = (startDate: Date, businessDays: number): Date => {
  const result = new Date(startDate);
  let addedDays = 0;

  while (addedDays < businessDays) {
    result.setDate(result.getDate() + 1);
    // Skip weekends (0 = Sunday, 6 = Saturday)
    if (result.getDay() !== 0 && result.getDay() !== 6) {
      addedDays++;
    }
  }

  return result;
};

export const isHolidayPeriod = (date: Date): boolean => {
  const month = date.getMonth();
  const day = date.getDate();

  // Check for major holidays that affect shipping
  const holidays = [
    { month: 0, day: 1 },   // New Year's Day
    { month: 6, day: 4 },   // Independence Day
    { month: 10, day: 11 }, // Veterans Day
    { month: 11, day: 25 }, // Christmas Day
  ];

  return holidays.some(holiday => holiday.month === month && holiday.day === day);
};

export const getShippingInsurance = (orderValue: number): number => {
  if (orderValue <= 100) return 0;
  if (orderValue <= 500) return 2.99;
  if (orderValue <= 1000) return 4.99;
  return Math.ceil(orderValue * 0.005); // 0.5% of order value
};

export const validateShippingAddress = (address: Address): string[] => {
  const errors: string[] = [];

  if (!address.firstName?.trim()) errors.push('First name is required');
  if (!address.lastName?.trim()) errors.push('Last name is required');
  if (!address.address1?.trim()) errors.push('Address line 1 is required');
  if (!address.city?.trim()) errors.push('City is required');
  if (!address.state?.trim()) errors.push('State is required');
  if (!address.zipCode?.trim()) errors.push('ZIP code is required');
  if (address.zipCode && !/^\d{5}(-\d{4})?$/.test(address.zipCode)) {
    errors.push('Invalid ZIP code format');
  }

  return errors;
};
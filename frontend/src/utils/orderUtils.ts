import { Order, TrackingEvent, ShippingMethod } from '../types';

export const generateTrackingNumber = (): string => {
const prefix = 'TRK';
const numbers = Math.random().toString().slice(2, 12);
return `${prefix}${numbers}`;
};

export const generateOrderNumber = (): string => {
return `ORD-${Date.now().toString().slice(-8)}`;
};

export const createMockTrackingEvents = (status: string): TrackingEvent[] => {
  const events: TrackingEvent[] = [
    {
      id: '1',
      status: 'confirmed',
      description: 'Order confirmed and payment processed',
      location: 'Processing Center',
      timestamp: new Date(Date.now() - 4 * 24 * 60 * 60 * 1000).toISOString(),
      isCompleted: true,
    },
    {
      id: '2',
      status: 'processing',
      description: 'Items are being prepared for shipment',
      location: 'Fulfillment Center',
      timestamp: new Date(Date.now() - 3 * 24 * 60 * 60 * 1000).toISOString(),
      isCompleted: true,
    },
  ];

  if (['shipped', 'out_for_delivery', 'delivered'].includes(status)) {
    events.push({
      id: '3',
      status: 'shipped',
      description: 'Package has been shipped and is in transit',
      location: 'Distribution Center',
      timestamp: new Date(Date.now() - 2 * 24 * 60 * 60 * 1000).toISOString(),
      isCompleted: true,
    });
  }

  if (['out_for_delivery', 'delivered'].includes(status)) {
    events.push({
      id: '4',
      status: 'out_for_delivery',
      description: 'Package is out for delivery',
      location: 'Local Delivery Hub',
      timestamp: new Date(Date.now() - 1 * 24 * 60 * 60 * 1000).toISOString(),
      isCompleted: true,
    });
  }

  if (status === 'delivered') {
    events.push({
      id: '5',
      status: 'delivered',
      description: 'Package has been delivered successfully',
      location: 'Customer Address',
      timestamp: new Date().toISOString(),
      isCompleted: true,
    });
  }

  return events;
};

export const calculateEstimatedDelivery = (shippingMethod: ShippingMethod): string => {
  const today = new Date();
  let deliveryDate = new Date(today);

  // Parse estimated days from shipping method
  const daysMatch = shippingMethod.estimatedDays.match(/(\d+)/);
  const maxDays = daysMatch ? parseInt(daysMatch[0]) : 7;

  // Add business days (skip weekends)
  let addedDays = 0;
  while (addedDays < maxDays) {
    deliveryDate.setDate(deliveryDate.getDate() + 1);
    // Skip weekends
    if (deliveryDate.getDay() !== 0 && deliveryDate.getDay() !== 6) {
      addedDays++;
    }
  }

  return deliveryDate.toISOString();
};

export const createMockOrders = (userId: string): Order[] => {
  const mockShippingMethod: ShippingMethod = {
    id: 'standard',
    name: 'Standard Shipping',
    description: 'Reliable delivery with tracking',
    price: 9.99,
    estimatedDays: '5-7 business days',
    courier: 'USPS',
    icon: 'truck',
    features: ['Tracking included', 'Insurance up to $100']
  };

  const mockAddress = {
    id: '1',
    type: 'shipping' as const,
    firstName: 'John',
    lastName: 'Doe',
    address1: '123 Main Street',
    city: 'New York',
    state: 'NY',
    zipCode: '10001',
    country: 'US',
    isDefault: true,
  };

  const mockPaymentMethod = {
    id: '1',
    type: 'card' as const,
    last4: '4242',
    brand: 'visa',
    expiryMonth: 12,
    expiryYear: 2025,
    isDefault: true,
  };

  return [
    {
      id: '1',
      orderNumber: 'ORD-12345678',
      buyerId: userId,
      sellerId: 'seller1',
      products: [],
      total: 124.98,
      subtotal: 99.99,
      tax: 8.00,
      shipping: 16.99,
      status: 'shipped',
      orderDate: new Date(Date.now() - 3 * 24 * 60 * 60 * 1000).toISOString(),
      shippingAddress: mockAddress,
      billingAddress: mockAddress,
      paymentMethod: mockPaymentMethod,
      shippingMethod: mockShippingMethod,
      trackingNumber: 'TRK1234567890',
      estimatedDelivery: new Date(Date.now() + 2 * 24 * 60 * 60 * 1000).toISOString(),
      trackingEvents: createMockTrackingEvents('shipped'),
    },
    {
      id: '2',
      orderNumber: 'ORD-87654321',
      buyerId: userId,
      sellerId: 'seller2',
      products: [],
      total: 299.97,
      subtotal: 249.99,
      tax: 20.00,
      shipping: 29.98,
      status: 'delivered',
      orderDate: new Date(Date.now() - 10 * 24 * 60 * 60 * 1000).toISOString(),
      shippingAddress: mockAddress,
      billingAddress: mockAddress,
      paymentMethod: mockPaymentMethod,
      shippingMethod: mockShippingMethod,
      trackingNumber: 'TRK0987654321',
      actualDelivery: new Date(Date.now() - 2 * 24 * 60 * 60 * 1000).toISOString(),
      trackingEvents: createMockTrackingEvents('delivered'),
    },
  ];
};
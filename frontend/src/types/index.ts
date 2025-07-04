export interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  category: string;
  image: string;
  images?: string[];
  sellerId: string;
  sellerName: string;
  rating: number;
  reviews: Review[];
  stock: number;
  tags: string[];
  weight?: number; // in kg for shipping calculations
  dimensions?: {
    length: number;
    width: number;
    height: number;
  };
}

export interface Review {
  id: string;
  userId: string;
  userName: string;
  rating: number;
  comment: string;
  date: string;
}

export interface CartItem {
  product: Product;
  quantity: number;
}

export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
  isSeller: boolean;
  sellerProfile?: SellerProfile;
  addresses?: Address[];
  paymentMethods?: PaymentMethod[];
  orders?: Order[];
}

export interface SellerProfile {
  businessName: string;
  description: string;
  location: string;
  phone: string;
  website?: string;
  logo?: string;
  rating: number;
  totalSales: number;
  joinDate: string;
  verified: boolean;
  socialMedia?: {
    facebook?: string;
    twitter?: string;
    instagram?: string;
  };
}

export interface ShippingMethod {
  id: string;
  name: string;
  description: string;
  price: number;
  estimatedDays: string;
  courier: string;
  icon: string;
  features: string[];
  supportsCOD?: boolean; // Cash on Delivery support
}

export interface TrackingEvent {
  id: string;
  status: string;
  description: string;
  location: string;
  timestamp: string;
  isCompleted: boolean;
}

export interface Order {
  id: string;
  orderNumber: string;
  buyerId: string;
  sellerId: string;
  products: CartItem[];
  total: number;
  subtotal: number;
  tax: number;
  shipping: number
  codFee?: number; // Cash on Delivery fee
  status: 'pending' | 'confirmed' | 'processing' | 'shipped' | 'out_for_delivery' | 'delivered' | 'cancelled' | 'returned';
  orderDate: string;
  shippingAddress: Address;
  billingAddress: Address;
  paymentMethod: PaymentMethod;
  shippingMethod: ShippingMethod;
  trackingNumber?: string;
  estimatedDelivery?: string;
  actualDelivery?: string;
  trackingEvents: TrackingEvent[];
  notes?: string;
}

export interface Address {
  id: string;
  type: 'shipping' | 'billing';
  firstName: string;
  lastName: string;
  company?: string;
  address1: string;
  address2?: string;
  city: string;
  state: string;
  zipCode: string;
  country: string;
  phone?: string;
  isDefault: boolean;
}

export interface PaymentMethod {
  id: string;
  type: 'card' | 'paypal' | 'bank' | 'cod'; // Added COD type;
  last4?: string;
  brand?: string;
  expiryMonth?: number;
  expiryYear?: number;
  isDefault: boolean;
  billingAddress?: Address;
}

export interface CheckoutData {
  items: CartItem[];
  shippingAddress: Address;
  billingAddress: Address;
  paymentMethod: PaymentMethod;
  shippingMethod: ShippingMethod;
  subtotal: number;
  tax: number;
  shipping: number;
  codFee?: number; // Cash on Delivery fee
  total: number;
}
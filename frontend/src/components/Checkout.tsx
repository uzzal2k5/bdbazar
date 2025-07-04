import React, { useState } from 'react';
import { X, CreditCard, MapPin, Truck, Lock, ArrowLeft, ArrowRight } from 'lucide-react';
import { CartItem, Address, PaymentMethod, CheckoutData, ShippingMethod } from '../types';
import { AddressForm } from './AddressForm';
import { PaymentForm } from './PaymentForm';
import { ShippingMethodSelector } from './ShippingMethodSelector';

interface CheckoutProps {
  items: CartItem[];
  onClose: () => void;
  onOrderComplete: (orderData: CheckoutData) => void;
}

export const Checkout: React.FC<CheckoutProps> = ({
  items,
  onClose,
  onOrderComplete,
}) => {
  const [currentStep, setCurrentStep] = useState<'shipping' | 'method' | 'payment' | 'review'>('shipping');
  const [shippingAddress, setShippingAddress] = useState<Address | null>(null);
  const [billingAddress, setBillingAddress] = useState<Address | null>(null);
  const [paymentMethod, setPaymentMethod] = useState<PaymentMethod | null>(null);
  const [shippingMethod, setShippingMethod] = useState<ShippingMethod | null>(null);
  const [useSameAddress, setUseSameAddress] = useState(true);

  const subtotal = items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0);
  const tax = subtotal * 0.08; // 8% tax
  const shippingCost = shippingMethod?.price || 0;

  // Calculate COD fee if COD is selected
  const codFee = paymentMethod?.type === 'cod' ? Math.min(Math.max(subtotal * 0.02, 2), 10) : 0;

  const total = subtotal + tax + shippingCost + codFee;

  // Mock shipping methods with COD support
  const shippingMethods: ShippingMethod[] = [
    {
      id: 'standard',
      name: 'Standard Shipping',
      description: 'Reliable delivery with tracking',
      price: subtotal > 50 ? 0 : 9.99,
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
      price: 19.99,
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
      price: 39.99,
      estimatedDays: '1 business day',
      courier: 'UPS',
      icon: 'package',
      features: ['Next day delivery', 'Signature required', 'Insurance up to $1000'],
      supportsCOD: false // Overnight doesn't support COD
    }
  ];

  const handleShippingSubmit = (address: Address) => {
    setShippingAddress(address);
    if (useSameAddress) {
      setBillingAddress(address);
    }
    setCurrentStep('method');
  };

  const handleShippingMethodSelect = (method: ShippingMethod) => {
    setShippingMethod(method);
    // Reset payment method if COD was selected but new shipping method doesn't support it
    if (paymentMethod?.type === 'cod' && !method.supportsCOD) {
      setPaymentMethod(null);
    }
  };

  const handlePaymentSubmit = (payment: PaymentMethod, billing?: Address) => {
    setPaymentMethod(payment);
    if (billing && !useSameAddress) {
      setBillingAddress(billing);
    }
    setCurrentStep('review');
  };

  const handleOrderSubmit = () => {
    if (shippingAddress && paymentMethod && billingAddress && shippingMethod) {
      const orderData: CheckoutData = {
        items,
        shippingAddress,
        billingAddress,
        paymentMethod,
        shippingMethod,
        subtotal,
        tax,
        shipping: shippingCost,
        codFee,
        total,
      };
      onOrderComplete(orderData);
    }
  };

  const steps = [
    { id: 'shipping', label: 'Shipping', icon: MapPin },
    { id: 'method', label: 'Delivery', icon: Truck },
    { id: 'payment', label: 'Payment', icon: CreditCard },
    { id: 'review', label: 'Review', icon: Lock },
  ];

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-2xl max-w-6xl w-full max-h-[95vh] overflow-y-auto">
        <div className="sticky top-0 bg-white border-b border-gray-200 p-4 flex justify-between items-center">
          <h2 className="text-xl font-semibold">Secure Checkout</h2>
          <button
            onClick={onClose}
            className="p-2 hover:bg-gray-100 rounded-full transition-colors"
          >
            <X className="h-6 w-6" />
          </button>
        </div>

        <div className="p-6">
          {/* Progress Steps */}
          <div className="mb-8">
            <div className="flex items-center justify-center space-x-8">
              {steps.map((step, index) => {
                const Icon = step.icon;
                const isActive = step.id === currentStep;
                const isCompleted = steps.findIndex(s => s.id === currentStep) > index;

                return (
                  <div key={step.id} className="flex items-center">
                    <div className={`flex items-center space-x-2 ${
                      isActive ? 'text-blue-600' : isCompleted ? 'text-green-600' : 'text-gray-400'
                    }`}>
                      <div className={`w-10 h-10 rounded-full flex items-center justify-center ${
                        isActive ? 'bg-blue-100' : isCompleted ? 'bg-green-100' : 'bg-gray-100'
                      }`}>
                        <Icon className="h-5 w-5" />
                      </div>
                      <span className="font-medium">{step.label}</span>
                    </div>
                    {index < steps.length - 1 && (
                      <div className={`w-16 h-0.5 mx-4 ${
                        isCompleted ? 'bg-green-600' : 'bg-gray-300'
                      }`} />
                    )}
                  </div>
                );
              })}
            </div>
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            {/* Main Content */}
            <div className="lg:col-span-2">
              {currentStep === 'shipping' && (
                <div>
                  <h3 className="text-lg font-semibold mb-4">Shipping Address</h3>
                  <AddressForm
                    type="shipping"
                    onSubmit={handleShippingSubmit}
                    initialData={shippingAddress}
                  />

                  <div className="mt-6">
                    <label className="flex items-center space-x-2">
                      <input
                        type="checkbox"
                        checked={useSameAddress}
                        onChange={(e) => setUseSameAddress(e.target.checked)}
                        className="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                      />
                      <span className="text-sm text-gray-700">
                        Use same address for billing
                      </span>
                    </label>
                  </div>
                </div>
              )}

              {currentStep === 'method' && (
                <div>
                  <ShippingMethodSelector
                    methods={shippingMethods}
                    selectedMethod={shippingMethod}
                    onSelect={handleShippingMethodSelect}
                  />

                  <div className="mt-6 flex justify-between">
                    <button
                      onClick={() => setCurrentStep('shipping')}
                      className="flex items-center space-x-2 px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
                    >
                      <ArrowLeft className="h-4 w-4" />
                      <span>Back to Shipping</span>
                    </button>

                    <button
                      onClick={() => setCurrentStep('payment')}
                      disabled={!shippingMethod}
                      className="flex items-center space-x-2 bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                      <span>Continue to Payment</span>
                      <ArrowRight className="h-4 w-4" />
                    </button>
                  </div>
                </div>
              )}

              {currentStep === 'payment' && (
                <div>
                  <h3 className="text-lg font-semibold mb-4">Payment Information</h3>
                  <PaymentForm
                    onSubmit={handlePaymentSubmit}
                    requiresBilling={!useSameAddress}
                    initialBilling={billingAddress}
                    orderTotal={subtotal}
                    shippingMethodSupportsCOD={shippingMethod?.supportsCOD}
                  />

                  <div className="mt-6 flex justify-between">
                    <button
                      onClick={() => setCurrentStep('method')}
                      className="flex items-center space-x-2 px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
                    >
                      <ArrowLeft className="h-4 w-4" />
                      <span>Back to Shipping Method</span>
                    </button>
                  </div>
                </div>
              )}

              {currentStep === 'review' && (
                <div>
                  <h3 className="text-lg font-semibold mb-6">Review Your Order</h3>

                  {/* Order Items */}
                  <div className="bg-gray-50 rounded-lg p-4 mb-6">
                    <h4 className="font-medium mb-4">Order Items</h4>
                    <div className="space-y-3">
                      {items.map((item) => (
                        <div key={item.product.id} className="flex items-center space-x-4">
                          <img
                            src={item.product.image}
                            alt={item.product.name}
                            className="w-12 h-12 object-cover rounded"
                          />
                          <div className="flex-1">
                            <p className="font-medium">{item.product.name}</p>
                            <p className="text-sm text-gray-600">Qty: {item.quantity}</p>
                          </div>
                          <p className="font-medium">${(item.product.price * item.quantity).toFixed(2)}</p>
                        </div>
                      ))}
                    </div>
                  </div>

                  {/* Addresses and Shipping */}
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
                    <div className="bg-gray-50 rounded-lg p-4">
                      <h4 className="font-medium mb-2 flex items-center space-x-2">
                        <MapPin className="h-4 w-4" />
                        <span>Shipping Address</span>
                      </h4>
                      {shippingAddress && (
                        <div className="text-sm text-gray-600">
                          <p>{shippingAddress.firstName} {shippingAddress.lastName}</p>
                          <p>{shippingAddress.address1}</p>
                          {shippingAddress.address2 && <p>{shippingAddress.address2}</p>}
                          <p>{shippingAddress.city}, {shippingAddress.state} {shippingAddress.zipCode}</p>
                        </div>
                      )}
                    </div>

                    <div className="bg-gray-50 rounded-lg p-4">
                      <h4 className="font-medium mb-2 flex items-center space-x-2">
                        <Truck className="h-4 w-4" />
                        <span>Shipping Method</span>
                      </h4>
                      {shippingMethod && (
                        <div className="text-sm text-gray-600">
                          <p className="font-medium">{shippingMethod.name}</p>
                          <p>{shippingMethod.estimatedDays}</p>
                          <p>{shippingMethod.courier}</p>
                        </div>
                      )}
                    </div>
                  </div>

                  <div className="bg-gray-50 rounded-lg p-4 mb-6">
                    <h4 className="font-medium mb-2 flex items-center space-x-2">
                      <CreditCard className="h-4 w-4" />
                      <span>Payment Method</span>
                    </h4>
                    {paymentMethod && (
                      <div className="text-sm text-gray-600">
                        {paymentMethod.type === 'card' ? (
                          <>
                            <p className="capitalize">{paymentMethod.brand} ending in {paymentMethod.last4}</p>
                            <p>Expires {paymentMethod.expiryMonth}/{paymentMethod.expiryYear}</p>
                          </>
                        ) : paymentMethod.type === 'paypal' ? (
                          <p>PayPal</p>
                        ) : (
                          <div>
                            <p className="font-medium text-green-700">Cash on Delivery</p>
                            <p>Pay when you receive your order</p>
                            <p className="text-xs text-green-600">COD Fee: ${codFee.toFixed(2)}</p>
                          </div>
                        )}
                      </div>
                    )}
                  </div>

                  <div className="flex justify-between">
                    <button
                      onClick={() => setCurrentStep('payment')}
                      className="flex items-center space-x-2 px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
                    >
                      <ArrowLeft className="h-4 w-4" />
                      <span>Back to Payment</span>
                    </button>

                    <button
                      onClick={handleOrderSubmit}
                      className="flex items-center space-x-2 bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors font-medium"
                    >
                      <Lock className="h-4 w-4" />
                      <span>
                        {paymentMethod?.type === 'cod' ? 'Place COD Order' : 'Place Order'} - ${total.toFixed(2)}
                      </span>
                    </button>
                  </div>
                </div>
              )}
            </div>

            {/* Order Summary */}
            <div className="lg:col-span-1">
              <div className="bg-gray-50 rounded-lg p-6 sticky top-6">
                <h3 className="font-semibold mb-4">Order Summary</h3>

                <div className="space-y-3 mb-4">
                  <div className="flex justify-between">
                    <span>Subtotal</span>
                    <span>${subtotal.toFixed(2)}</span>
                  </div>
                  <div className="flex justify-between">
                    <span>Shipping</span>
                    <span>
                      {shippingMethod
                        ? (shippingMethod.price === 0 ? 'Free' : `$${shippingMethod.price.toFixed(2)}`)
                        : 'TBD'
                      }
                    </span>
                  </div>
                  <div className="flex justify-between">
                    <span>Tax</span>
                    <span>${tax.toFixed(2)}</span>
                  </div>
                  {codFee > 0 && (
                    <div className="flex justify-between text-green-700">
                      <span>COD Fee</span>
                      <span>${codFee.toFixed(2)}</span>
                    </div>
                  )}
                  <div className="border-t border-gray-300 pt-3">
                    <div className="flex justify-between font-semibold text-lg">
                      <span>Total</span>
                      <span>${total.toFixed(2)}</span>
                    </div>
                  </div>
                </div>

                <div className="text-xs text-gray-500 space-y-1">
                  <p className="flex items-center space-x-1">
                    <Lock className="h-3 w-3" />
                    <span>Secure SSL encryption</span>
                  </p>
                  <p>30-day return policy</p>
                  {shippingMethod && (
                    <p>Estimated delivery: {shippingMethod.estimatedDays}</p>
                  )}
                  {paymentMethod?.type === 'cod' && (
                    <p className="text-green-600">✓ Cash on Delivery available</p>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
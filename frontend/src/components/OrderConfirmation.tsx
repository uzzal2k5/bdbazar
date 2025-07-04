import React from 'react';
import { CheckCircle, Package, Truck, CreditCard, Download, DollarSign } from 'lucide-react';
import { CheckoutData } from '../types';

interface OrderConfirmationProps {
  orderData: CheckoutData;
  orderNumber: string;
  onClose: () => void;
}

export const OrderConfirmation: React.FC<OrderConfirmationProps> = ({
  orderData,
  orderNumber,
  onClose,
}) => {
  const estimatedDelivery = new Date();
  estimatedDelivery.setDate(estimatedDelivery.getDate() + 5);

  const isCOD = orderData.paymentMethod.type === 'cod';

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-2xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div className="p-8 text-center">
          <div className="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
            <CheckCircle className="h-8 w-8 text-green-600" />
          </div>

          <h2 className="text-2xl font-bold text-gray-900 mb-2">Order Confirmed!</h2>
          <p className="text-gray-600 mb-6">
            {isCOD
              ? "Your Cash on Delivery order has been successfully placed."
              : "Thank you for your purchase. Your order has been successfully placed."
            }
          </p>

          <div className="bg-gray-50 rounded-lg p-4 mb-6">
            <div className="flex justify-between items-center mb-2">
              <span className="text-sm text-gray-600">Order Number</span>
              <span className="font-mono font-medium">{orderNumber}</span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-sm text-gray-600">Total Amount</span>
              <span className="font-bold text-lg">${orderData.total.toFixed(2)}</span>
            </div>
            {isCOD && orderData.codFee && (
              <div className="flex justify-between items-center mt-2 text-green-700">
                <span className="text-sm">COD Fee (included)</span>
                <span className="text-sm font-medium">${orderData.codFee.toFixed(2)}</span>
              </div>
            )}
          </div>

          {/* COD Specific Information */}
          {isCOD && (
            <div className="bg-green-50 border border-green-200 rounded-lg p-4 mb-6">
              <div className="flex items-center justify-center space-x-2 mb-3">
                <DollarSign className="h-5 w-5 text-green-600" />
                <h3 className="font-semibold text-green-800">Cash on Delivery</h3>
              </div>
              <div className="text-sm text-green-700 space-y-2">
                <p>• Please keep <strong>${orderData.total.toFixed(2)}</strong> ready in cash</p>
                <p>• Payment will be collected upon delivery</p>
                <p>• Please have exact change if possible</p>
                <p>• COD orders may take 1-2 additional days to process</p>
              </div>
            </div>
          )}

          {/* Order Timeline */}
          <div className="mb-8">
            <h3 className="font-semibold mb-4">What's Next?</h3>
            <div className="space-y-4">
              <div className="flex items-center space-x-3">
                <div className="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                </div>
                <div className="text-left">
                  <p className="font-medium">Order Confirmed</p>
                  <p className="text-sm text-gray-600">
                    {isCOD ? "We've received your COD order" : "We've received your order"}
                  </p>
                </div>
              </div>

              <div className="flex items-center space-x-3">
                <div className="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <Package className="h-4 w-4 text-blue-600" />
                </div>
                <div className="text-left">
                  <p className="font-medium">
                    {isCOD ? "Verification & Processing" : "Processing"}
                  </p>
                  <p className="text-sm text-gray-600">
                    {isCOD
                      ? "We're verifying your order and preparing items"
                      : "We're preparing your items"
                    }
                  </p>
                </div>
              </div>

              <div className="flex items-center space-x-3">
                <div className="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center">
                  <Truck className="h-4 w-4 text-gray-600" />
                </div>
                <div className="text-left">
                  <p className="font-medium">Shipped</p>
                  <p className="text-sm text-gray-600">
                    Estimated delivery: {estimatedDelivery.toLocaleDateString()}
                    {isCOD && " (COD orders may take 1-2 extra days)"}
                  </p>
                </div>
              </div>
            </div>
          </div>

          {/* Order Summary */}
          <div className="border-t border-gray-200 pt-6 mb-6">
            <h3 className="font-semibold mb-4 text-left">Order Summary</h3>
            <div className="space-y-3">
              {orderData.items.map((item) => (
                <div key={item.product.id} className="flex items-center space-x-4">
                  <img
                    src={item.product.image}
                    alt={item.product.name}
                    className="w-12 h-12 object-cover rounded"
                  />
                  <div className="flex-1 text-left">
                    <p className="font-medium">{item.product.name}</p>
                    <p className="text-sm text-gray-600">Qty: {item.quantity}</p>
                  </div>
                  <p className="font-medium">${(item.product.price * item.quantity).toFixed(2)}</p>
                </div>
              ))}
            </div>
          </div>

          {/* Shipping & Payment Info */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6 text-left">
            <div className="bg-gray-50 rounded-lg p-4">
              <h4 className="font-medium mb-2 flex items-center space-x-2">
                <Truck className="h-4 w-4" />
                <span>Shipping Address</span>
              </h4>
              <div className="text-sm text-gray-600">
                <p>{orderData.shippingAddress.firstName} {orderData.shippingAddress.lastName}</p>
                <p>{orderData.shippingAddress.address1}</p>
                {orderData.shippingAddress.address2 && <p>{orderData.shippingAddress.address2}</p>}
                <p>{orderData.shippingAddress.city}, {orderData.shippingAddress.state} {orderData.shippingAddress.zipCode}</p>
              </div>
            </div>

            <div className="bg-gray-50 rounded-lg p-4">
              <h4 className="font-medium mb-2 flex items-center space-x-2">
                {isCOD ? <DollarSign className="h-4 w-4" /> : <CreditCard className="h-4 w-4" />}
                <span>Payment Method</span>
              </h4>
              <div className="text-sm text-gray-600">
                {isCOD ? (
                  <div>
                    <p className="font-medium text-green-700">Cash on Delivery</p>
                    <p>Pay ${orderData.total.toFixed(2)} upon delivery</p>
                  </div>
                ) : orderData.paymentMethod.type === 'card' ? (
                  <>
                    <p className="capitalize">{orderData.paymentMethod.brand} ending in {orderData.paymentMethod.last4}</p>
                    <p>Expires {orderData.paymentMethod.expiryMonth}/{orderData.paymentMethod.expiryYear}</p>
                  </>
                ) : (
                  <p>PayPal</p>
                )}
              </div>
            </div>
          </div>

          {/* Action Buttons */}
          <div className="flex flex-col sm:flex-row gap-3">
            <button
              onClick={onClose}
              className="flex-1 bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition-colors font-medium"
            >
              Continue Shopping
            </button>

            <button className="flex-1 border border-gray-300 text-gray-700 py-3 rounded-lg hover:bg-gray-50 transition-colors font-medium flex items-center justify-center space-x-2">
              <Download className="h-4 w-4" />
              <span>Download Receipt</span>
            </button>
          </div>

          <p className="text-xs text-gray-500 mt-4">
            A confirmation email has been sent to your email address.
            {isCOD && " You will receive a call to confirm your COD order within 24 hours."}
          </p>
        </div>
      </div>
    </div>
  );
};
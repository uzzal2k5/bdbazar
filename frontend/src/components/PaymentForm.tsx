import React, { useState } from 'react';
import { CreditCard, Lock, DollarSign, AlertCircle } from 'lucide-react';
import { PaymentMethod, Address } from '../types';
import { AddressForm } from './AddressForm';

interface PaymentFormProps {
  onSubmit: (paymentMethod: PaymentMethod, billingAddress?: Address) => void;
  requiresBilling: boolean;
  initialBilling?: Address | null;
  orderTotal?: number;
  shippingMethodSupportsCOD?: boolean;
}

export const PaymentForm: React.FC<PaymentFormProps> = ({
  onSubmit,
  requiresBilling,
  initialBilling,
  orderTotal = 0,
  shippingMethodSupportsCOD = true,
}) => {
  const [paymentType, setPaymentType] = useState<'card' | 'paypal' | 'cod'>('card');
  const [showBillingForm, setShowBillingForm] = useState(false);
  const [billingAddress, setBillingAddress] = useState<Address | null>(initialBilling || null);

  const [cardData, setCardData] = useState({
    cardNumber: '',
    expiryMonth: '',
    expiryYear: '',
    cvv: '',
    cardholderName: '',
  });

  // COD fee calculation (2% of order total, minimum $2, maximum $10)
  const codFee = Math.min(Math.max(orderTotal * 0.02, 2), 10);
  const codLimit = 500; // Maximum order value for COD

  const handleCardSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (requiresBilling && !billingAddress) {
      setShowBillingForm(true);
      return;
    }

    const paymentMethod: PaymentMethod = {
      id: Date.now().toString(),
      type: 'card',
      last4: cardData.cardNumber.slice(-4),
      brand: getCardBrand(cardData.cardNumber),
      expiryMonth: parseInt(cardData.expiryMonth),
      expiryYear: parseInt(cardData.expiryYear),
      isDefault: false,
    };

    onSubmit(paymentMethod, billingAddress || undefined);
  };

  const handlePayPalSubmit = () => {
    const paymentMethod: PaymentMethod = {
      id: Date.now().toString(),
      type: 'paypal',
      isDefault: false,
    };

    onSubmit(paymentMethod);
  };

  const handleCODSubmit = () => {
    if (requiresBilling && !billingAddress) {
      setShowBillingForm(true);
      return;
    }

    const paymentMethod: PaymentMethod = {
      id: Date.now().toString(),
      type: 'cod',
      isDefault: false,
    };

    onSubmit(paymentMethod, billingAddress || undefined);
  };

  const handleBillingSubmit = (address: Address) => {
    setBillingAddress(address);
    setShowBillingForm(false);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;

    if (name === 'cardNumber') {
      // Format card number with spaces
      const formatted = value.replace(/\s/g, '').replace(/(.{4})/g, '$1 ').trim();
      setCardData({ ...cardData, [name]: formatted });
    } else {
      setCardData({ ...cardData, [name]: value });
    }
  };

  const getCardBrand = (cardNumber: string): string => {
    const number = cardNumber.replace(/\s/g, '');
    if (number.startsWith('4')) return 'visa';
    if (number.startsWith('5') || number.startsWith('2')) return 'mastercard';
    if (number.startsWith('3')) return 'amex';
    return 'unknown';
  };

  const currentYear = new Date().getFullYear();
  const years = Array.from({ length: 20 }, (_, i) => currentYear + i);
  const months = Array.from({ length: 12 }, (_, i) => (i + 1).toString().padStart(2, '0'));

  if (showBillingForm) {
    return (
      <div>
        <h4 className="font-medium mb-4">Billing Address</h4>
        <AddressForm
          type="billing"
          onSubmit={handleBillingSubmit}
          initialData={billingAddress}
        />
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Payment Method Selection */}
      <div className="space-y-3">
        <label className="flex items-center space-x-3 p-4 border border-gray-300 rounded-lg cursor-pointer hover:bg-gray-50">
          <input
            type="radio"
            name="paymentType"
            value="card"
            checked={paymentType === 'card'}
            onChange={(e) => setPaymentType(e.target.value as 'card')}
            className="text-blue-600 focus:ring-blue-500"
          />
          <CreditCard className="h-5 w-5 text-gray-600" />
          <span className="font-medium">Credit or Debit Card</span>
        </label>

        <label className="flex items-center space-x-3 p-4 border border-gray-300 rounded-lg cursor-pointer hover:bg-gray-50">
          <input
            type="radio"
            name="paymentType"
            value="paypal"
            checked={paymentType === 'paypal'}
            onChange={(e) => setPaymentType(e.target.value as 'paypal')}
            className="text-blue-600 focus:ring-blue-500"
          />
          <div className="w-5 h-5 bg-blue-600 rounded text-white text-xs flex items-center justify-center font-bold">
            P
          </div>
          <span className="font-medium">PayPal</span>
        </label>

        {/* Cash on Delivery Option */}
        <label className={`flex items-center space-x-3 p-4 border rounded-lg cursor-pointer transition-all ${
          !shippingMethodSupportsCOD || orderTotal > codLimit
            ? 'border-gray-200 bg-gray-50 cursor-not-allowed opacity-60'
            : 'border-gray-300 hover:bg-gray-50'
        }`}>
          <input
            type="radio"
            name="paymentType"
            value="cod"
            checked={paymentType === 'cod'}
            onChange={(e) => setPaymentType(e.target.value as 'cod')}
            disabled={!shippingMethodSupportsCOD || orderTotal > codLimit}
            className="text-blue-600 focus:ring-blue-500 disabled:opacity-50"
          />
          <DollarSign className="h-5 w-5 text-green-600" />
          <div className="flex-1">
            <span className="font-medium">Cash on Delivery (COD)</span>
            <p className="text-sm text-gray-600">Pay when you receive your order</p>
            {orderTotal > 0 && orderTotal <= codLimit && (
              <p className="text-xs text-green-600">COD Fee: ${codFee.toFixed(2)}</p>
            )}
          </div>
        </label>

        {/* COD Limitations Notice */}
        {(!shippingMethodSupportsCOD || orderTotal > codLimit) && (
          <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-3">
            <div className="flex items-start space-x-2">
              <AlertCircle className="h-4 w-4 text-yellow-600 mt-0.5 flex-shrink-0" />
              <div className="text-sm text-yellow-800">
                {!shippingMethodSupportsCOD && (
                  <p>Cash on Delivery is not available for the selected shipping method.</p>
                )}
                {orderTotal > codLimit && (
                  <p>Cash on Delivery is only available for orders under ${codLimit}.</p>
                )}
              </div>
            </div>
          </div>
        )}
      </div>

      {paymentType === 'card' && (
        <form onSubmit={handleCardSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Cardholder Name *
            </label>
            <input
              type="text"
              name="cardholderName"
              value={cardData.cardholderName}
              onChange={handleInputChange}
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Card Number *
            </label>
            <input
              type="text"
              name="cardNumber"
              value={cardData.cardNumber}
              onChange={handleInputChange}
              placeholder="1234 5678 9012 3456"
              maxLength={19}
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div className="grid grid-cols-3 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Month *
              </label>
              <select
                name="expiryMonth"
                value={cardData.expiryMonth}
                onChange={handleInputChange}
                required
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">MM</option>
                {months.map(month => (
                  <option key={month} value={month}>{month}</option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Year *
              </label>
              <select
                name="expiryYear"
                value={cardData.expiryYear}
                onChange={handleInputChange}
                required
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">YYYY</option>
                {years.map(year => (
                  <option key={year} value={year}>{year}</option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                CVV *
              </label>
              <input
                type="text"
                name="cvv"
                value={cardData.cvv}
                onChange={handleInputChange}
                placeholder="123"
                maxLength={4}
                required
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
          </div>

          {requiresBilling && (
            <div className="bg-gray-50 rounded-lg p-4">
              <h4 className="font-medium mb-2">Billing Address</h4>
              {billingAddress ? (
                <div className="text-sm text-gray-600">
                  <p>{billingAddress.firstName} {billingAddress.lastName}</p>
                  <p>{billingAddress.address1}</p>
                  <p>{billingAddress.city}, {billingAddress.state} {billingAddress.zipCode}</p>
                  <button
                    type="button"
                    onClick={() => setShowBillingForm(true)}
                    className="text-blue-600 hover:text-blue-700 mt-2"
                  >
                    Change billing address
                  </button>
                </div>
              ) : (
                <button
                  type="button"
                  onClick={() => setShowBillingForm(true)}
                  className="text-blue-600 hover:text-blue-700"
                >
                  Add billing address
                </button>
              )}
            </div>
          )}

          <div className="flex items-center space-x-2 text-sm text-gray-600">
            <Lock className="h-4 w-4" />
            <span>Your payment information is encrypted and secure</span>
          </div>

          <button
            type="submit"
            className="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition-colors font-medium"
          >
            Continue to Review
          </button>
        </form>
      )}

      {paymentType === 'paypal' && (
        <div className="space-y-4">
          <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
            <p className="text-blue-800 text-sm">
              You will be redirected to PayPal to complete your payment securely.
            </p>
          </div>

          <button
            onClick={handlePayPalSubmit}
            className="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition-colors font-medium"
          >
            Continue with PayPal
          </button>
        </div>
      )}

      {paymentType === 'cod' && (
        <div className="space-y-4">
          <div className="bg-green-50 border border-green-200 rounded-lg p-4">
            <h4 className="font-medium text-green-800 mb-2">Cash on Delivery</h4>
            <div className="text-sm text-green-700 space-y-2">
              <p>• Pay in cash when your order is delivered to your doorstep</p>
              <p>• No advance payment required</p>
              <p>• COD fee: ${codFee.toFixed(2)} (included in total)</p>
              <p>• Available for orders up to ${codLimit}</p>
              <p>• Please keep exact change ready for the delivery person</p>
            </div>
          </div>

          <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
            <div className="flex items-start space-x-2">
              <AlertCircle className="h-4 w-4 text-yellow-600 mt-0.5 flex-shrink-0" />
              <div className="text-sm text-yellow-800">
                <p className="font-medium mb-1">Important Notes:</p>
                <ul className="list-disc list-inside space-y-1">
                  <li>Orders may take longer to process due to verification</li>
                  <li>COD orders cannot be cancelled once shipped</li>
                  <li>Refusal to accept delivery may result in additional charges</li>
                </ul>
              </div>
            </div>
          </div>

          {requiresBilling && (
            <div className="bg-gray-50 rounded-lg p-4">
              <h4 className="font-medium mb-2">Billing Address</h4>
              {billingAddress ? (
                <div className="text-sm text-gray-600">
                  <p>{billingAddress.firstName} {billingAddress.lastName}</p>
                  <p>{billingAddress.address1}</p>
                  <p>{billingAddress.city}, {billingAddress.state} {billingAddress.zipCode}</p>
                  <button
                    type="button"
                    onClick={() => setShowBillingForm(true)}
                    className="text-blue-600 hover:text-blue-700 mt-2"
                  >
                    Change billing address
                  </button>
                </div>
              ) : (
                <button
                  type="button"
                  onClick={() => setShowBillingForm(true)}
                  className="text-blue-600 hover:text-blue-700"
                >
                  Add billing address
                </button>
              )}
            </div>
          )}

          <button
            onClick={handleCODSubmit}
            className="w-full bg-green-600 text-white py-3 rounded-lg hover:bg-green-700 transition-colors font-medium"
          >
            Continue with Cash on Delivery
          </button>
        </div>
      )}
    </div>
  );
};
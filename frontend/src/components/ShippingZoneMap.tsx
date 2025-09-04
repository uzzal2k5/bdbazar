import React, { useState } from 'react';
import { MapPin, Info, Clock, DollarSign } from 'lucide-react';

interface ShippingZone {
  id: string;
  name: string;
  states: string[];
  baseRate: number;
  estimatedDays: string;
  color: string;
}

interface ShippingZoneMapProps {
  onZoneSelect?: (zone: ShippingZone) => void;
  selectedZip?: string;
}

export const ShippingZoneMap: React.FC<ShippingZoneMapProps> = ({
  onZoneSelect,
  selectedZip,
}) => {
  const [selectedZone, setSelectedZone] = useState<ShippingZone | null>(null);

  const shippingZones: ShippingZone[] = [
    {
      id: 'zone1',
      name: 'Zone 1 - East Coast',
      states: ['NY', 'NJ', 'CT', 'MA', 'RI', 'VT', 'NH', 'ME', 'PA', 'DE', 'MD', 'DC', 'VA', 'WV', 'NC', 'SC', 'GA', 'FL'],
      baseRate: 5.99,
      estimatedDays: '2-4 business days',
      color: '#3B82F6',
    },
    {
      id: 'zone2',
      name: 'Zone 2 - Central',
      states: ['OH', 'KY', 'TN', 'AL', 'MS', 'IN', 'IL', 'WI', 'MI', 'MN', 'IA', 'MO', 'AR', 'LA', 'ND', 'SD', 'NE', 'KS', 'OK', 'TX'],
      baseRate: 7.99,
      estimatedDays: '3-5 business days',
      color: '#10B981',
    },
    {
      id: 'zone3',
      name: 'Zone 3 - Mountain',
      states: ['MT', 'WY', 'CO', 'NM', 'ID', 'UT', 'AZ', 'NV'],
      baseRate: 9.99,
      estimatedDays: '4-6 business days',
      color: '#F59E0B',
    },
    {
      id: 'zone4',
      name: 'Zone 4 - West Coast',
      states: ['WA', 'OR', 'CA', 'AK', 'HI'],
      baseRate: 11.99,
      estimatedDays: '5-7 business days',
      color: '#EF4444',
    },
  ];

  const getZoneByZip = (zipCode: string): ShippingZone | null => {
    if (!zipCode || zipCode.length < 5) return null;

    const firstDigit = parseInt(zipCode.charAt(0));

    if (firstDigit <= 2) return shippingZones[0]; // East Coast
    if (firstDigit <= 5) return shippingZones[1]; // Central
    if (firstDigit <= 7) return shippingZones[2]; // Mountain
    return shippingZones[3]; // West Coast
  };

  const handleZoneClick = (zone: ShippingZone) => {
    setSelectedZone(zone);
    if (onZoneSelect) {
      onZoneSelect(zone);
    }
  };

  const currentZone = selectedZip ? getZoneByZip(selectedZip) : null;

  return (
    <div className="bg-white rounded-lg border border-gray-200 p-6">
      <div className="flex items-center space-x-2 mb-4">
        <MapPin className="h-5 w-5 text-blue-600" />
        <h3 className="text-lg font-semibold">Shipping Zones</h3>
      </div>

      {/* Current Zone Display */}
      {currentZone && (
        <div className="mb-6 p-4 bg-blue-50 border border-blue-200 rounded-lg">
          <div className="flex items-center space-x-3">
            <div
              className="w-4 h-4 rounded-full"
              style={{ backgroundColor: currentZone.color }}
            ></div>
            <div>
              <p className="font-medium text-blue-900">Your Shipping Zone</p>
              <p className="text-sm text-blue-700">
                ZIP {selectedZip} is in {currentZone.name}
              </p>
            </div>
          </div>
          <div className="mt-3 grid grid-cols-2 gap-4 text-sm">
            <div className="flex items-center space-x-2">
              <DollarSign className="h-4 w-4 text-blue-600" />
              <span>Base Rate: ${currentZone.baseRate}</span>
            </div>
            <div className="flex items-center space-x-2">
              <Clock className="h-4 w-4 text-blue-600" />
              <span>{currentZone.estimatedDays}</span>
            </div>
          </div>
        </div>
      )}

      {/* Zone List */}
      <div className="space-y-3 mb-6">
        {shippingZones.map((zone) => (
          <div
            key={zone.id}
            onClick={() => handleZoneClick(zone)}
            className={`p-4 border rounded-lg cursor-pointer transition-all ${
              selectedZone?.id === zone.id || currentZone?.id === zone.id
                ? 'border-blue-500 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300 hover:bg-gray-50'
            }`}
          >
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-3">
                <div
                  className="w-4 h-4 rounded-full"
                  style={{ backgroundColor: zone.color }}
                ></div>
                <div>
                  <h4 className="font-medium">{zone.name}</h4>
                  <p className="text-sm text-gray-600">
                    {zone.states.length} states covered
                  </p>
                </div>
              </div>
              <div className="text-right">
                <p className="font-medium">${zone.baseRate}</p>
                <p className="text-sm text-gray-600">{zone.estimatedDays}</p>
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Zone Details */}
      {selectedZone && (
        <div className="border-t border-gray-200 pt-4">
          <h4 className="font-medium mb-3 flex items-center space-x-2">
            <Info className="h-4 w-4" />
            <span>{selectedZone.name} Details</span>
          </h4>

          <div className="bg-gray-50 rounded-lg p-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
              <div>
                <p className="text-sm text-gray-600 mb-1">Base Shipping Rate</p>
                <p className="font-medium text-lg">${selectedZone.baseRate}</p>
              </div>
              <div>
                <p className="text-sm text-gray-600 mb-1">Estimated Delivery</p>
                <p className="font-medium">{selectedZone.estimatedDays}</p>
              </div>
            </div>

            <div>
              <p className="text-sm text-gray-600 mb-2">States in this zone:</p>
              <div className="flex flex-wrap gap-1">
                {selectedZone.states.map((state) => (
                  <span
                    key={state}
                    className="px-2 py-1 bg-white border border-gray-200 rounded text-xs font-medium"
                  >
                    {state}
                  </span>
                ))}
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Shipping Notes */}
      <div className="mt-6 p-4 bg-yellow-50 border border-yellow-200 rounded-lg">
        <div className="flex items-start space-x-2">
          <Info className="h-4 w-4 text-yellow-600 mt-0.5 flex-shrink-0" />
          <div className="text-sm text-yellow-800">
            <p className="font-medium mb-1">Shipping Information</p>
            <ul className="list-disc list-inside space-y-1 text-yellow-700">
              <li>Rates shown are base rates and may vary based on package weight and dimensions</li>
              <li>Free shipping available on orders over $50 within Zone 1 and Zone 2</li>
              <li>Express and overnight options available for all zones</li>
              <li>Alaska and Hawaii may have additional surcharges</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};
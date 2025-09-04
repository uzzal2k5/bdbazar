export interface ShippingZone {
    id: string;
    name: string;
    states: string[];
    countries?: string[];
    baseRate: number;
    freeShippingThreshold: number;
    estimatedDays: string;
    restrictions?: string[];
}

export interface ShippingRestriction {
    productId: string;
    restrictedZones: string[];
    reason: string;
    alternatives?: string[];
}

export interface ShippingInsurance {
    id: string;
    name: string;
    description: string;
    cost: number;
    coverage: number;
    required: boolean;
}

export interface DeliveryWindow {
    id: string;
    name: string;
    description: string;
    timeRange: string;
    additionalCost: number;
    available: boolean;
}

export interface ShippingNotification {
    id: string;
    type: 'sms' | 'email' | 'push';
    event: 'shipped' | 'out_for_delivery' | 'delivered' | 'exception';
    enabled: boolean;
    contact: string;
}

export interface PackageTracking {
    trackingNumber: string;
    carrier: string;
    status: string;
    estimatedDelivery: string;
    currentLocation: string;
    events: TrackingEvent[];
    deliveryInstructions?: string;
    signatureRequired: boolean;
}

export interface ShippingLabel {
    id: string;
    trackingNumber: string;
    carrier: string;
    service: string;
    labelUrl: string;
    cost: number;
    weight: number;
    dimensions: {
    length: number;
    width: number;
    height: number;
};
    createdAt: string;
}

export interface ShippingRate {
    carrierId: string;
    carrierName: string;
    serviceId: string;
    serviceName: string;
    cost: number;
    estimatedDays: number;
    deliveryDate: string;
    features: string[];
    restrictions?: string[];
}

export interface ShippingQuote {
    id: string;
    origin: Address;
    destination: Address;
    packages: PackageInfo[];
    rates: ShippingRate[];
    createdAt: string;
    expiresAt: string;
}

export interface PackageInfo {
    weight: number;
    dimensions: {
        length: number;
        width: number;
        height: number;
    };
    value: number;
    description: string;
    fragile: boolean;
    hazardous: boolean;
}

export interface ShippingPreferences {
    defaultMethod: string;
    notifications: ShippingNotification[];
    deliveryInstructions: string;
    signatureRequired: boolean;
    leaveAtDoor: boolean;
    preferredDeliveryWindow?: DeliveryWindow;
    insurancePreference: 'none' | 'basic' | 'full';
}

export interface ShippingCost {
    base: number;
    weight: number;
    distance: number;
    size: number;
    insurance: number;
    handling: number;
    fuel: number;
    total: number;
    discounts: ShippingDiscount[];
}

export interface ShippingDiscount {
    id: string;
    name: string;
    type: 'percentage' | 'fixed' | 'free_shipping';
    value: number;
    applied: number;
    reason: string;
}

export interface ShippingCarrier {
    id: string;
    name: string;
    logo: string;
    services: ShippingService[];
    trackingUrl: string;
    supportsCOD: boolean;
    supportsInsurance: boolean;
    maxWeight: number;
    maxDimensions: {
        length: number;
        width: number;
        height: number;
    };
}

export interface ShippingService {
    id: string;
    name: string;
    description: string;
    type: 'ground' | 'air' | 'express' | 'overnight' | 'international';
    features: string[];
    restrictions: string[];
    transitTime: string;
    cutoffTime: string;
}
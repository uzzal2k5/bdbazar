import { Product } from '../types';

export const generateMockProducts = (): Product[] => {
  return [
    {
      id: '1',
      name: 'Wireless Bluetooth Headphones',
      description: 'Premium quality wireless headphones with noise cancellation and 30-hour battery life.',
      price: 99.99,
      category: 'electronics',
      image: 'https://images.pexels.com/photos/3394650/pexels-photo-3394650.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller1',
      sellerName: 'TechGear Pro',
      rating: 4.5,
      reviews: [
        {
          id: '1',
          userId: 'user1',
          userName: 'John Doe',
          rating: 5,
          comment: 'Amazing sound quality and comfortable fit!',
          date: '2024-01-15'
        },
        {
          id: '2',
          userId: 'user2',
          userName: 'Jane Smith',
          rating: 4,
          comment: 'Great headphones, battery life is excellent.',
          date: '2024-01-10'
        }
      ],
      stock: 25,
      tags: ['wireless', 'bluetooth', 'headphones', 'noise-cancellation']
    },
    {
      id: '2',
      name: 'Organic Cotton T-Shirt',
      description: 'Soft and comfortable 100% organic cotton t-shirt available in multiple colors.',
      price: 24.99,
      category: 'clothing',
      image: 'https://images.pexels.com/photos/996329/pexels-photo-996329.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller2',
      sellerName: 'EcoFashion',
      rating: 4.2,
      reviews: [
        {
          id: '3',
          userId: 'user3',
          userName: 'Mike Johnson',
          rating: 4,
          comment: 'Very comfortable and good quality material.',
          date: '2024-01-20'
        }
      ],
      stock: 50,
      tags: ['organic', 'cotton', 'eco-friendly', 'casual']
    },
    {
      id: '3',
      name: 'Smart Fitness Watch',
      description: 'Advanced fitness tracking with heart rate monitoring, GPS, and smartphone integration.',
      price: 199.99,
      category: 'electronics',
      image: 'https://images.pexels.com/photos/393047/pexels-photo-393047.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller3',
      sellerName: 'FitTech Solutions',
      rating: 4.7,
      reviews: [
        {
          id: '4',
          userId: 'user4',
          userName: 'Sarah Wilson',
          rating: 5,
          comment: 'Perfect for tracking my workouts and daily activities.',
          date: '2024-01-18'
        }
      ],
      stock: 15,
      tags: ['fitness', 'smartwatch', 'health', 'gps']
    },
    {
      id: '4',
      name: 'Vintage Leather Jacket',
      description: 'Classic vintage-style leather jacket made from genuine leather with premium finishing.',
      price: 149.99,
      category: 'clothing',
      image: 'https://images.pexels.com/photos/1040945/pexels-photo-1040945.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller4',
      sellerName: 'Vintage Wardrobe',
      rating: 4.8,
      reviews: [
        {
          id: '5',
          userId: 'user5',
          userName: 'David Brown',
          rating: 5,
          comment: 'Excellent quality leather and perfect fit.',
          date: '2024-01-12'
        }
      ],
      stock: 8,
      tags: ['leather', 'vintage', 'jacket', 'fashion']
    },
    {
      id: '5',
      name: 'Minimalist Desk Lamp',
      description: 'Modern minimalist LED desk lamp with adjustable brightness and USB charging port.',
      price: 39.99,
      category: 'home',
      image: 'https://images.pexels.com/photos/1112598/pexels-photo-1112598.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller5',
      sellerName: 'Modern Living',
      rating: 4.3,
      reviews: [
        {
          id: '6',
          userId: 'user6',
          userName: 'Lisa Chen',
          rating: 4,
          comment: 'Great design and very functional.',
          date: '2024-01-14'
        }
      ],
      stock: 30,
      tags: ['led', 'desk', 'lamp', 'minimalist', 'usb']
    },
    {
      id: '6',
      name: 'Professional Camera Tripod',
      description: 'Heavy-duty aluminum tripod with adjustable height and 360-degree rotation.',
      price: 79.99,
      category: 'electronics',
      image: 'https://images.pexels.com/photos/51383/photo-camera-subject-photographer-51383.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller6',
      sellerName: 'Photo Pro Gear',
      rating: 4.6,
      reviews: [
        {
          id: '7',
          userId: 'user7',
          userName: 'Mark Thompson',
          rating: 5,
          comment: 'Sturdy and reliable, perfect for professional photography.',
          date: '2024-01-16'
        }
      ],
      stock: 12,
      tags: ['camera', 'tripod', 'photography', 'professional']
    },
    {
      id: '7',
      name: 'Bestselling Mystery Novel',
      description: 'Gripping mystery novel that will keep you on the edge of your seat.',
      price: 14.99,
      category: 'books',
      image: 'https://images.pexels.com/photos/46274/pexels-photo-46274.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller7',
      sellerName: 'BookWorld',
      rating: 4.4,
      reviews: [
        {
          id: '8',
          userId: 'user8',
          userName: 'Emily Davis',
          rating: 4,
          comment: 'Couldn\'t put it down! Great plot twists.',
          date: '2024-01-13'
        }
      ],
      stock: 100,
      tags: ['mystery', 'novel', 'fiction', 'bestseller']
    },
    {
      id: '8',
      name: 'Yoga Mat with Carrying Strap',
      description: 'Non-slip yoga mat made from eco-friendly materials with convenient carrying strap.',
      price: 29.99,
      category: 'sports',
      image: 'https://images.pexels.com/photos/3822668/pexels-photo-3822668.jpeg?auto=compress&cs=tinysrgb&w=800',
      sellerId: 'seller8',
      sellerName: 'Zen Fitness',
      rating: 4.1,
      reviews: [
        {
          id: '9',
          userId: 'user9',
          userName: 'Amy Rodriguez',
          rating: 4,
          comment: 'Good quality mat, non-slip surface works well.',
          date: '2024-01-17'
        }
      ],
      stock: 40,
      tags: ['yoga', 'mat', 'fitness', 'eco-friendly']
    }
  ];
};
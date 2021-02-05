db.createCollection('customer', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['username', 'email', 'password', 'dob', 'avatar', 'favorite', 'verified'], properties: {
        username: {
          bsonType: 'string', minLength: 3,
          maxLength: 20
        }, email: { bsonType: 'string' }, password: {
          bsonType: 'string', minLength: 60,
          maxLength: 60
        }, dob: { bsonType: 'date' }, avatar: { bsonType: 'string' }, favorite: { bsonType: 'array', items: { bsonType: 'objectId', uniqueItems: true } }, verified: { bsonType: 'bool' }
      }
    }
  }
});

db.createCollection('business', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['username', 'email', 'password', 'businessName', 'type', 'tags', 'description', 'location', 'address', 'displayImage', 'images', 'placement', 'menu', 'status', 'verified'], properties: {
        username: {
          bsonType: 'string', minLength: 3,
          maxLength: 20
        }, email: { bsonType: 'string' }, password: {
          bsonType: 'string', minLength: 60,
          maxLength: 60
        }, businessName: {
          bsonType: 'string', minLength: 3,
          maxLength: 50
        }, type: { bsonType: 'string' }, tags: {
          bsonType: 'array', items: {
            bsonType: 'string', uniqueItems: true,
            maxItems: 5
          }
        }, description: { bsonType: 'string', maxLength: 50 }, location: { bsonType: 'object' }, address: {
          bsonType: 'string', minLength: 3,
          maxLength: 50
        }, displayImage: { bsonType: 'string' }, images: {
          bsonType: 'array', items: {
            bsonType: 'string', uniqueItems: true,
            maxItems: 10
          }
        }, placement: {
          bsonType: 'array', items: {
            required: ['name', 'floor', 'type', 'space', 'price', 'x', 'y', 'width', 'height', 'rotation'], properties: {
              name: {
                bsonType: 'string', minLength: 1,
                maxLength: 2
              }, floor: { bsonType: 'int', min: 1 }, type: { bsonType: 'string', enum: ['SEAT', 'TABLE'] }, space: { bsonType: 'int', min: 1 }, price: { bsonType: 'decimal' }, x: { bsonType: 'double' }, y: { bsonType: 'double' }, width: { bsonType: 'double' }, height: { bsonType: 'double' }, rotation: { bsonType: 'double' }
            }
          }
        }, menu: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'image', 'price', 'max'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, description: { bsonType: 'string', maxLength: 125 }, image: { bsonType: 'string' }, price: { bsonType: 'decimal', min: 0 }, max: { bsonType: 'int' }
            }
          }
        }, status: { bsonType: 'int' }, verified: { bsonType: 'bool' }
      }
    }
  }
});

db.createCollection('reservation', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['businessId', 'name', 'start', 'end', 'placement', 'menu'], properties: {
        businessId: { bsonType: 'objectId' }, name: {
          bsonType: 'string', minLength: 3,
          maxLength: 20
        }, start: { bsonType: 'date' }, end: { bsonType: 'date' }, placement: {
          bsonType: 'array', items: {
            required: ['name', 'floor', 'type', 'space', 'price', 'user', 'status', 'x', 'y', 'width', 'height', 'rotation'], properties: {
              name: {
                bsonType: 'string', minLength: 1,
                maxLength: 2
              }, floor: { bsonType: 'int', min: 1 }, type: { bsonType: 'string', enum: ['SEAT', 'TABLE'] }, space: { bsonType: 'int', min: 1 }, price: { bsonType: 'decimal' }, user: { bsonType: 'objectId' }, status: { bsonType: 'string', enum: ['EMPTY', 'TAKEN', 'IN PROGRESS'] }, x: { bsonType: 'double' }, y: { bsonType: 'double' }, width: { bsonType: 'double' }, height: { bsonType: 'double' }, rotation: { bsonType: 'double' }
            }
          }
        }, menu: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'image', 'price', 'max'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, description: { bsonType: 'string', maxLength: 125 }, image: { bsonType: 'string' }, price: { bsonType: 'decimal', min: 0 }, max: { bsonType: 'int' }
            }
          }
        }
      }
    }
  }
});

db.createCollection('order', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['reservationId', 'customerId', 'businessId', 'start', 'end', 'seats', 'preorder', 'totalPrice', 'status', 'image', 'extraSpace'], properties: {
        reservationId: { bsonType: 'objectId' }, customerId: { bsonType: 'objectId' }, businessId: { bsonType: 'objectId' }, start: { bsonType: 'date' }, end: { bsonType: 'date' }, seats: {
          bsonType: 'array', items: {
            required: ['name', 'floor', 'type', 'space', 'price', 'x', 'y', 'width', 'height', 'rotation'], properties: {
              name: {
                bsonType: 'string', minLength: 1,
                maxLength: 2
              }, floor: { bsonType: 'int', min: 1 }, type: { bsonType: 'string', enum: ['SEAT', 'TABLE'] }, space: { bsonType: 'int', min: 1 }, price: { bsonType: 'decimal' }, x: { bsonType: 'double' }, y: { bsonType: 'double' }, width: { bsonType: 'double' }, height: { bsonType: 'double' }, rotation: { bsonType: 'double' }
            }
          }
        }, preorder: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'image', 'quantity', 'price'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, description: { bsonType: 'string' }, image: { bsonType: 'string' }, quantity: { bsonType: 'int', min: 1 }, price: { bsonType: 'decimal', min: 0 }
            }
          }
        }, totalPrice: { bsonType: 'decimal', min: 0 }, status: { bsonType: 'string', enum: ['PAID', 'USED', 'EXPIRED', 'CANCELLED'] }, image: { bsonType: 'string' }, extraSpace: { bsonType: 'int' }
      }
    }
  }
});

db.createCollection('admin', { validator: { $jsonSchema: { bsonType: 'object', required: ['username', 'password'], properties: { username: { bsonType: 'string' }, password: { bsonType: 'string' } } } } });

db.createCollection('request', { validator: { $jsonSchema: { bsonType: 'object', required: ['businessName', 'type', 'tags', 'description', 'location', 'address'], properties: { businessName: { bsonType: 'string' }, type: { bsonType: 'string' }, tags: { bsonType: 'array', items: { bsonType: 'string' } }, description: { bsonType: 'string' }, location: { bsonType: 'object' }, address: { bsonType: 'string' } } } } });  
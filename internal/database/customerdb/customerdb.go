package customerdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CustomerDB struct {
	CusCol *mongo.Collection
	BusCol *mongo.Collection
}

func (c *CustomerDB) AuthenticateCustomer(ctx context.Context, customer *typedb.Customer) (string, error) {
	res := c.CusCol.FindOne(ctx, bson.M{"_id": customer.Id})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", commonErr.NOTFOUND
		}

		return "", commonErr.INTERNAL
	}

	pw := customer.Password
	if err := res.Decode(customer); err != nil {
		return "", commonErr.INTERNAL
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(pw)); err != nil {
		return "", commonErr.NOTFOUND
	}

	return customer.Id.String(), nil
}

func (c *CustomerDB) CreateCustomer(ctx context.Context, customer *typedb.Customer) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 12)
	if err != nil {
		return "", commonErr.INTERNAL
	}

	customer.Password = string(pw)
	customer.Id = primitive.NewObjectIDFromTimestamp(time.Now())

	_, err = c.CusCol.InsertOne(ctx, customer)
	if err != nil {
		// TODO: better error handling
		return "", commonErr.INTERNAL
	}

	return customer.Id.String(), nil
}

func (c *CustomerDB) AppendFavorite(ctx context.Context, customerId string, businessId string) error {
	if res := c.BusCol.FindOne(ctx, bson.M{"_id": businessId}); res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	_, err := c.CusCol.UpdateOne(
		ctx,
		bson.M{"_id": customerId},
		bson.D{
			{"$push",
				bson.D{
					{"favorite", businessId},
				},
			},
		})

	// TODO: better error handling
	if err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

func (c *CustomerDB) RemoveFavorite(ctx context.Context, customerId string, businessId string) error {
	_, err := c.CusCol.UpdateOne(
		ctx,
		bson.M{"_id": customerId},
		bson.D{
			{"$pull",
				bson.D{
					{"favorite", businessId},
				},
			},
		})

	// TODO: better error handling
	if err != nil {
		return commonErr.INTERNAL
	}

	return nil
}
package repository

import (
	"context"
	"fmt"

	"github.com/Todari/hgt-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r repository) UpdateProperty(ctx context.Context, studentId string, property model.Property) (model.Property, error) {
	var result user
	err := r.db.Collection("hgtUser").FindOne(ctx, bson.M{"studentId": studentId}).Decode(&result)
	fmt.Println("user", result)
	if err != nil {
		return model.Property{}, ErrUserNotFound
	}
	in := bson.M{"userId": primitive.ObjectID(result.ID), "smoke": property.Smoke, "height": property.Height, "religion": property.Religion, "p": property.P}
	fmt.Println("in", in)
	out, err := r.db.
		Collection("property").
		UpdateOne(ctx, bson.M{"userId": result.ID}, bson.M{"$set": in})
	fmt.Println("out", out)
	if err != nil {
		return model.Property{}, err
	}
	if out.MatchedCount == 0 {
		return model.Property{}, ErrUserNotFound
	}
	return property, nil
}

type property struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"userId,omitempty"`
	Smoke    string             `bson:"smoke, omitempty"`
	Height   string             `bson:"height, omitempty"`
	Religion string             `bson:"religion, omitempty"`
	P        []string           `bson:"p,omitempty"`
}

func fromModelProperty(in model.Property) property {
	UserID, err := primitive.ObjectIDFromHex(in.UserID)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(UserID)
	return property{
		UserID:   UserID,
		Smoke:    in.Smoke,
		Height:   in.Height,
		Religion: in.Religion,
		P:        in.P,
	}
}

func toModelProperty(in property) model.Property {
	return model.Property{
		ID:       in.ID.String(),
		UserID:   in.UserID.String(),
		Smoke:    in.Smoke,
		Height:   in.Height,
		Religion: in.Religion,
		P:        in.P,
	}
}
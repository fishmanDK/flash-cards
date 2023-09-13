package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryDb struct {
	client *mongo.Client
}

type TitleCategory struct {
	Title string `db:"title"`
}

type Category struct {
	Title         string               `db:"title"`
	QuestionsId   []primitive.ObjectID `bson:"questions"`
	QuestionsText []string
}

type IdCategory struct {
	Id   primitive.ObjectID `bson:"_id"`
	Test primitive.ObjectID `bson:"user_id"`
}

type UserCategories struct {
	Categories      []primitive.ObjectID `db:"categories"`
	TitleCategories []string
}

func NewCategoryDb(client *mongo.Client) *CategoryDb {
	return &CategoryDb{
		client: client,
	}
}

func (a *CategoryDb) GetUserCategories(id string) (UserCategories, error) {
	var userCategories UserCategories

	userCollection := a.client.Database("anki").Collection("user")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return userCategories, err
	}
	filter := bson.M{
		"_id": objectId,
	}

	err = userCollection.FindOne(context.Background(), filter, nil).Decode(&userCategories)
	if err != nil {
		return userCategories, err
	}

	categoriesCollection := a.client.Database("anki").Collection("categories")

	filter = bson.M{
		"_id": bson.M{
			"$in": userCategories.Categories,
		},
	}

	cur, err := categoriesCollection.Find(context.Background(), filter)
	if err != nil {
		return UserCategories{}, err
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var title TitleCategory

		err = cur.Decode(&title)
		if err != nil {
			return UserCategories{}, err
		}

		userCategories.TitleCategories = append(userCategories.TitleCategories, title.Title)
	}

	return userCategories, nil
}

func (a *CategoryDb) GetUserCategoryById(userId, categoryName string) (Category, error) {
	var category Category

	categoryCollection := a.client.Database("anki").Collection("categories")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Category{}, err
	}

	filter := bson.D{
		{"user_id", objectId},
		{"title", categoryName},
	}

	err = categoryCollection.FindOne(context.Background(), filter).Decode(&category)

	if err != nil {
		return Category{}, err
	}

	resCategory, err := a.parceQuestions(&category, category.QuestionsId)
	if err != nil {
		return Category{}, err
	}

	return *resCategory, nil
}

func (a *CategoryDb) parceQuestions(category *Category, questionsIndsList []primitive.ObjectID) (*Category, error) {
	questionCollection := a.client.Database("anki").Collection("questions")

	filter := bson.M{
		"_id": bson.M{
			"$in": questionsIndsList,
		},
	}

	cur, err := questionCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var question Question

		err := cur.Decode(&question)
		if err != nil {
			return nil, err
		}

		category.QuestionsText = append(category.QuestionsText, question.QestionText)
	}

	return category, nil
}

func (a *CategoryDb) CreateCategory(userId, categoryName string) (string, error) {
	categoryCollection := a.client.Database("anki").Collection("categories")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", err
	}

	document := bson.D{
		{"user_id", objectId},
		{"title", categoryName},
		{"questions", bson.A{}},
	}

	res, err := categoryCollection.InsertOne(context.Background(), document)
	if err != nil {
		return "", err
	}

	filter := bson.D{
		{"_id", objectId},
	}

	updateData := bson.M{
		"$push": bson.M{
			"categories": res.InsertedID,
		},
	}

	userCollection := a.client.Database("anki").Collection("user")

	_, err = userCollection.UpdateOne(context.Background(), filter, updateData)
	if err != nil {
		return "", err
	}

	return objectId.Hex(), nil
}

func (a *CategoryDb) UpdateTitle(userId, oldName, newName string) error {
	categoryCollection := a.client.Database("anki").Collection("categories")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"user_id", objectId},
		{"title", oldName},
	}

	updateData := bson.M{
		"$set": bson.M{
			"title": newName,
		},
	}

	_, err = categoryCollection.UpdateOne(context.Background(), filter, updateData)
	if err != nil {
		return err
	}

	return nil
}

func (a *CategoryDb) DeleteCategory(userId, categoryName string) error {
	var categoryId IdCategory

	categoryCollection := a.client.Database("anki").Collection("categories")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"user_id", objectId},
		{"title", categoryName},
	}

	err = categoryCollection.FindOne(context.Background(), filter).Decode(&categoryId)
	if err != nil {
		return err
	}

	_, err = categoryCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	filter = bson.D{
		{"_id", objectId},
	}
	update := bson.M{
		"$pull": bson.M{
			"categories": categoryId.Id,
		},
	}

	userCollection := a.client.Database("anki").Collection("user")

	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

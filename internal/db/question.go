package db

import (
	"context"
	"log"
	"reflect"
	"strings"

	anki "github.com/fishmanDK/anki_telegram"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionDb struct {
	client *mongo.Client
}

func NewQuestionDb(client *mongo.Client) *QuestionDb {
	return &QuestionDb{
		client: client,
	}
}

type Question struct {
	Name        string `bson:"name"`
	QestionText string `bson:"question_text"`
	Answer      string `bson:"answer"`
}

type ResUserQuestions struct {
	Qestions []Question
}

type UserQestions struct {
	Qestoins []primitive.ObjectID `bson:"questions"`
}

func (a *QuestionDb) GetQuestions(userId, categoryName string) (ResUserQuestions, error) {
	var res ResUserQuestions

	categoryCollection := a.client.Database("anki").Collection("categories")
	questionCollection := a.client.Database("anki").Collection("questions")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return ResUserQuestions{}, err
	}

	categoryFilter := bson.D{
		{"user_id", objectId},
		{"title", categoryName},
	}

	var userQestions UserQestions
	err = categoryCollection.FindOne(context.Background(), categoryFilter).Decode(&userQestions)
	if err != nil {
		return ResUserQuestions{}, err
	}

	qestionFilter := bson.M{
		"_id": bson.M{
			"$in": userQestions.Qestoins,
		},
	}

	cur, err := questionCollection.Find(context.Background(), qestionFilter)
	if err != nil {
		return ResUserQuestions{}, err
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var question Question

		err = cur.Decode(&question)
		if err != nil {
			return ResUserQuestions{}, err
		}

		res.Qestions = append(res.Qestions, question)
	}
	return res, nil
}

func (a *QuestionDb) GetQuestion(userId, qestionName string) (Question, error) {
	var question Question
	questionCollection := a.client.Database("anki").Collection("questions")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Question{}, err
	}

	filter := bson.D{
		{"user_id", objectId},
		{"name", qestionName},
	}

	err = questionCollection.FindOne(context.Background(), filter).Decode(&question)
	if err != nil {
		return Question{}, err
	}
	log.Println("    ffff  ", question)
	return question, nil
}

func (a *QuestionDb) CreateQuestion(userId, categoryName string, input anki.CreateQuestion) (string, error) {
	categoryCollection := a.client.Database("anki").Collection("categories")
	questionCollection := a.client.Database("anki").Collection("questions")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", err
	}

	documentQuestion := bson.D{
		{"user_id", objectId},
		{"name", input.Name},
		{"question_text", input.QuestionText},
		{"answer", input.Answer},
	}

	res, err := questionCollection.InsertOne(context.Background(), documentQuestion)
	if err != nil {
		return "", err
	}

	filter := bson.D{
		{"user_id", objectId},
		{"title", categoryName},
	}

	updateData := bson.M{
		"$push": bson.M{
			"questions": res.InsertedID,
		},
	}

	_, err = categoryCollection.UpdateOne(context.Background(), filter, updateData)
	if err != nil {
		return "", err
	}
	return objectId.Hex(), nil
}

func (a *QuestionDb) UpdateQustion(userId, categoryName, questionName string, input anki.UpdateQuestion) error {
	questionCollection := a.client.Database("anki").Collection("questions")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"user_id", objectId},
		{"name", questionName},
	}

	v := reflect.ValueOf(input)
	t := v.Type()

	updateData := bson.M{"$set": bson.M{}}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		if !reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
			strings.ToLower(field.Name)
			updateData["$set"].(bson.M)[strings.ToLower(field.Name)] = value
		}
	}

	_, err = questionCollection.UpdateOne(context.Background(), filter, updateData)
	if err != nil {
		return err
	}

	return nil
}

type QuestionId struct {
	Id primitive.ObjectID `bson:"_id"`
}

func (a *QuestionDb) DeleteQuestion(userId, categoryName, questionName string) error {
	var questionId QuestionId

	questionCollection := a.client.Database("anki").Collection("questions")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"user_id", objectId},
		{"name", questionName},
	}

	err = questionCollection.FindOne(context.Background(), filter).Decode(&questionId)
	if err != nil {
		return err
	}

	_, err = questionCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	categoryCollection := a.client.Database("anki").Collection("categories")

	filter = bson.D{
		{"user_id", objectId},
		{"title", categoryName},
	}

	update := bson.M{
		"$pull": bson.M{
			"questions": questionId.Id,
		},
	}

	_, err = categoryCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

package db

// "log"

// type ResultInfoAboutUser struct {
// 	Username   string
// 	Email      string
// 	Categories []string
// }

// type InfoCollectionUser struct {
// 	Username   string               `db:"username"`
// 	Email      string               `db:"email"`
// 	Categories []primitive.ObjectID `db:"categories"`
// }

// func (a *AuthDb) GetInfoCollectionUser(id string) (InfoCollectionUser, error) {
// 	var dataUser InfoCollectionUser

// 	userCollection := a.client.Database("anki").Collection("user")

// 	objectId, err := primitive.ObjectIDFromHex(id)

// 	if err != nil {
// 		return InfoCollectionUser{}, errors.New("ошибка при приведении id к ObjectId (GetInfoCollectionUser)")
// 	}

// 	filter := bson.M{
// 		"_id": objectId,
// 	}

// 	err = userCollection.FindOne(context.Background(), filter, nil).Decode(&dataUser)
// 	if err != nil {
// 		return InfoCollectionUser{}, errors.New("ошибка при поиске пользователя (getInfoCollectionUser)")
// 	}

// 	return dataUser, nil
// }

// func (a *AuthDb) parceTitleCategories(resInfo *ResultInfoAboutUser, categoriesIndsList []primitive.ObjectID) (*ResultInfoAboutUser, error) {
// 	categoriesCollection := a.client.Database("anki").Collection("categories")

// 	filter := bson.M{
// 		"_id": bson.M{
// 			"$in": categoriesIndsList,
// 		},
// 	}

// 	cur, err := categoriesCollection.Find(context.Background(), filter)

// 	if err != nil {
// 		return nil, errors.New("ошибка при поиске категорий (parceTitleCategories)")
// 	}
// 	defer cur.Close(context.Background())

// 	for cur.Next(context.Background()) {
// 		var title TitleCategory

// 		err := cur.Decode(&title)
// 		if err != nil {
// 			return nil, errors.New("ошибка при декодировании названия категории (parceTitleCategories)")
// 		}

// 		resInfo.Categories = append(resInfo.Categories, title.Title)
// 	}

// 	return resInfo, nil
// }

// func (a *AuthDb) GetUserData(id string) (ResultInfoAboutUser, error) {
// 	var resultInfo ResultInfoAboutUser

// 	dataInCollectionUser, err := a.GetInfoCollectionUser(id)
// 	if err != nil {
// 		return ResultInfoAboutUser{}, err
// 	}

// 	resultInfo.Username = dataInCollectionUser.Username
// 	resultInfo.Email = dataInCollectionUser.Email

// 	resultInf, err := a.parceTitleCategories(&resultInfo, dataInCollectionUser.Categories)
// 	if err != nil {
// 		return ResultInfoAboutUser{}, err
// 	}

// 	return *resultInf, nil
// }


## Go persistence with MongoDB

In this example, we persist data in MongoDB with `gRPC`.

To connect to MongoDB:
```go
import (
    ...
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DATABASE_URI    = "mongodb://root:root@localhost:27017/"
	DATABASE_NAME   = "blogdb"
	COLLECTION_NAME = "blog"
)

func connectToMongoDB() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(DATABASE_URI))
	if err != nil {
		log.Printf("failed to create new mongoDB client: %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Printf("failed to connect to to mongoDB: %v\n", err)
	}

	return client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
}

type Server struct {
	pb.BlogServiceServer
}

var collection *mongo.Collection

func main() {
	collection = connectToMongoDB()
    ...
}
```

#### CRUD operations with MongoDB and gRPC services
Create:
```go
func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal err: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error casting to oid: %v\n", err),
		)
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}
```
Read:
```go
func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("error parse ID: %v\n", err),
		)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	err = res.Decode(data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("cannot find blog with the ID provided: %v\n", err),
		)
	}

	return documentToBlog(data), nil
}
```
Update:
```go
func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("error parse ID: %v\n", err),
		)
	}

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with Id",
		)
	}

	return &emptypb.Empty{}, nil
}
```
Delete:
```go
func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error) {
    oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("error parse ID: %v\n", err),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not delete object in mongoDB: %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}

	return &emptypb.Empty{}, nil
}
```
List All:
```go
func (s *Server) ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	cur, err := collection.Find(context.Background(), primitive.D{{}}) // empty filter
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v\n", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}

	return nil
}
```

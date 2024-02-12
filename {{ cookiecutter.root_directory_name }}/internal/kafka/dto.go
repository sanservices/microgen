package kafka

// nolint: all  // Example on how to create a new DTO
type Thing struct {
	Name string `avro:"name"`
}

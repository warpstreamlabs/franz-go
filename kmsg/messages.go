package kmsg

// Code generated by kgo/generate. DO NOT EDIT.

const MaxKey = 18

// Header is user provided metadata for a record. Kafka does not look at
// headers at all; they are solely for producers and consumers.
type Header struct {
	Key string

	Value []byte
}

// A Record is a Kafka v0.11.0.0 record. It corresponds to an individual
// message as it is written on the wire.
type Record struct {
	// Length is the length of this record on the wire of everything that
	// follows this field. It is an int32 encoded as a varint.
	Length int32

	// Attributes are record level attributes. This field currently is unused.
	Attributes int8

	// TimestampDelta is the millisecond delta of this record's timestamp
	// from the record's RecordBatch's FirstTimestamp.
	TimestampDelta int32

	// OffsetDelta is the delta of this record's offset from the record's
	// RecordBatch's FirstOffset.
	//
	// For producing, this is usually equal to the index of the record in
	// the record batch.
	OffsetDelta int32

	// Key is an blob of data for a record.
	//
	// Key's are usually used for hashing the record to specific Kafka partitions.
	Key []byte

	// Value is  a blob of data. This field is the main "message" portion of a
	// record.
	Value []byte

	// Headers are optional user provided metadata for records. Unlike normal
	// arrays, the number of headers is encoded as a varint.
	Headers []Header
}

// RecordBatch is a Kafka concept that groups many individual records together
// in a more optimized format.
type RecordBatch struct {
	// NullableBytesLength is not officially a field in a RecordBatch, however,
	// RecordBatch is a special form of NULLABLE_BYTES. Since all nullable bytes
	// must be prefixed with a int32 length, we throw that here.
	// This length should not include itself, only data that follows.
	NullableBytesLength int32

	// FirstOffset is the first offset in a record batch.
	//
	// For producing, this is usually 0.
	FirstOffset int64

	// Length is the wire length of everything that follows this field.
	Length int32

	// PartitionLeaderEpoch is a number that Kafka uses for cluster
	// communication. Clients generally do not need to worry about this
	// field and producers should set it to -1.
	PartitionLeaderEpoch int32

	// Magic is the current "magic" number of this message format.
	// The current magic number is 2.
	Magic int8

	// CRC is the crc of everything that follows this field using the
	// Castagnoli polynomial.
	CRC int32

	// Attributes describe the records array of this batch.
	//
	// Bits 0 thru 3 correspond to compression:
	//   - 000 is no compression
	//   - 001 is snappy compression
	//   - 010 is lz4 compression
	//   - 011 is zstd compression (produce request version 7+)
	//
	// Bit 4 is the timestamp type, with 0 meaning CreateTime corresponding
	// to the timestamp being from the producer, and 1 meaning LogAppendTime
	// corresponding to the timestamp being from the broker.
	//
	// Bit 5 indicates whether the batch is part of a transaction (1 is yes).
	//
	// Bit 6 indicates if the batch includes a control message (1 is yes).
	// Control messages are used to enable transactions and are generated from
	// the broker. Clients should not return control batches to applications.
	Attributes int16

	// LastOffsetDelta is the offset of the last message in a batch. This is
	// by the broker to ensure correct behavior even with batch compaction.
	LastOffsetDelta int32

	// FirstTimestamp is the timestamp (in milliseconds) of the first record
	// in a batch.
	FirstTimestamp int64

	// MaxTimestamp is the timestamp (in milliseconds) of the last record
	// in a batch. Similar to LastOffsetDelta, this is used to ensure correct
	// behavior with compacting.
	MaxTimestamp int64

	// ProducerId is the broker assigned producerId from an InitProducerId
	// request.
	//
	// Clients that wish to support idempotent messages and transactions must
	// set this field.
	ProducerId int64

	// ProducerEpoch is the broker assigned producerEpoch from an InitProducerId
	// request.
	//
	// Clients that wish to support idempotent messages and transactions must
	// set this field.
	ProducerEpoch int16

	// FirstSequence is the producer assigned sequence number used by the
	// broker to deduplicate messages.
	//
	// Clients that wish to support idempotent messages and transactions must
	// set this field.
	//
	// The sequence number for each record in a batch is OffsetDelta + FirstSequence.
	FirstSequence int32

	// Records are the batch of records to send.
	//
	// Note that to compress a batch, you compress the entire set of records
	// and then use that as the value for a single record.
	Records []Record
}
type ProduceRequestTopicDataData struct {
	Partition int32

	Records RecordBatch
}
type ProduceRequestTopicData struct {
	Topic string

	Data []ProduceRequestTopicDataData
}
type ProduceRequest struct {
	// Version is the version of this message used with a Kafka broker.
	Version       int16
	TransactionID *string // v3+

	Acks int16

	Timeout int32

	TopicData []ProduceRequestTopicData
}

func (*ProduceRequest) Key() int16                 { return 0 }
func (*ProduceRequest) MaxVersion() int16          { return 7 }
func (*ProduceRequest) MinVersion() int16          { return 2 }
func (v *ProduceRequest) SetVersion(version int16) { v.Version = version }
func (v *ProduceRequest) GetVersion() int16        { return v.Version }
func (v *ProduceRequest) ResponseKind() Response   { return &ProduceResponse{Version: v.Version} }

func (v *ProduceRequest) AppendTo(dst []byte) []byte {
	version := v.Version
	_ = version
	if version >= 3 {
		v := v.TransactionID
		dst = AppendNullableString(dst, v)
	}
	{
		v := v.Acks
		dst = AppendInt16(dst, v)
	}
	{
		v := v.Timeout
		dst = AppendInt32(dst, v)
	}
	{
		v := v.TopicData
		dst = AppendArrayLen(dst, len(v))
		for i := range v {
			v := &v[i]
			{
				v := v.Topic
				dst = AppendString(dst, v)
			}
			{
				v := v.Data
				dst = AppendArrayLen(dst, len(v))
				for i := range v {
					v := &v[i]
					{
						v := v.Partition
						dst = AppendInt32(dst, v)
					}
					{
						v := &v.Records
						{
							v := v.NullableBytesLength
							dst = AppendInt32(dst, v)
						}
						{
							v := v.FirstOffset
							dst = AppendInt64(dst, v)
						}
						{
							v := v.Length
							dst = AppendInt32(dst, v)
						}
						{
							v := v.PartitionLeaderEpoch
							dst = AppendInt32(dst, v)
						}
						{
							v := v.Magic
							dst = AppendInt8(dst, v)
						}
						{
							v := v.CRC
							dst = AppendInt32(dst, v)
						}
						{
							v := v.Attributes
							dst = AppendInt16(dst, v)
						}
						{
							v := v.LastOffsetDelta
							dst = AppendInt32(dst, v)
						}
						{
							v := v.FirstTimestamp
							dst = AppendInt64(dst, v)
						}
						{
							v := v.MaxTimestamp
							dst = AppendInt64(dst, v)
						}
						{
							v := v.ProducerId
							dst = AppendInt64(dst, v)
						}
						{
							v := v.ProducerEpoch
							dst = AppendInt16(dst, v)
						}
						{
							v := v.FirstSequence
							dst = AppendInt32(dst, v)
						}
						{
							v := v.Records
							dst = AppendArrayLen(dst, len(v))
							for i := range v {
								v := &v[i]
								{
									v := v.Length
									dst = AppendVarint(dst, v)
								}
								{
									v := v.Attributes
									dst = AppendInt8(dst, v)
								}
								{
									v := v.TimestampDelta
									dst = AppendVarint(dst, v)
								}
								{
									v := v.OffsetDelta
									dst = AppendVarint(dst, v)
								}
								{
									v := v.Key
									dst = AppendVarintBytes(dst, v)
								}
								{
									v := v.Value
									dst = AppendVarintBytes(dst, v)
								}
								{
									v := v.Headers
									dst = AppendVarint(dst, int32(len(v)))
									for i := range v {
										v := &v[i]
										{
											v := v.Key
											dst = AppendVarintString(dst, v)
										}
										{
											v := v.Value
											dst = AppendVarintBytes(dst, v)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return dst
}

type ProduceResponseResponsesPartitionResponses struct {
	Partition int32

	ErrorCode int16

	BaseOffset int64

	LogAppendTime int64 // v2+

	LogStartOffset int64 // v5+
}
type ProduceResponseResponses struct {
	Topic string

	PartitionResponses []ProduceResponseResponsesPartitionResponses
}
type ProduceResponse struct {
	// Version is the version of this message used with a Kafka broker.
	Version   int16
	Responses []ProduceResponseResponses

	ThrottleTimeMs int32
}

func (v *ProduceResponse) ReadFrom(src []byte) error {
	version := v.Version
	_ = version
	b := BinReader{Src: src}
	{
		s := v
		{
			v := s.Responses
			a := v
			for i := b.ArrayLen(); i > 0; i-- {
				a = append(a, ProduceResponseResponses{})
				v := &a[len(a)-1]
				{
					s := v
					{
						v := b.String()
						s.Topic = v
					}
					{
						v := s.PartitionResponses
						a := v
						for i := b.ArrayLen(); i > 0; i-- {
							a = append(a, ProduceResponseResponsesPartitionResponses{})
							v := &a[len(a)-1]
							{
								s := v
								{
									v := b.Int32()
									s.Partition = v
								}
								{
									v := b.Int16()
									s.ErrorCode = v
								}
								{
									v := b.Int64()
									s.BaseOffset = v
								}
								if version >= 2 {
									v := b.Int64()
									s.LogAppendTime = v
								}
								if version >= 5 {
									v := b.Int64()
									s.LogStartOffset = v
								}
							}
						}
						v = a
						s.PartitionResponses = v
					}
				}
			}
			v = a
			s.Responses = v
		}
		{
			v := b.Int32()
			s.ThrottleTimeMs = v
		}
	}
	return b.Complete()
}

// MetadataRequest requests metadata from Kafka.
type MetadataRequest struct {
	// Version is the version of this message used with a Kafka broker.
	Version int16
	// Topics is a list of topics to return metadata about. If this is null,
	// all topics are included. If this is empty, no topics are.
	// For v0 (<Kafka 0.10.0.0), if this is empty, all topics are included.
	Topics []string

	// AllowAutoTopicCreation, introduced in Kafka 0.11.0.0, allows topic
	// auto creation of the topics in this request if they do not exist.
	AllowAutoTopicCreation bool // v4+

	// IncludeClusterAuthorizedOperations, introduced in Kakfa 2.3.0, specifies
	// whether to return a bitfield of AclOperations that this client can perform
	// on the cluster.
	IncludeClusterAuthorizedOperations bool // v8+

	// IncludeTopicAuthorizedOperations, introduced in Kakfa 2.3.0, specifies
	// whether to return a bitfield of AclOperations that this client can perform
	// on individual topics.
	IncludeTopicAuthorizedOperations bool // v8+
}

func (*MetadataRequest) Key() int16                 { return 3 }
func (*MetadataRequest) MaxVersion() int16          { return 8 }
func (*MetadataRequest) MinVersion() int16          { return 0 }
func (v *MetadataRequest) SetVersion(version int16) { v.Version = version }
func (v *MetadataRequest) GetVersion() int16        { return v.Version }
func (v *MetadataRequest) ResponseKind() Response   { return &MetadataResponse{Version: v.Version} }

func (v *MetadataRequest) AppendTo(dst []byte) []byte {
	version := v.Version
	_ = version
	{
		v := v.Topics
		dst = AppendArrayLen(dst, len(v))
		for i := range v {
			v := v[i]
			dst = AppendString(dst, v)
		}
	}
	if version >= 4 {
		v := v.AllowAutoTopicCreation
		dst = AppendBool(dst, v)
	}
	if version >= 8 {
		v := v.IncludeClusterAuthorizedOperations
		dst = AppendBool(dst, v)
	}
	if version >= 8 {
		v := v.IncludeTopicAuthorizedOperations
		dst = AppendBool(dst, v)
	}
	return dst
}

type MetadataResponseBrokers struct {
	// NodeID is the node ID of a Kafka broker.
	NodeID int32

	// Host is the hostname of a Kafka broker.
	Host string

	// Port is the port of a Kafka broker.
	Port int32

	// Rack is the rack this Kafka broker is in.
	Rack *string // v1+
}
type MetadataResponseTopicMetadataPartitionMetadata struct {
	// ErrorCode is any error for a partition in topic metadata.
	//
	// LEADER_NOT_AVAILABLE is returned if a leader is unavailable for this
	// partition. For v0 metadata responses, this is also returned if a
	// partition leader's listener does not exist.
	//
	// LISTENER_NOT_FOUND is returned if a leader ID is known but the
	// listener for it is not (v1+).
	//
	// REPLICA_NOT_AVAILABLE is returned in v0 responses if any replica is
	// unavailable.
	ErrorCode int16

	// Partition is a partition number for a topic.
	Partition int32

	// Leader is the broker leader for this partition. This will be -1
	// on leader / listener error.
	Leader int32

	// LeaderEpoch, proposed in KIP-320 and introduced in Kafka 2.1.0  is the
	// epoch of the broker leader.
	LeaderEpoch int32 // v7+

	// Replicas returns all broker IDs containing replicas of this partition.
	Replicas []int32

	// ISR returns all broker IDs of in-sync replicas of this partition.
	ISR []int32

	// OfflineReplicas, proposed in KIP-112 and introduced in Kafka 1.0,
	// returns all offline broker IDs that should be replicating this partition.
	OfflineReplicas []int32 // v5+
}
type MetadataResponseTopicMetadata struct {
	// ErrorCode is any error for a topic in a metadata request.
	//
	// TOPIC_AUTHORIZATION_FAILED: returned if the client is not authorized
	// to describe the topic, or if the metadata request specified topic auto
	// creation, the topic did not exist, and the user lacks permission to create.
	//
	// UNKNOWN_TOPIC_OR_PARTITION: returned if a topic does not exist and
	// the request did not specify autocreation.
	//
	// LEADER_NOT_AVAILABLE: returned if a new topic is created successfully
	// (since there is no leader on an immediately new topic).
	//
	// There can be a myriad of other errors for unsuccessful topic creation.
	ErrorCode int16

	// Topic is the topic this metadata corresponds to.
	Topic string

	// IsInternal signifies whether this topic is a Kafka internal topic.
	IsInternal bool // v1+

	// PartitionMetadata contains metadata about partitions for a topic.
	PartitionMetadata []MetadataResponseTopicMetadataPartitionMetadata

	// AuthorizedOperations, proposed in KIP-430 and introduced in Kafka 2.3.0,
	// returns a bitfield (corresponding to AclOperation) containing which
	// operations the client is allowed to perform on this topic.
	// This is only returned if requested.
	AuthorizedOperations int32 // v8+
}

// MetadataResponse is returned for a MetdataRequest.
type MetadataResponse struct {
	// Version is the version of this message used with a Kafka broker.
	Version int16
	// ThrottleTimeMs is how long of a throttle Kafka will apply to the client
	// after this request.
	// For Kafka < 2.0.0, the throttle is applied before issuing a response.
	// For Kafka >= 2.0.0, the throttle is applied after issuing a response.
	ThrottleTimeMs int32 // v3+

	// Brokers is a set of alive Kafka brokers.
	Brokers []MetadataResponseBrokers

	// ClusterID, proposed in KIP-78 and introduced in Kafka 0.10.1.0, is a
	// unique string specifying the cluster that the replying Kafka belongs to.
	ClusterID *string // v2+

	// ControllerID is the ID of the controller broker (the admin broker).
	ControllerID int32 // v1+

	// TopicMetadata contains metadata about each topic requested in the
	// MetadataRequest.
	TopicMetadata []MetadataResponseTopicMetadata

	// AuthorizedOperations returns a bitfield containing which operations the
	// client is allowed to perform on this cluster.
	AuthorizedOperations int32 // v8+
}

func (v *MetadataResponse) ReadFrom(src []byte) error {
	version := v.Version
	_ = version
	b := BinReader{Src: src}
	{
		s := v
		if version >= 3 {
			v := b.Int32()
			s.ThrottleTimeMs = v
		}
		{
			v := s.Brokers
			a := v
			for i := b.ArrayLen(); i > 0; i-- {
				a = append(a, MetadataResponseBrokers{})
				v := &a[len(a)-1]
				{
					s := v
					{
						v := b.Int32()
						s.NodeID = v
					}
					{
						v := b.String()
						s.Host = v
					}
					{
						v := b.Int32()
						s.Port = v
					}
					if version >= 1 {
						v := b.NullableString()
						s.Rack = v
					}
				}
			}
			v = a
			s.Brokers = v
		}
		if version >= 2 {
			v := b.NullableString()
			s.ClusterID = v
		}
		if version >= 1 {
			v := b.Int32()
			s.ControllerID = v
		}
		{
			v := s.TopicMetadata
			a := v
			for i := b.ArrayLen(); i > 0; i-- {
				a = append(a, MetadataResponseTopicMetadata{})
				v := &a[len(a)-1]
				{
					s := v
					{
						v := b.Int16()
						s.ErrorCode = v
					}
					{
						v := b.String()
						s.Topic = v
					}
					if version >= 1 {
						v := b.Bool()
						s.IsInternal = v
					}
					{
						v := s.PartitionMetadata
						a := v
						for i := b.ArrayLen(); i > 0; i-- {
							a = append(a, MetadataResponseTopicMetadataPartitionMetadata{})
							v := &a[len(a)-1]
							{
								s := v
								{
									v := b.Int16()
									s.ErrorCode = v
								}
								{
									v := b.Int32()
									s.Partition = v
								}
								{
									v := b.Int32()
									s.Leader = v
								}
								if version >= 7 {
									v := b.Int32()
									s.LeaderEpoch = v
								}
								{
									v := s.Replicas
									a := v
									for i := b.ArrayLen(); i > 0; i-- {
										v := b.Int32()
										a = append(a, v)
									}
									v = a
									s.Replicas = v
								}
								{
									v := s.ISR
									a := v
									for i := b.ArrayLen(); i > 0; i-- {
										v := b.Int32()
										a = append(a, v)
									}
									v = a
									s.ISR = v
								}
								if version >= 5 {
									v := s.OfflineReplicas
									a := v
									for i := b.ArrayLen(); i > 0; i-- {
										v := b.Int32()
										a = append(a, v)
									}
									v = a
									s.OfflineReplicas = v
								}
							}
						}
						v = a
						s.PartitionMetadata = v
					}
					if version >= 8 {
						v := b.Int32()
						s.AuthorizedOperations = v
					}
				}
			}
			v = a
			s.TopicMetadata = v
		}
		if version >= 8 {
			v := b.Int32()
			s.AuthorizedOperations = v
		}
	}
	return b.Complete()
}

type ApiVersionsRequest struct {
	// Version is the version of this message used with a Kafka broker.
	Version int16
}

func (*ApiVersionsRequest) Key() int16                 { return 18 }
func (*ApiVersionsRequest) MaxVersion() int16          { return 2 }
func (*ApiVersionsRequest) MinVersion() int16          { return 0 }
func (v *ApiVersionsRequest) SetVersion(version int16) { v.Version = version }
func (v *ApiVersionsRequest) GetVersion() int16        { return v.Version }
func (v *ApiVersionsRequest) ResponseKind() Response   { return &ApiVersionsResponse{Version: v.Version} }

func (v *ApiVersionsRequest) AppendTo(dst []byte) []byte {
	version := v.Version
	_ = version
	return dst
}

type ApiVersionsResponseApiVersions struct {
	ApiKey int16

	MinVersion int16

	MaxVersion int16
}
type ApiVersionsResponse struct {
	// Version is the version of this message used with a Kafka broker.
	Version   int16
	ErrorCode int16

	ApiVersions []ApiVersionsResponseApiVersions

	ThrottleTimeMs int32 // v1+
}

func (v *ApiVersionsResponse) ReadFrom(src []byte) error {
	version := v.Version
	_ = version
	b := BinReader{Src: src}
	{
		s := v
		{
			v := b.Int16()
			s.ErrorCode = v
		}
		{
			v := s.ApiVersions
			a := v
			for i := b.ArrayLen(); i > 0; i-- {
				a = append(a, ApiVersionsResponseApiVersions{})
				v := &a[len(a)-1]
				{
					s := v
					{
						v := b.Int16()
						s.ApiKey = v
					}
					{
						v := b.Int16()
						s.MinVersion = v
					}
					{
						v := b.Int16()
						s.MaxVersion = v
					}
				}
			}
			v = a
			s.ApiVersions = v
		}
		if version >= 1 {
			v := b.Int32()
			s.ThrottleTimeMs = v
		}
	}
	return b.Complete()
}

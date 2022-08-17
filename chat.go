package main

//type RealtimeHub struct {
//	Messenger KafkaProvider
//}
//
//func NewRealtimeHub(messenger KafkaProvider) RealtimeHub {
//	return RealtimeHub{
//		Messenger: messenger,
//	}
//}
//func (h *RealtimeHub) Get() interface{} {
//	return h
//}
//func (h *RealtimeHub) Run() {
//	go h.SubscribeAllPartitions()
//}
//
//func (h *RealtimeHub) SubscribeAllPartitions() {
//	kafkaHost := "localhost:9092"
//
//	conn, err := kafka.Dial("tcp", kafkaHost)
//	if err != nil {
//		logrus.Errorf("SubscribeAllPartitions: error kafka dialing %v", err)
//		return
//	}
//
//	var controllerConn *kafka.Conn
//	defer func() {
//		_ = conn.Close()
//		_ = controllerConn.Close()
//	}()
//
//	controller, err := conn.Controller()
//	if err != nil {
//		logrus.Errorf("startNotifier: error kafka dialing controller %v", err)
//		return
//	}
//
//	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
//	if err != nil {
//		panic(err.Error())
//	}
//
//	err = controllerConn.CreateTopics(kafka.TopicConfig{
//		Topic:             "testing101",
//		NumPartitions:     5,
//		ReplicationFactor: 1,
//	})
//	if err != nil {
//		logrus.Errorf("SubscribeAllPartitions: error creating topic %v", err)
//		return
//	}
//
//	h.Subscribe()
//}
//
//func (h *RealtimeHub) Subscribe() {
//	kafkaHost := "localhost:9092"
//
//	r := kafka.NewReader(kafka.ReaderConfig{
//		Brokers: []string{kafkaHost},
//		Topic:   "testing101",
//		GroupID: "g1",
//	})
//	_ = r.SetOffset(kafka.LastOffset)
//
//	for {
//		m, err := r.ReadMessage(context.Background())
//		if err != nil {
//			break
//		}
//
//		var message Message
//		if err = json.Unmarshal(m.Value, &message); err != nil {
//			logrus.Errorf("Failed to Unmarshal message from topic: %q error: %v", m.Topic, err)
//			return
//		}
//
//		logrus.Println(message)
//		fmt.Println(message)
//
//	}
//}

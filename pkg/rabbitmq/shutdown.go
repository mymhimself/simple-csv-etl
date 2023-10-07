package rabbitmq

// GracefulShutdown can be called by a consumer to gracfully shutdown after getting all deliveris processed
func (rb *rabbit) GracefulShutdown() error {
	for _, consumerName := range rb.consumerNames {
		rb.gracefulShutdown <- struct{}{}
		err := rb.chann.Cancel(consumerName, false)
		if err != nil {
			return err
		}
	}
	rb.wg.Wait()
	if rb.conn != nil {
		if err := rb.conn.Close(); err != nil {
			return err
		}
	}
	return nil
}

// CloseConnection closes the connection. It is used to be called from the publisher application
// before terminationg the app to make sure all publishings has been received to the server
func (rb *rabbit) CloseConnection() error {
	return rb.conn.Close()
}

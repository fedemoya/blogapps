package messages

import (
    "bytes"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "io/ioutil"
    "net/http"
    "sync"
    "testing"
    "time"
)

func TestWebhookMessageConsumer(t *testing.T) {

    webhookMessageConsumer := NewWebhookMessageConsumer()

    messagesCh, err := webhookMessageConsumer.Consume()
    require.NoError(t, err, "error starting webhook message consumer")

    // wait server start
    time.Sleep(1 * time.Second)

    payload := []byte(`message payload`)
    require.NoError(t, err, "failed reading test data")

    wg := &sync.WaitGroup{}
    wg.Add(1)

    go func() {
        msg := <-messagesCh
        assert.Equal(t, payload, msg.Payload())
        wg.Done()
    }()

    resp, err := http.DefaultClient.Post("http://127.0.0.1:8080/webhooks", "application/json", ioutil.NopCloser(bytes.NewReader(payload)))
    require.NoError(t, err, "error sending post request to webhook message consumer")
    assert.Equal(t, resp.StatusCode, 201)

    wg.Wait()

    err = webhookMessageConsumer.Close()
    require.NoError(t, err, "failed closing webhook message consumer")
}

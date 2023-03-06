package messages

import (
    "context"
    "errors"
    "github.com/gin-gonic/gin"
    "io"
    "log"
    "net/http"
    "time"
)

type WebhookMessageConsumer struct {
    server *http.Server
}

func NewWebhookMessageConsumer() *WebhookMessageConsumer {
    return &WebhookMessageConsumer{}
}

func (w *WebhookMessageConsumer) Consume() (<-chan Message, error) {

    router := gin.Default()

    messagesCh := make(chan Message)

    router.POST("/webhooks", func(c *gin.Context) {
        body, err := io.ReadAll(c.Request.Body)
        if err != nil {
            log.Printf("failed reading webhook request body: %s", err.Error())
            c.Status(http.StatusInternalServerError)
            return
        }
        message := NewWebhookMessage(body)
        messagesCh <- message

        // TODO wait for the message to be Acked or Nacked

        c.Status(http.StatusCreated)
    })

    w.server = &http.Server{
        Addr:    "127.0.0.1:8080",
        Handler: router,
    }

    // Initializing the server in a goroutine so that
    // it won't block the graceful shutdown handling below
    go func() {
        if err := w.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Printf("webhook message consumer closed unexpectedly: %s\n", err)
        }
    }()

    return messagesCh, nil
}

func (w *WebhookMessageConsumer) Close() error {
    log.Printf("closing message consumer")

    // The context is used to inform the server it has 5 seconds to finish
    // the request it is currently handling
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := w.server.Shutdown(ctx); err != nil {
        log.Printf("server forced to shutdown: %s", err.Error())
    }
    return nil
}

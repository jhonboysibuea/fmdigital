package middleware

import (
	"context"
	"crud-app/consts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ContextKey is a custom type to avoid key collisions in context values

// Middleware function to set trace-id in the context and response headers
func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a trace-id
		traceID := uuid.New().String()

		// Set trace-id in response headers
		c.Header("X-Trace-ID", traceID)

		// Create a new context with trace-id
		ctx := context.WithValue(c.Request.Context(), consts.TraceIDKey, traceID)

		// Use the new context in the request
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

package constant

type contextKey string

const ContextKeyTraceID contextKey = "TraceID"
const ContextKeySpanID contextKey = "SpanID"
const ContextKeyUserID contextKey = "UserID"
const ContextKeyUsername contextKey = "Username"
const ContextKeyEmail contextKey = "Email"
const ContextKeyUserType contextKey = "UserType"
const ContextKeyOrganizationID contextKey = "OrganizationID"
const ContextKeyOrganizationName contextKey = "OrganizationName"
const TokenExpiresAt contextKey = "TokenExpiresAt"

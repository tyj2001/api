package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/samber/lo"
	"github.com/swuecho/chatgpt_backend/sqlc_queries"
)

// ChatSessionService provides methods for interacting with chat sessions.
type ChatSessionService struct {
	q *sqlc_queries.Queries
}

// NewChatSessionService creates a new ChatSessionService.
func NewChatSessionService(q *sqlc_queries.Queries) *ChatSessionService {
	return &ChatSessionService{q: q}
}

// CreateChatSession creates a new chat session.
func (s *ChatSessionService) CreateChatSession(ctx context.Context, session_params sqlc_queries.CreateChatSessionParams) (sqlc_queries.ChatSession, error) {
	session, err := s.q.CreateChatSession(ctx, session_params)
	if err != nil {
		return sqlc_queries.ChatSession{}, err
	}
	return session, nil
}

// GetChatSessionByID returns a chat session by ID.
func (s *ChatSessionService) GetChatSessionByID(ctx context.Context, id int32) (sqlc_queries.ChatSession, error) {
	session, err := s.q.GetChatSessionByID(ctx, id)
	if err != nil {
		return sqlc_queries.ChatSession{}, fmt.Errorf("failed to retrieve session: %w", err)
	}
	return session, nil
}

// UpdateChatSession updates an existing chat session.
func (s *ChatSessionService) UpdateChatSession(ctx context.Context, session_params sqlc_queries.UpdateChatSessionParams) (sqlc_queries.ChatSession, error) {
	session_u, err := s.q.UpdateChatSession(ctx, session_params)
	if err != nil {
		return sqlc_queries.ChatSession{}, errors.New("failed to update session")
	}
	return session_u, nil
}

// DeleteChatSession deletes a chat session by ID.
func (s *ChatSessionService) DeleteChatSession(ctx context.Context, id int32) error {
	err := s.q.DeleteChatSession(ctx, id)
	if err != nil {
		return errors.New("failed to delete session by id")
	}
	return nil
}

// GetAllChatSessions returns all chat sessions.
func (s *ChatSessionService) GetAllChatSessions(ctx context.Context) ([]sqlc_queries.ChatSession, error) {
	sessions, err := s.q.GetAllChatSessions(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve sessions")
	}
	return sessions, nil
}

func (s *ChatSessionService) GetChatSessionsByUserID(ctx context.Context, userID int32) ([]sqlc_queries.ChatSession, error) {
	sessions, err := s.q.GetChatSessionsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to retrieve sessions")
	}
	return sessions, nil
}

func (s *ChatSessionService) GetSimpleChatSessionsByUserID(ctx context.Context, userID int32) ([]SimpleChatSession, error) {
	sessions, err := s.q.GetChatSessionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	simple_sessions := lo.Map(sessions, func(session sqlc_queries.ChatSession, _idx int) SimpleChatSession {
		return SimpleChatSession{
			Uuid:   session.Uuid,
			IsEdit: false,
			Title:  session.Topic,
		}
	})
	return simple_sessions, nil
}

// GetChatSessionByUUID returns an authentication user record by ID.
func (s *ChatSessionService) GetChatSessionByUUID(ctx context.Context, uuid string) (sqlc_queries.ChatSession, error) {
	chatSession, err := s.q.GetChatSessionByUUID(ctx, uuid)
	if err != nil {
		return sqlc_queries.ChatSession{}, fmt.Errorf("failed to retrieve session by uuid, %w", err)
	}
	return chatSession, nil
}

// UpdateChatSessionByUUID updates an existing chat session.
func (s *ChatSessionService) UpdateChatSessionByUUID(ctx context.Context, session_params sqlc_queries.UpdateChatSessionByUUIDParams) (sqlc_queries.ChatSession, error) {
	session_u, err := s.q.UpdateChatSessionByUUID(ctx, session_params)
	if err != nil {
		return sqlc_queries.ChatSession{}, fmt.Errorf("failed to update session, %w", err)
	}
	return session_u, nil
}

// UpdateChatSessionTopicByUUID updates an existing chat session topic.
func (s *ChatSessionService) UpdateChatSessionTopicByUUID(ctx context.Context, session_params sqlc_queries.UpdateChatSessionTopicByUUIDParams) (sqlc_queries.ChatSession, error) {
	session_u, err := s.q.UpdateChatSessionTopicByUUID(ctx, session_params)
	if err != nil {
		return sqlc_queries.ChatSession{}, fmt.Errorf("failed to update session, %w", err)
	}
	return session_u, nil
}

// CreateOrUpdateChatSessionByUUID updates an existing chat session.
func (s *ChatSessionService) CreateOrUpdateChatSessionByUUID(ctx context.Context, session_params sqlc_queries.CreateOrUpdateChatSessionByUUIDParams) (sqlc_queries.ChatSession, error) {
	session_u, err := s.q.CreateOrUpdateChatSessionByUUID(ctx, session_params)
	if err != nil {
		return sqlc_queries.ChatSession{}, fmt.Errorf("failed to update session, %w", err)
	}
	return session_u, nil
}

// DeleteChatSessionByUUID deletes a chat session by UUID.
func (s *ChatSessionService) DeleteChatSessionByUUID(ctx context.Context, uuid string) error {
	err := s.q.DeleteChatSessionByUUID(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete session by uuid, %w", err)

	}
	return nil
}

// UpdateSessionMaxLength
func (s *ChatSessionService) UpdateSessionMaxLength(ctx context.Context, session_params sqlc_queries.UpdateSessionMaxLengthParams) (sqlc_queries.ChatSession, error) {
	session_u, err := s.q.UpdateSessionMaxLength(ctx, session_params)
	if err != nil {
		return sqlc_queries.ChatSession{}, fmt.Errorf("failed to update session, %w", err)
	}
	return session_u, nil
}
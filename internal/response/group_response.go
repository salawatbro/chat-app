package response

import "github.com/salawatbro/chat-app/internal/models"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GroupResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedBy User   `json:"created_by"`
}

type GroupsResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewGroupResponse(group *models.Chat) *GroupResponse {
	return &GroupResponse{
		ID:   group.ID.Hex(),
		Name: group.ChatName,
		CreatedBy: User{
			ID:    group.Owner.ID.Hex(),
			Name:  group.Owner.Name,
			Email: group.Owner.Email,
		},
	}
}

func NewGroupsResponse(groups []*models.Chat) *[]GroupsResponse {
	groupResponses := make([]GroupsResponse, 0)
	for _, group := range groups {
		groupResponses = append(groupResponses, GroupsResponse{
			ID:   group.ID.Hex(),
			Name: group.ChatName,
		})
	}
	return &groupResponses
}

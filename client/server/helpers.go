package server

import (
	"compendium/models"
	"context"
)

func (s *Server) GetTokenIdentity(token string) *models.Identity {
	i := s.db.Temp.IdentityRead(context.TODO(), token)
	if i.Token != "" {
		return &i
	}
	return nil
}

func (s *Server) GetCorpData(i *models.Identity, roleId string) models.CorpData {
	c := models.CorpData{}
	c.Members = []models.CorpMemberint{}

	if i.Guild.ID != "" {
		c.Roles = s.ds.GetRoles(i.Guild.ID)
		cm := s.db.Temp.CorpMemberReadAllByGuildId(context.TODO(), i.Guild.ID)
		for _, member := range cm {
			if s.ds.CheckRole(i.Guild.ID, member.UserId, roleId) {
				c.Members = append(c.Members, member)
			}
		}
	}
	return c
}

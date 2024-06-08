package entity

import "github.com/Borislavv/go-seo/internal/shared/domain/vo"

type Project struct {
	ID         vo.ID  `json:"id"`
	RefID      vo.ID  `json:"ref_id"`
	DomainName string `json:"domain_name"`
}

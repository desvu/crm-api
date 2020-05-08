package url_builder

import (
	"strings"

	"github.com/qilin/crm-api/internal/domain/entity"
)

func (h Helper) BuildGameMedia(i *entity.GameMedia) string {
	return strings.Join([]string{h.Env.App.StorageURL, i.FilePath}, "")
}

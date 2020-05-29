package document

import (
	"github.com/qilin/crm-api/internal/domain/enum/document"
	"github.com/qilin/crm-api/internal/domain/enum/language"
)

var titleNames = map[language.Language]map[document.Type]string{
	language.English: map[document.Type]string{
		document.TypePrivacyPolicy: "Privacy Policy",
		document.TypeGameRules:     "Game Rules",
		document.TypeLicense:       "EULA",
	},
	language.Russian: map[document.Type]string{
		document.TypePrivacyPolicy: "Политика конфиденциальности",
		document.TypeGameRules:     "Правила игра",
		document.TypeLicense:       "Лицензионное соглашение",
	},
}

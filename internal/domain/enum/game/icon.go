package game

type Icon string

const (
	IconUndefined Icon = ""
	IconGamepad   Icon = "gamepad" // 🎮 video game (U+1F3AE)
	IconCloud     Icon = "cloud"   // ☁️ cloud (U+2601)
	IconPeople    Icon = "people"  // 👥 busts in silhouette (U+1F465)
)

func NewIcon(name string) Icon {
	switch Icon(name) {
	case IconGamepad:
		return IconGamepad
	case IconCloud:
		return IconCloud
	case IconPeople:
		return IconPeople
	}
	return IconUndefined
}

func (i Icon) String() string {
	return string(i)
}

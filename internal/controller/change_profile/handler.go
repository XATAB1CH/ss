package changeprofile

type ChangeProfileController interface {
	// ChangeCity() (string, error)
	// ChangeDescription() (string, error)
}

type Handler struct {
	ChangeProfileController ChangeProfileController
}

func New(ChangeProfileController ChangeProfileController) *Handler {
	return &Handler{
		ChangeProfileController: ChangeProfileController,
	}
}

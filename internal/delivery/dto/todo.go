package dto

type CreateToDoRequest struct {
	//UserID uint   `json:"user_id"`
	Text string `json:"text" binding:"required,min=1,max=500"`
}

type ToDorResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
}
